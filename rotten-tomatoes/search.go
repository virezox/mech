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

func NewSearch(search string) (Search, error) {
   req, err := http.NewRequest("GET", AddrSearch, nil)
   if err != nil {
      return Search{}, err
   }
   val := req.URL.Query()
   val.Set("search", search)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Search{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Search{}, fmt.Errorf("status %v", res.Status)
   }
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return Search{}, err
   }
   script := doc.ByAttr("id", "movies-json")
   script.Scan()
   data := []byte(script.Text())
   var s Search
   if err := json.Unmarshal(data, &s); err != nil {
      return Search{}, err
   }
   return s, nil
}
