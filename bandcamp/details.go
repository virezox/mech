package bandcamp

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "strings"
)

const nilZ0 = "nilZ0"

type Details struct {
   Band_ID int
   Tralbum_ID int
   Tralbum_Type byte
}

func NewDetails(addr string) (*Details, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   for _, c := range res.Cookies() {
      if c.Name != "session" {
         continue
      }
      r := strings.Index(c.Value, nilZ0)
      if r == -1 {
         continue
      }
      x := strings.IndexByte(c.Value, 'x')
      if x == -1 {
         continue
      }
      r += len(nilZ0)
      id, err := strconv.Atoi(c.Value[r+1:x])
      if err != nil {
         continue
      }
      return &Details{
         Tralbum_Type: c.Value[r], Tralbum_ID: id,
      }, nil
   }
   return nil, fmt.Errorf("cookies %v", res.Cookies())
}
