package bandcamp

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "regexp"
   "strconv"
)

var Verbose = mech.Verbose

func (d Detail) Tralbum() (*Tralbum, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(d); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", ApiMobile, buf)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

// URL to track_id or album_id, anonymous
func (d *Detail) Head(addr string) error {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return err
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
      d.Band_ID = 1
      d.Tralbum_ID = id
      d.Tralbum_Type = find[1]
      return nil
   }
   return fmt.Errorf("cookies %v", res.Cookies())
}
