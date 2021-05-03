package youtube

import (
   "bufio"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "regexp"
   "strconv"
   "strings"
)

const API = "https://www.youtube.com/get_video_info"

func readAll(addr string) ([]byte, error) {
   println("Get", addr)
   res, err := http.Get(addr)
   if err != nil { return nil, err }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

type Video struct {
   StreamingData struct {
      AdaptiveFormats []struct {
         Bitrate int
         Height int
         Itag int
         MimeType string
         SignatureCipher string
      }
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   VideoDetails struct {
      ShortDescription string
      Title string
      VideoId string
      ViewCount int `json:",string"`
   }
}

// NewVideo fetches video metadata
func NewVideo(id string) (Video, error) {
   val := make(url.Values)
   val.Set("eurl", API)
   val.Set("video_id", id)
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

func (v Video) Description() string { return v.VideoDetails.ShortDescription }

const split = `.split("");`

// GetStream returns the url for a specific format
func (v Video) GetStream(itag int) (string, error) {
   if len(v.StreamingData.AdaptiveFormats) == 0 {
      return "", errors.New("AdaptiveFormats empty")
   }
   // get cipher text
   cipher, err := v.cipher(itag)
   if err != nil { return "", err }
   query, err := url.ParseQuery(cipher)
   if err != nil { return "", err }
   sig := []byte(query.Get("s"))
   // decrypt
   body, err := readAll("https://www.youtube.com/embed/" + v.VideoDetails.VideoId)
   if err != nil { return "", err }
   player := regexp.MustCompile("/player/([^/]+)/player_").FindSubmatch(body)
   if len(player) < 2 {
      return "", errors.New("unable to find basejs URL in playerConfig")
   }
   base := fmt.Sprintf(
      "https://www.youtube.com/s/player/%s/player_ias.vflset/en_US/base.js",
      player[1],
   )
   res, err := http.Get(base)
   if err != nil { return "", err }
   defer res.Body.Close()
   scan := bufio.NewScanner(res.Body)
   for scan.Scan() {
      if ! strings.Contains(scan.Text(), `.split("");`) { continue }
      for _, match := range regexp.MustCompile(`\d+`).FindAllString(scan.Text(), -1) {
         index, err := strconv.Atoi(match)
         if err != nil { return "", err }
         swap(sig, index)
      }
      return query.Get("url") + "&sig=" + string(sig), nil
   }
   return "", fmt.Errorf("%q not found", split)
}

func swap(sig []byte, index int) {
   c := sig[0]
   sig[0] = sig[index % len(sig)]
   sig[index % len(sig)] = c
}

func (v Video) PublishDate() string {
   return v.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (v Video) Title() string { return v.VideoDetails.Title }

func (v Video) ViewCount() int { return v.VideoDetails.ViewCount }

func (v Video) cipher(itag int) (string, error) {
   for _, format := range v.StreamingData.AdaptiveFormats {
      if format.Itag == itag { return format.SignatureCipher, nil }
   }
   return "", errors.New("itag not found")
}


type cipher struct {
   matches [][]string
   reverse string
   splice string
   swap string
}

func (ci cipher) decrypt(sig []byte) error {
   for _, match := range ci.matches {
      switch match[1] {
      case ci.swap:
         arg, err := strconv.Atoi(match[2])
         if err != nil { return err }
         pos := arg % len(sig)
         sig[0], sig[pos] = sig[pos], sig[0]
      case ci.splice:
         arg, err := strconv.Atoi(match[2])
         if err != nil { return err }
         sig = sig[arg:]
      case ci.reverse:
         for n := len(sig) - 2; n >= 0; n-- {
            sig = append(sig[:n], append(sig[n + 1:], sig[n])...)
         }
      }
   }
   return nil
}

const (
   jsReverse = `:function\(a\)\{(?:return )?a\.reverse\(\)\}`
   jsSplice = `:function\(a,b\)\{a\.splice\(0,b\)\}`
   jsSwap = `:function\(a,b\)\{var c=a\[0\];a\[0\]=a\[b(?:%a\.length)?\];a\[b(?:%a\.length)?\]=c(?:;return a)?\}`
   jsVar = `[a-zA-Z_\$][a-zA-Z_0-9]*`
)
