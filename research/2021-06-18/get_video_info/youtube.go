package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func request(req *http.Request) error {
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   date := []byte("publishDate")
   if ! bytes.Contains(body, date) {
      return fmt.Errorf("missing %q", date)
   }
   return nil
}

func values() error {
   req, err := http.NewRequest(
      "GET", "https://www.youtube.com/get_video_info", nil,
   )
   if err != nil {
      return err
   }
   for _, c := range cs {
      for _, cplayer := range cplayers {
         for _, cver := range cvers {
            for _, el := range els {
               for _, hl := range hls {
                  for _, sts := range stss{
                     val := make(url.Values)
                     val.Set("c", c)
                     if cplayer != "" {
                        val.Set("cplayer", cplayer)
                     }
                     val.Set("cver", cver)
                     if el != "" {
                        val.Set("el", el)
                     }
                     // stackoverflow.com/a/67629882
                     val.Set("eurl", "https://youtube.googleapis.com/v/NMYIVsdGfoo")
                     if hl != "" {
                        val.Set("hl", hl)
                     }
                     // stackoverflow.com/a/67629882
                     val.Set("html5", "1")
                     if sts != "" {
                        val.Set("sts", sts)
                     }
                     val.Set("video_id", "NMYIVsdGfoo")
                     req.URL.RawQuery = val.Encode()
                     err := request(req)
                     fmt.Print(val, err, "\n\n")
                     time.Sleep(100 * time.Millisecond)
                  }
               }
            }
         }
      }
   }
   return nil
}

var cplayers = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/yt-dlp/yt-dlp
   "UNIPLAYER",
}

var cs = []string{
   // stackoverflow.com/a/67629882
   "TVHTML5",
   // github.com/yt-dlp/yt-dlp
   "WEB_REMIX",
}

var cvers = []string{
   // stackoverflow.com/a/67629882
   "6.20180913",
   // github.com/yt-dlp/yt-dlp
   "0.1",
}

var els = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/yt-dlp/yt-dlp
   "detailpage",
   // github.com/Hexer10/youtube_explode_dart
   "embedded",
}

var hls = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/Hexer10/youtube_explode_dart
   "en",
}

var stss = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/Hexer10/youtube_explode_dart
   "18795",
}

func main() {
   err := values()
   if err != nil {
      panic(err)
   }
}
