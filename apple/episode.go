package apple

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

type Episode struct {
   Data struct {
      Playables map[string]struct {
         Assets Asset
      }
   }
}

func NewEpisode(contentID string) (*Episode, error) {
   req, err := http.NewRequest(
      "GET", "https://tv.apple.com/api/uts/v3/episodes/" + contentID, nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "caller": {"web"},
      "locale": {"en-US"},
      "pfm": {"web"},
      "sf": {strconv.Itoa(sf_max)},
      "v": {strconv.Itoa(v_max)},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   epi := new(Episode)
   if err := json.NewDecoder(res.Body).Decode(epi); err != nil {
      return nil, err
   }
   return epi, nil
}

func (e Episode) Asset() *Asset {
   for _, play := range e.Data.Playables {
      return &play.Assets
   }
   return nil
}
