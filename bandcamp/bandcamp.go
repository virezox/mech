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

type Detail struct {
   Band_ID int `json:"band_id"`
   Tralbum_ID int `json:"tralbum_id,omitempty"`
   Tralbum_Type string `json:"tralbum_type,omitempty"`
}

const Mobile = "http://bandcamp.com/api/mobile/24/tralbum_details"

var Verbose = mech.Verbose

func (d Detail) Tralbum() (*Tralbum, error) {
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(d); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", Mobile, buf)
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

type Tralbum struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}

////////////////////////////////////////////////////////////////////////////////

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

// URL to ID, key

// URL to ID, anonymous
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
