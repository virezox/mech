package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type Search struct {
   GlobalObjects struct {
      Tweets map[int64]struct {
         Entities struct {
            URLs []struct {
               Expanded_URL string
            }
         }
      }
   }
}

func (g Guest) Search(query string) (*Search, error) {
   req, err := http.NewRequest(
      "GET", "https://twitter.com/i/api/2/search/adaptive.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   req.URL.RawQuery = "q=" + url.QueryEscape(query)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sea := new(Search)
   if err := json.NewDecoder(res.Body).Decode(sea); err != nil {
      return nil, err
   }
   return sea, nil
}
