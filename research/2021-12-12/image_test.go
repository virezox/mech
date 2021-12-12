package youtube

import (
   "net/http"
   "os"
   "path"
   "strings"
   "testing"
   "time"
)

const id = "SrHP61Ko7pY"

func TestImage(t *testing.T) {
   for _, img := range images {
      var addr strings.Builder
      addr.WriteString("http://i.ytimg.com/")
      addr.WriteString(img.format.dir)
      addr.WriteByte('/')
      addr.WriteString(id)
      addr.WriteByte('/')
      addr.WriteString(img.base)
      addr.WriteByte('.')
      addr.WriteString(img.format.ext)
      if img.format == webP {
         res, err := http.Get(addr.String())
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         file, err := os.Create(path.Base(addr.String()))
         if err != nil {
            t.Fatal(err)
         }
         file.ReadFrom(res.Body)
         file.Close()
         time.Sleep(99 * time.Millisecond)
      }
   }
}
