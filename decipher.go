package youtube

import (
   "errors"
   "fmt"
   "log"
   "net/url"
   "regexp"
   "strconv"
   "time"
)

func (c *Client) decipherURL(videoID string, cipher string) (string, error) {
   queryParams, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   /* eg:
   extract decipher from  https://youtube.com/s/player/4fbb4d5b/player_ias.vflset/en_US/base.js

   var Mt={
   splice:function(a,b){a.splice(0,b)},
   reverse:function(a){a.reverse()},
   EQ:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c}};

   a=a.split("");
   Mt.splice(a,3);
   Mt.EQ(a,39);
   Mt.splice(a,2);
   Mt.EQ(a,1);
   Mt.splice(a,1);
   Mt.EQ(a,35);
   Mt.EQ(a,51);
   Mt.splice(a,2);
   Mt.reverse(a,52);
   return a.join("")
   */
   operations, err := c.parseDecipherOpsWithCache(videoID)
   if err != nil { return "", err }
   // apply operations
   bs := []byte(queryParams.Get("s"))
   for _, op := range operations {
      bs = op(bs)
   }
   decipheredURL := fmt.Sprintf("%s&%s=%s", queryParams.Get("url"), queryParams.Get("sp"), string(bs))
   return decipheredURL, nil
}

const (
   jsvarStr = "[a-zA-Z_\\$][a-zA-Z_0-9]*"
   reverseStr = ":function\\(a\\)\\{(?:return )?a\\.reverse\\(\\)\\}"
   spliceStr = ":function\\(a,b\\)\\{a\\.splice\\(0,b\\)\\}"
)

var (
   swapStr = fmt.Sprint(
      `:function\(a,b\)\{var c=a\[0\];a\[0\]=a\[b(?:%a\.length)?\];`,
      `a\[b(?:%a\.length)?\]=c(?:;return a)?\}`,
   )
   basejsPattern = regexp.MustCompile(`(/s/player/\w+/player_ias.vflset/\w+/base.js)`)
   actionsObjRegexp = regexp.MustCompile(fmt.Sprintf(
      "var (%s)=\\{((?:(?:%s%s|%s%s|%s%s),?\\n?)+)\\};",
      jsvarStr, jsvarStr, swapStr, jsvarStr, spliceStr, jsvarStr, reverseStr,
   ))
   reverseRegexp = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, reverseStr))
   spliceRegexp  = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, spliceStr))
   swapRegexp    = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, swapStr))
   actionsFuncRegexp = regexp.MustCompile(fmt.Sprintf(
      `function(?: %s)?\(a\)\{a=a\.split\(""\);` +
      `\s*((?:(?:a=)?%s\.%s\(a,\d+\);)+)return a\.join\(""\)\}`,
      jsvarStr, jsvarStr, jsvarStr,
   ))
)

func (c *Client) parseDecipherOps(videoID string) (operations []DecipherOperation, err error) {
	embedURL := fmt.Sprintf("https://youtube.com/embed/%s?hl=en", videoID)
	embedBody, err := c.httpGetBodyBytes(embedURL)
	if err != nil {
		return nil, err
	}

	// example: /s/player/f676c671/player_ias.vflset/en_US/base.js
	escapedBasejsURL := string(basejsPattern.Find(embedBody))
	if escapedBasejsURL == "" {
		log.Println("playerConfig:", string(embedBody))
		return nil, errors.New("unable to find basejs URL in playerConfig")
	}

	basejsBody, err := c.httpGetBodyBytes("https://youtube.com"+escapedBasejsURL)
	if err != nil {
		return nil, err
	}

	objResult := actionsObjRegexp.FindSubmatch(basejsBody)
	funcResult := actionsFuncRegexp.FindSubmatch(basejsBody)
	if len(objResult) < 3 || len(funcResult) < 2 {
		return nil, fmt.Errorf("error parsing signature tokens (#obj=%d, #func=%d)", len(objResult), len(funcResult))
	}

	obj := objResult[1]
	objBody := objResult[2]
	funcBody := funcResult[1]

	var reverseKey, spliceKey, swapKey string

	if result := reverseRegexp.FindSubmatch(objBody); len(result) > 1 {
		reverseKey = string(result[1])
	}
	if result := spliceRegexp.FindSubmatch(objBody); len(result) > 1 {
		spliceKey = string(result[1])
	}
	if result := swapRegexp.FindSubmatch(objBody); len(result) > 1 {
		swapKey = string(result[1])
	}

	regex, err := regexp.Compile(fmt.Sprintf("(?:a=)?%s\\.(%s|%s|%s)\\(a,(\\d+)\\)", obj, reverseKey, spliceKey, swapKey))
	if err != nil {
		return nil, err
	}

	var ops []DecipherOperation
	for _, s := range regex.FindAllSubmatch(funcBody, -1) {
		switch string(s[1]) {
		case reverseKey:
			ops = append(ops, reverseFunc)
		case swapKey:
			arg, _ := strconv.Atoi(string(s[2]))
			ops = append(ops, newSwapFunc(arg))
		case spliceKey:
			arg, _ := strconv.Atoi(string(s[2]))
			ops = append(ops, newSpliceFunc(arg))
		}
	}
	return ops, nil
}

