package apple

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

type Progress struct {
   *http.Response
   metric []string
   x, xMax int
   y int64
}

func NewProgress(res *http.Response) *Progress {
   var pro Progress
   pro.Response = res
   pro.metric = []string{" B", " kB", " MB", " GB"}
   pro.xMax = 10_000_000
   return &pro
}

func (p *Progress) Read(buf []byte) (int, error) {
   if p.x == 0 {
      bytes := NumberFormat(float64(p.y), p.metric)
      fmt.Println(mech.Percent(p.y, p.ContentLength), bytes)
   }
   num, err := p.Body.Read(buf)
   if err != nil {
      return 0, err
   }
   p.y += int64(num)
   p.x += num
   if p.x >= p.xMax {
      p.x = 0
   }
   return num, nil
}
