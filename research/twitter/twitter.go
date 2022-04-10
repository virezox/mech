package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel format.LogLevel

type Search struct {
   GlobalObjects struct {
      Tweets map[int64]struct {
         Entities struct {
            URLs []struct {
               // twitter.com/i/spaces/1ynKOZVDnbkxR
               Expanded_URL string
            }
         }
      }
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
   req.URL.RawQuery = "q=" + url.QueryEscape(q)
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

func (s Search) String() string {
   var buf []byte
   for key, val := range s.GlobalObjects.Tweets {
      if buf != nil {
         buf = append(buf, '\n')
      }
      buf = append(buf, "Tweet: "...)
      buf = strconv.AppendInt(buf, key, 10)
      for _, addr := range val.Entities.URLs {
         buf = append(buf, "\nURL: "...)
         buf = append(buf, addr.Expanded_URL...)
      }
   }
   return string(buf)
}

func pass() (*Search, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/2/search/adaptive.json"
   val := make(url.Values)
   val["q"] = []string{"filter:spaces"}
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   // This ensures Spaces Tweets will include Spaces URL
   val["tweet_mode"] = []string{"extended"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   search := new(Search)
   if err := json.Unmarshal(buf, search); err != nil {
      return nil, err
   }
   return search, nil
}
