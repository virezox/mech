package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)


func (s Search) Post() (*Result, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(s)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search", buf,
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8")
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   r := new(Result)
   if err := json.NewDecoder(res.Body).Decode(r); err != nil {
      return nil, err
   }
   return r, nil
}
