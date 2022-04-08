package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "io"
   "net/http"
   "net/url"
   "sort"
   "strings"
   "testing"
)

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

func WorkerGet(addr string, workers int) ([]*bytes.Buffer, error) {
   if workers < 1 {
      return nil, fmt.Errorf("workers out of range")
   }
   res, err := http.Head(addr)
   if err != nil { return nil, err }
   w := worker{
      addr,
      make(chan error),
      make([]*bytes.Buffer, workers),
      res.ContentLength / int64(workers) + 1,
   }
   for job := range w.jobs {
      go w.work(job)
   }
   for range w.jobs {
      err := <-w.err
      if err != nil { return nil, err }
   }
   return w.jobs, nil
}

func TestFormat(t *testing.T) {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.URL.RawQuery = val.Encode()
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
   sort.Sort(youtube.Height{play.StreamingData.AdaptiveFormats, 480})
   for _, form := range play.StreamingData.AdaptiveFormats {
      fmt.Println(form)
   }
}

var body = strings.NewReader(`
{
  "context": {
    "client": {
      "clientName": "TVHTML5_SIMPLY_EMBEDDED_PLAYER",
      "clientVersion": "2.0",
    },
    "thirdParty": {
      "embedUrl": "https://www.youtube.com"
    }
  },
  "videoId": "HsUATh_Nc2U"
}
`)
