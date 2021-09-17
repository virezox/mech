package bandcamp

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type Band struct {
   Band_ID json.Number
   URL string
}

func BandGet(id string) (*Band, error) {
   req, err := http.NewRequest("GET", Origin + "/api/band/3/info", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("band_id", id)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   ban := new(Band)
   if err := roundTrip(req, ban); err != nil {
      return nil, err
   }
   return ban, nil
}

func BandPost(id json.Number) (*Band, error) {
   br := bandClient{id, key}
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(br); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", Origin + "/api/band/3/info", buf)
   if err != nil {
      return nil, err
   }
   ban := new(Band)
   if err := roundTrip(req, ban); err != nil {
      return nil, err
   }
   return ban, nil
}

type bandClient struct {
   Band_ID json.Number `json:"band_id"`
   Key string `json:"key"`
}
