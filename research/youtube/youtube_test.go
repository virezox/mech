package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "io"
   "net/http"
   "net/url"
   "strings"
   "testing"
)

func TestYouTube(t *testing.T) {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   var play youtube.Player
   if err := json.NewDecoder(res.Body).Decode(&play); err != nil {
      t.Fatal(err)
   }
   for _, form := range play.StreamingData.AdaptiveFormats {
      if form.ContentLength >= 1 && form.ContentLength <= 9_999_999 {
         fmt.Println(form)
         break
      }
   }
}

type worker struct {
   addr string
   err chan error
   jobs []*bytes.Buffer
   step int64
}

func (w worker) work(job int) {
   req, err := http.NewRequest("GET", w.addr, nil)
   if err != nil {
      w.err <- err
      return
   }
   req.Header.Set("Range", fmt.Sprintf("bytes=%v-", int64(job) * w.step))
   res, err := new(http.Client).Do(req)
   if err != nil {
      w.err <- err
      return
   }
   defer res.Body.Close()
   w.jobs[job] = new(bytes.Buffer)
   io.CopyN(w.jobs[job], res.Body, w.step)
   w.err <- nil
}

func Get(addr string, length int64, workers int) ([]*bytes.Buffer, error) {
   work := worker{
      addr,
      make(chan error),
      make([]*bytes.Buffer, workers),
      length / int64(workers) + 1,
   }
   for job := range work.jobs {
      go work.work(job)
   }
   for range work.jobs {
      err := <-work.err
      if err != nil {
         return nil, err
      }
   }
   return work.jobs, nil
}

var body = strings.NewReader(`{
  "context": {
    "client": {
      "clientName": "TVHTML5_SIMPLY_EMBEDDED_PLAYER",
      "clientVersion": "2.0"
    }
  },
  "videoId": "Tq92D6wQ1mg"
}
`)
