package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

const (
   chunk = 10_000_000
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func NewAndroid(id string) (*Android, error) {
   res, err := newPlayer(
      id, ClientAndroid.ClientName, ClientAndroid.ClientVersion,
   ).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Android)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func (a Android) NewFormat(itag int) (*Format, error) {
   for _, format := range a.StreamingData.AdaptiveFormats {
      if format.Itag == itag {
         return &format, nil
      }
   }
   return nil, fmt.Errorf("itag %v", itag)
}

func (f Format) Write(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   var pos int64
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("status %v", res.Status)
      }
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

func newPlayer(id, name, version string) player {
   var p player
   p.Context.Client.ClientName = name
   p.Context.Client.ClientVersion = version
   p.VideoID = id
   return p
}

func (p player) post() (*http.Response, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}


func (r Result) VideoRenderers() []VideoRenderer {
   var vids []VideoRenderer
   for _, sect := range r.Contents.PrimaryContents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.VideoRenderer)
      }
   }
   return vids
}
