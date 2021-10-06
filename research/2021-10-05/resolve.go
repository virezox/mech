package resolve

import (
   "fmt"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

const NilZ0 = "nilZ0"

type details struct {
   Band_ID int
   Tralbum_ID int
   Tralbum_Type byte
}

func newDetails(s string) (*details, error) {
   r := strings.Index(s, NilZ0)
   if r == -1 {
      return nil, notFound{NilZ0}
   }
   x := strings.IndexByte(s, 'x')
   if x == -1 {
      return nil, notFound{'x'}
   }
   r += len(NilZ0)
   id, err := strconv.Atoi(s[r+1:x])
   if err != nil {
      return nil, err
   }
   return &details{
      Tralbum_Type: s[r], Tralbum_ID: id,
   }, nil
}

func oldDetails(addr string) (*details, error) {
   fmt.Println("HEAD", addr)
   res, err := http.Head(addr)
   if err != nil {
      return nil, err
   }
   for _, cook := range res.Cookies() {
      if cook.Name != "session" {
         continue
      }
      val, err := url.PathUnescape(cook.Value)
      if err != nil {
         return nil, err
      }
      var d details
      if _, err := fmt.Sscanf(
         val, "1\tr:[\"nilZ0%c%vx%v\"]", &d.Tralbum_Type, &d.Tralbum_ID,
      ); err != nil {
         fmt.Println(err)
      } else {
         return &d, nil
      }
   }
   return nil, fmt.Errorf("cookies %v", res.Cookies())
}

type notFound struct {
   value interface{}
}

func (n notFound) Error() string {
   return fmt.Sprintf("%q not found", n.value)
}
