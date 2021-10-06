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

const (
   ApiAlbum = "http://bandcamp.com/api/album/2/info"
   ApiMobile = "http://bandcamp.com/api/mobile/24/tralbum_details"
   ApiTrack = "http://bandcamp.com/api/track/3/info"
   ApiUrl = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

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

type Album struct {
   Large_Art_URL string // 350 x 350
   Release_Date int64
}

func NewAlbum(id int) (*Album, error) {
   req, err := http.NewRequest("GET", ApiAlbum, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("album_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   alb := new(Album)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}

func (a Album) Unix() time.Time {
   return time.Unix(a.Release_Date, 0)
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

type Track struct {
   Album_ID int
   Streaming_URL string
   Title string
}

func NewTrack(id int) (*Track, error) {
   req, err := http.NewRequest("GET", ApiTrack, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("track_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Track)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
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
