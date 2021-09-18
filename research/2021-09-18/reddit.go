package reddit

import (
   "encoding/json"
   "fmt"
   "net/http"
)

var Verbose bool

type listing struct {
   Data struct {
      Children []struct {
         Data struct {
            Media struct {
               Reddit_Video struct {
                  DASH_URL string
               }
            }
         }
      }
   }
}

func listings(id string) ([]listing, error) {
   req, err := http.NewRequest(
      "GET", "https://www.reddit.com/comments/" + id + ".json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "Mozilla")
   if Verbose {
      fmt.Println(req.Method, req.URL)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   var lists []listing
   if err := json.NewDecoder(res.Body).Decode(&lists); err != nil {
      return nil, err
   }
   return lists, nil
}
