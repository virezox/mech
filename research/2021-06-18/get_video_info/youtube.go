package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
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
   file, err := os.Create("info.txt")
   if err != nil {
      return err
   }
   defer file.Close()
   for _, c := range cs {
      for _, cplayer := range cplayers {
         for _, cver := range cvers {
            for _, el := range els {
               for _, eurl := range eurls {
                  for _, hl := range hls {
                     for _, ps := range pss {
                        for _, sts := range stss {
                           val := make(url.Values)
                           val.Set("c", c)
                           if cplayer != "" {
                              val.Set("cplayer", cplayer)
                           }
                           val.Set("cver", cver)
                           if el != "" {
                              val.Set("el", el)
                           }
                           val.Set("eurl", eurl)
                           if hl != "" {
                              val.Set("hl", hl)
                           }
                           // stackoverflow.com/a/67629882
                           val.Set("html5", "1")
                           if ps != "" {
                              val.Set("ps", ps)
                           }
                           if sts != "" {
                              val.Set("sts", sts)
                           }
                           val.Set("video_id", "NMYIVsdGfoo")
                           fmt.Println(val)
                           req.URL.RawQuery = val.Encode()
                           err := request(req)
                           fmt.Fprintln(file, val, err)
                           time.Sleep(100 * time.Millisecond)
                        }
                     }
                  }
               }
            }
         }
      }
   }
   return nil
}

var cs = []string{
   // stackoverflow.com/a/67629882
   "TVHTML5",
   // github.com/yt-dlp/yt-dlp
   "WEB_REMIX",
}

var cplayers = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/yt-dlp/yt-dlp
   "UNIPLAYER",
}

var cvers = []string{
   // github.com/yt-dlp/yt-dlp
   "0.1",
   // stackoverflow.com/a/67629882
   "6.20180913",
   // github.com/pytube/pytube
   "7.20201028",
}

var els = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/yt-dlp/yt-dlp
   "detailpage",
   // github.com/Hexer10/youtube_explode_dart
   "embedded",
}

var eurls = []string{
   // stackoverflow.com/a/67629882
   "https://youtube.googleapis.com/v/NMYIVsdGfoo",
   // github.com/pytube/pytube
   "https://www.youtube.com/watch?v=NMYIVsdGfoo",
}

var hls = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/Hexer10/youtube_explode_dart
   "en",
   // github.com/pytube/pytube
   "en_US",
}

var pss = []string{
   // stackoverflow.com/a/67629882
   "",
   // github.com/pytube/pytube
   "default",
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
