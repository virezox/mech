package bandcamp

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "os"
   "strconv"
)

const (
   Origin = "http://bandcamp.com"
   key = "veidihundr"
)

var Verbose bool

func roundTrip(req *http.Request, v interface{}) error {
   if Verbose {
      d, err := httputil.DumpRequest(req, true)
      if err != nil {
         return err
      }
      os.Stdout.Write(d)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(v)
}

type Band struct {
   Discography []struct {
      URL string
   }
}

func (b *Band) Get(id int) error {
   req, err := http.NewRequest("GET", Origin + "/api/band/3/discography", nil)
   if err != nil {
      return err
   }
   q := req.URL.Query()
   q.Set("band_id", strconv.Itoa(id))
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   return roundTrip(req, b)
}

func (b *Band) Post(id int) error {
   br := bandRequest{id, key}
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(br); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", Origin + "/api/band/3/discography", buf)
   if err != nil {
      return err
   }
   return roundTrip(req, b)
}

type BandInfo struct {
   Band_ID int
   URL string
}

func (b *BandInfo) Get(id int) error {
   req, err := http.NewRequest("GET", Origin + "/api/band/3/info", nil)
   if err != nil {
      return err
   }
   q := req.URL.Query()
   q.Set("band_id", strconv.Itoa(id))
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   return roundTrip(req, b)
}

func (b *BandInfo) Post(id int) error {
   br := bandRequest{id, key}
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(br); err != nil {
      return err
   }
   req, err := http.NewRequest("POST", Origin + "/api/band/3/info", buf)
   if err != nil {
      return err
   }
   return roundTrip(req, b)
}

type bandRequest struct {
   Band_ID int `json:"band_id"`
   Key string `json:"key"`
}
