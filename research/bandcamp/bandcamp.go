package bandcamp

import (
   "fmt"
   "github.com/89z/format"
   "net/http"
   "strings"
   "text/scanner"
)

var LogLevel format.LogLevel

type session struct {
   Type string
   ID int
}

func newSession(addr string) (*session, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {   
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   var ses session
   for _, cook := range res.Cookies() {
      if cook.Name == "session" {
         strings.Index(cook.Value, "nilZ0")
         fmt.Println(cook.Value)
         // 1%09r%3A%5B%22nilZ0i3543854834x1646019095%22%5D%09t%3A1646019095
         src := strings.NewReader("nilZ0i3543854834x1646014103")
         var scan scanner.Scanner
         scan.Init(src)
         scan.Mode = scanner.ScanInts
         for scan.Scan() != scanner.EOF {
            fmt.Println(scan.TokenText())
         }
      }
   }
   return &ses, nil
}
