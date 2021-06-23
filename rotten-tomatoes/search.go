package tomato

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

type Search struct {
   Items []Item
}

func NewSearch(search string) (*Search, error) {
   req, err := http.NewRequest("GET", AddrSearch, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("search", search)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return nil, err
   }
   script := doc.ByAttr("id", "movies-json")
   script.Scan()
   data := []byte(script.Text())
   s := new(Search)
   json.Unmarshal(data, s)
   return s, nil
}
