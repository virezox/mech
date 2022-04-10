package youtube

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
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
