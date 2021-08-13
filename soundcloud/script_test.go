package soundcloud

import (
   "bytes"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "io"
   "testing"
   "time"
)

/*
asset 49 consistently works:
https://a-v2.sndcdn.com/assets/49-4b976e4f.js
content-length: 1393972

but asset 2 is smaller:
https://a-v2.sndcdn.com/assets/2-b0e52b4d.js
content-length: 922378

does asset 2 consistently work? yes:
*/
func TestScript(t *testing.T) {
   addr := "https://soundcloud.com"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   scan := mech.NewScanner(res.Body)
   for scan.ScanAttr("crossorigin", "") {
      src := scan.Attr("src")
      res, err := http.Get(src)
      if err != nil {
         t.Fatal(err)
      }
      defer res.Body.Close()
      body, err := io.ReadAll(res.Body)
      if err != nil {
         t.Fatal(err)
      }
      if bytes.Contains(body, []byte(`client_id:"`)) {
         fmt.Println("pass", src)
      } else if bytes.Contains(body, []byte("?client_id=")) {
         //?client_id=fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj&
         fmt.Println("pass", src)
      } else {
         fmt.Println("fail", src)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