func (c *Client) parseDecipherOpsWithCache(videoID string) (operations []DecipherOperation, err error) {
	if c.decipherOpsCache == nil {
		c.decipherOpsCache = NewSimpleCache()
	}

	if ops := c.decipherOpsCache.Get(videoID); ops != nil {
		return ops, nil
	}

	ops, err := c.parseDecipherOps(videoID)
	if err != nil {
		return nil, err
	}

	c.decipherOpsCache.Set(videoID, ops)
	return ops, err
}


type DecipherOperation func([]byte) []byte

func newSpliceFunc(pos int) DecipherOperation {
	return func(bs []byte) []byte {
		return bs[pos:]
	}
}

func newSwapFunc(arg int) DecipherOperation {
	return func(bs []byte) []byte {
		pos := arg % len(bs)
		bs[0], bs[pos] = bs[pos], bs[0]
		return bs
	}
}

func reverseFunc(bs []byte) []byte {
	l, r := 0, len(bs)-1
	for l < r {
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
	return bs
}

var (
	_ DecipherOperationsCache = NewSimpleCache()
)

const defaultCacheExpiration = time.Minute * time.Duration(5)

type DecipherOperationsCache interface {
	Get(videoID string) []DecipherOperation
	Set(video string, operations []DecipherOperation)
}

type SimpleCache struct {
	videoID    string
	expiredAt  time.Time
	operations []DecipherOperation
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{}
}

// Get : get cache  when it has same video id and not expired
func (s SimpleCache) Get(videoID string) []DecipherOperation {
	return s.GetCacheBefore(videoID, time.Now())
}

// GetCacheBefore : can pass time for testing
func (s SimpleCache) GetCacheBefore(videoID string, time time.Time) []DecipherOperation {
	if videoID == s.videoID && s.expiredAt.After(time) {
		operations := make([]DecipherOperation, len(s.operations))
		copy(operations, s.operations)
		return operations
	}
	return nil
}

// Set : set cache with default expiration
func (s *SimpleCache) Set(videoID string, operations []DecipherOperation) {
	s.setWithExpiredTime(videoID, operations, time.Now().Add(defaultCacheExpiration))
}

func (s *SimpleCache) setWithExpiredTime(videoID string, operations []DecipherOperation, time time.Time) {
	s.videoID = videoID
	s.operations = make([]DecipherOperation, len(operations))
	copy(s.operations, operations)
	s.expiredAt = time
}


const (
	ErrCipherNotFound             = constError("cipher not found")
	ErrInvalidCharactersInVideoID = constError("invalid characters in video id")
	ErrVideoIDMinLength           = constError("the video id must be at least 10 characters long")
	ErrReadOnClosedResBody        = constError("http: read on closed response body")
	ErrNotPlayableInEmbed         = constError("embedding of this video has been disabled")
	ErrInvalidPlaylist            = constError("no playlist detected or invalid playlist ID")
)

type constError string

func (e constError) Error() string {
	return string(e)
}
