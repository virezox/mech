package tomato

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
)

const AddrSearch = "https://www.rottentomatoes.com/search"

func NewSearch(search string) error {
   req, err := http.NewRequest("GET", AddrSearch, nil)
   if err != nil { return err }
   val := req.URL.Query()
   val.Set("search", search)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil { return err }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   doc, err := mech.Parse(res.Body)
   if err != nil { return err }
   script := doc.ByAttr("id", "movies-json")
   script.Scan()
   fmt.Println(script.Text())
   return nil
}
