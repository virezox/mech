package bandcamp

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "regexp"
   "strconv"
   "time"
)

const ApiMobile = "http://bandcamp.com/api/mobile/24/tralbum_details"

var Heights = map[int]int{
   100: 3,
   124: 8,
   135: 15,
   138: 12,
   150: 7,
   172: 11,
   210: 9,
   300: 4,
   350: 2,
   368: 14,
   380: 13,
   700: 5,
   1200: 10,
   1500: 1,
}

var Verbose = mech.Verbose

func ArtUrl(id, height int) string {
   hID := Heights[height]
   return fmt.Sprintf("http://f4.bcbits.com/img/a%v_%v.jpg", id, hID)
}

// URL to type and ID. Request is anonymous.
func Head(addr string) (byte, int, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return 0, 0, err
   }
   if req.URL.Path == "" {
      req.URL.Path = "/music"
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return 0, 0, err
   }
   reg := regexp.MustCompile(`nilZ0([ait])(\d+)x`)
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

// All fields available with Track and Album
type Tralbum struct {
   Art_ID int
   Release_Date int64
   Title string
   Tracks []struct {
      Streaming_URL struct {
         MP3_128 string `json:"mp3-128"`
      }
   }
   Tralbum_Artist string
}

// ID to Tralbum. Request is anonymous.
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

func (t Tralbum) Unix() time.Time {
   return time.Unix(t.Release_Date, 0)
}
