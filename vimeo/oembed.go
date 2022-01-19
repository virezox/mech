package vimeo

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}

func Parse(id string) (uint64, error) {
   return strconv.ParseUint(id, 10, 64)
}

type Oembed struct {
   Title string
   Upload_Date string
   Thumbnail_URL string
}

func NewOembed(id uint64) (*Oembed, error) {
   req, err := http.NewRequest("GET", "https://vimeo.com/api/oembed.json", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "url=//vimeo.com/" + strconv.FormatUint(id, 10)
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   embed := new(Oembed)
   if err := json.NewDecoder(res.Body).Decode(embed); err != nil {
      return nil, err
   }
   return embed, nil
}
