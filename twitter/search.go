package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

type Tweet struct {
   Entities struct {
      URLs []struct {
         Expanded_URL string
      }
   }
}

func (t Tweet) String() string {
   var buf strings.Builder
   for i, addr := range t.Entities.URLs {
      if i >= 1 {
         buf.WriteByte('\n')
      }
      buf.WriteString("URL: ")
      buf.WriteString(addr.Expanded_URL)
   }
   return buf.String()
}

type Search struct {
   GlobalObjects struct {
      Tweets map[int64]Tweet
   }
}

func NewSearch(q string) (*Search, error) {
   req, err := http.NewRequest(
      "GET", "https://api.twitter.com/2/search/adaptive.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   req.URL.RawQuery = url.Values{
      "q": {q},
      // This ensures Spaces Tweets will include Spaces URL
      "tweet_mode": {"extended"},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   search := new(Search)
   if err := json.NewDecoder(res.Body).Decode(search); err != nil {
      return nil, err
   }
   return search, nil
}
