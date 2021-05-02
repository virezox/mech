package youtube

import (
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "log"
   "net/http"
   "net/url"
   "regexp"
   "sort"
   "strconv"
   "time"
)

const API = "https://www.youtube.com/get_video_info"

type Video struct {
   StreamingData struct {
      AdaptiveFormats []struct {
         Bitrate int
         Height int
         Itag int
         MimeType string
         SignatureCipher string
      }
      ExpiresInSeconds string
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         Description struct { SimpleText string }
         PublishDate string
         Title struct { SimpleText string }
         ViewCount int `json:",string"`
      }
   }
   VideoDetails struct { VideoId string }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   val := make(url.Values)
   val.Set("video_id", id)
   val.Set("eurl", "https://youtube.googleapis.com/v/" + id)
   body, err := readAll(API + "?" + val.Encode())
   if err != nil {
      return Video{}, err
   }
   val, err = url.ParseQuery(string(body))
   if err != nil {
      return Video{}, err
   }
   var (
      play = val.Get("player_response")
      vid Video
   )
   err = json.Unmarshal([]byte(play), &vid)
   if err != nil {
      return Video{}, err
   }
   return vid, nil
}

func (v Video) Description() string {
   return v.Microformat.PlayerMicroformatRenderer.Description.SimpleText
}

// GetStream returns the url for a specific format
func (v Video) GetStream(itag int) (string, error) {
   var cipher string
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { cipher = format.SignatureCipher }
   }
   query, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   decipherOpsCache := new(simpleCache)
   operations := decipherOpsCache.get(v.VideoDetails.VideoId)
   if operations == nil {
      operations, err = parseDecipherOps(v.VideoDetails.VideoId)
      if err != nil { return "", err }
      decipherOpsCache.set(v.VideoDetails.VideoId, operations)
   }
   // apply operations
   bs := []byte(query.Get("s"))
   for _, op := range operations {
      bs = op(bs)
   }
   return fmt.Sprintf("%s&%s=%s", query.Get("url"), query.Get("sp"), bs), nil
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string {
   return v.Microformat.PlayerMicroformatRenderer.Title.SimpleText
}

func (v Video) ViewCount() int {
   return v.Microformat.PlayerMicroformatRenderer.ViewCount
}


const (
   jsvarStr = `[a-zA-Z_\$][a-zA-Z_0-9]*`
   reverseStr = ":function\\(a\\)\\{(?:return )?a\\.reverse\\(\\)\\}"
   spliceStr = ":function\\(a,b\\)\\{a\\.splice\\(0,b\\)\\}"
)

var (
   actionsFuncRegexp = regexp.MustCompile(fmt.Sprintf(
      `function(?: %s)?\(a\)\{a=a\.split\(""\);` +
      `\s*((?:(?:a=)?%s\.%s\(a,\d+\);)+)return a\.join\(""\)\}`,
      jsvarStr, jsvarStr, jsvarStr,
   ))
   actionsObjRegexp = regexp.MustCompile(fmt.Sprintf(
      "var (%s)=\\{((?:(?:%s%s|%s%s|%s%s),?\\n?)+)\\};",
      jsvarStr, jsvarStr, swapStr, jsvarStr, spliceStr, jsvarStr, reverseStr,
   ))
   basejsPattern = regexp.MustCompile(`(/s/player/\w+/player_ias.vflset/\w+/base.js)`)
   reverseRegexp = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, reverseStr))
   spliceRegexp  = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, spliceStr))
   swapRegexp    = regexp.MustCompile(fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, swapStr))
   swapStr = fmt.Sprint(
      `:function\(a,b\)\{var c=a\[0\];a\[0\]=a\[b(?:%a\.length)?\];`,
      `a\[b(?:%a\.length)?\]=c(?:;return a)?\}`,
   )
)

func parseDecipherOps(videoID string) ([]decipherOperation, error) {
   embedBody, err := readAll("https://youtube.com/embed/" + videoID)
   if err != nil { return nil, err }
   // example: /s/player/f676c671/player_ias.vflset/en_US/base.js
   escapedBasejsURL := string(basejsPattern.Find(embedBody))
   if escapedBasejsURL == "" {
   log.Println("playerConfig:", string(embedBody))
      return nil, errors.New("unable to find basejs URL in playerConfig")
   }
   baseJsBody, err := readAll("https://youtube.com" + escapedBasejsURL)
   if err != nil { return nil, err }
   objResult := actionsObjRegexp.FindSubmatch(baseJsBody)
   funcResult := actionsFuncRegexp.FindSubmatch(baseJsBody)
   if len(objResult) < 3 || len(funcResult) < 2 {
      return nil, fmt.Errorf(
         "error parsing signature tokens (#obj=%d, #func=%d)",
         len(objResult), len(funcResult),
      )
   }
   var (
      funcBody = funcResult[1]
      obj = objResult[1]
      objBody = objResult[2]
      reverseKey, spliceKey, swapKey string
   )
   if result := reverseRegexp.FindSubmatch(objBody); len(result) > 1 {
      reverseKey = string(result[1])
   }
   if result := spliceRegexp.FindSubmatch(objBody); len(result) > 1 {
      spliceKey = string(result[1])
   }
   if result := swapRegexp.FindSubmatch(objBody); len(result) > 1 {
      swapKey = string(result[1])
   }
   regex, err := regexp.Compile(fmt.Sprintf(
      "(?:a=)?%s\\.(%s|%s|%s)\\(a,(\\d+)\\)",
      obj, reverseKey, spliceKey, swapKey,
   ))
   if err != nil { return nil, err }
   var ops []decipherOperation
   for _, s := range regex.FindAllSubmatch(funcBody, -1) {
      switch string(s[1]) {
      case swapKey:
         arg, _ := strconv.Atoi(string(s[2]))
         ops = append(ops, newSwapFunc(arg))
      case spliceKey:
         arg, _ := strconv.Atoi(string(s[2]))
         ops = append(ops, newSpliceFunc(arg))
      case reverseKey:
         ops = append(ops, reverseFunc)
      }
   }
   return ops, nil
}

func readAll(addr string) ([]byte, error) {
   println("Get", addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

func reverseFunc(bs []byte) []byte {
   sort.SliceStable(bs, func(d, e int) bool { return true })
   return bs
}

type decipherOperation func([]byte) []byte

func newSpliceFunc(pos int) decipherOperation {
   return func(bs []byte) []byte {
      return bs[pos:]
   }
}

func newSwapFunc(arg int) decipherOperation {
   return func(bs []byte) []byte {
      pos := arg % len(bs)
      bs[0], bs[pos] = bs[pos], bs[0]
      return bs
   }
}

type simpleCache struct {
   expiredAt time.Time
   operations []decipherOperation
   videoID string
}

// get cache  when it has same video id and not expired
func (s simpleCache) get(videoID string) []decipherOperation {
   if videoID == s.videoID && s.expiredAt.After(time.Now()) {
      operations := make([]decipherOperation, len(s.operations))
      copy(operations, s.operations)
      return operations
   }
   return nil
}

// set cache with default expiration
func (s *simpleCache) set(videoID string, operations []decipherOperation) {
   defaultCacheExpiration := time.Minute * time.Duration(5)
   s.videoID = videoID
   s.operations = make([]decipherOperation, len(operations))
   copy(s.operations, operations)
   s.expiredAt = time.Now().Add(defaultCacheExpiration)
}
