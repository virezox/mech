package bandcamp

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "regexp"
   "strconv"
)

const (
   ApiMobile = "http://bandcamp.com/api/mobile/24/tralbum_details"
   ApiUrl = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

var Verbose = mech.Verbose

// URL to track_id or album_id, anonymous
func Head(addr string) (byte, int, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return 0, 0, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return 0, 0, err
   }
   reg := regexp.MustCompile(`nilZ0([at])(\d+)x`)
   for _, c := range res.Cookies() {
      if c.Name == "session" {
         // [nilZ0t2809477874x t 2809477874]
         find := reg.FindStringSubmatch(c.Value)
         if find != nil {
            id, err := strconv.Atoi(find[2])
            if err == nil {
               return find[1][0], id, nil
            }
         }
      }
   }
   return 0, 0, fmt.Errorf("cookies %v", res.Cookies())
}

type Info struct {
   Album_ID int
   Band_ID int
   Track_ID int
}

// URL to track_id, album_id or band_id, key
func NewInfo(addr string) (*Info, error) {
   req, err := http.NewRequest("GET", ApiUrl, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("url", addr)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   inf := new(Info)
   if err := json.NewDecoder(res.Body).Decode(inf); err != nil {
      return nil, err
   }
   return inf, nil
}

type Tralbum struct {
   Bandcamp_URL string
   Tracks []struct {
      Title string
   }
}

func NewTralbum(typ byte, id int) (*Tralbum, error) {
   req, err := http.NewRequest("GET", ApiMobile, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("band_id", "1")
   val.Set("tralbum_type", string(typ))
   val.Set("tralbum_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
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
