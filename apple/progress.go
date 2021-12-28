package apple

import (
   "fmt"
   "github.com/89z/format/measure"
   "net/http"
)

type Progress struct {
   *http.Response
   callback func(int64, int64)
   x int
   y int64
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.x == 0 {
      p.callback(p.y, p.ContentLength)
   }
   num, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.y += int64(num)
   p.x += num
   if p.x >= 10_000_000 {
      p.x = 0
   }
   return num, nil
}

// Read method has pointer receiver
func NewProgress(res *http.Response) *Progress {
   var pro Progress
   pro.Response = res
   pro.callback = func(num, den int64) {
      percent := measure.Percent(num, den)
      size := measure.Size.FormatInt(num)
      fmt.Println(percent, size)
   }
   return &pro
}
