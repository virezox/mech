package bandcamp

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "regexp"
   "strconv"
)

type Detail struct {
   Band_ID int `json:"band_id"`
   Tralbum_ID int `json:"tralbum_id,omitempty"`
   Tralbum_Type string `json:"tralbum_type,omitempty"`
}

func TralbumDetail(addr string) (*Detail, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   reg := regexp.MustCompile(`nilZ0([at])(\d+)x`)
   for _, c := range res.Cookies() {
      if c.Name != "session" {
         continue
      }
      find := reg.FindStringSubmatch(c.Value)
      if find == nil {
         continue
      }
      // [nilZ0t2809477874x t 2809477874]
      id, err := strconv.Atoi(find[2])
      if err != nil {
         continue
      }
      return &Detail{
         1, id, find[1],
      }, nil
   }
   return nil, fmt.Errorf("cookies %v", res.Cookies())
}
