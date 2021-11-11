package progress

import (
   "fmt"
   "io"
   "net/http"
)

type progress struct {
   io.Reader
   begin time.Time
   pos int64
}

func (p *progress) Read(buf []byte) (int, error) {
   n, err := p.Reader.Read(buf)
   if err != nil {
      fmt.Println()
   } else {
      p.read += n
      fmt.Printf("\rRead %9v", p.read)
   }
   return n, err
}

func main() {
   r, err := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if err != nil {
      panic(err)
   }
   defer r.Body.Close()
   io.ReadAll(&progress{Reader: r.Body})
}
