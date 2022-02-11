package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
   "text/scanner"
   "time"
)

var LogLevel format.LogLevel

func scanInt(buf *scanner.Scanner) (int64, error) {
   for {
      switch buf.Scan() {
      case scanner.Int:
         return strconv.ParseInt(buf.TokenText(), 10, 64)
      case scanner.EOF:
         return 0, nil
      }
   }
}

type Clip struct {
   ID, UnlistedHash int64
}

func NewClip(address string) (*Clip, error) {
   var (
      clipPage Clip
      err error
   )
   buf := new(scanner.Scanner)
   buf.Init(strings.NewReader(address))
   buf.Mode = scanner.ScanInts
   clipPage.ID, err = scanInt(buf)
   if err != nil {
      return nil, err
   }
   clipPage.UnlistedHash, err = scanInt(buf)
   if err != nil {
      return nil, err
   }
   return &clipPage, nil
}

type JsonWeb struct {
   Token string
}

func NewJsonWeb() (*JsonWeb, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Requested-With", "XMLHttpRequest")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(JsonWeb)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

func (w JsonWeb) Video(clip *Clip) (*Video, error) {
   buf := []byte("https://api.vimeo.com/videos/")
   buf = strconv.AppendInt(buf, clip.ID, 10)
   if clip.UnlistedHash >= 1 {
      buf = append(buf, ':')
      buf = strconv.AppendInt(buf, clip.UnlistedHash, 10)
   }
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "JWT " + w.Token)
   req.URL.RawQuery = "fields=duration,download"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

type Video struct {
   Duration int64
   Download []struct {
      Public_Name string
      Size int64
      Link string
   }
}

func (v Video) String() string {
   buf := []byte("Duration: ")
   buf = append(buf, v.Time().String()...)
   buf = append(buf, "\nDownloads:"...)
   for _, dow := range v.Download {
      buf = append(buf, "\nName:"...)
      buf = append(buf, dow.Public_Name...)
      buf = append(buf, " Size:"...)
      buf = append(buf, format.Size.GetInt64(dow.Size)...)
      if dow.Link != "" {
         buf = append(buf, " Link:"...)
         buf = append(buf, dow.Link...)
      }
   }
   return string(buf)
}

func (v Video) Time() time.Duration {
   return time.Duration(v.Duration) * time.Second
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
