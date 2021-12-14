package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
)

const root = "https://api.twitter.com/1.1"

const authorization =
   "AAAAAAAAAAAAAAAAAAAAAPYXBAAAAAAACLXUNDekMxqa8h/" +
   "40K4moUkGsoc=TYfbDKbT3jJPCEVnMYqilB28NHfOPqkca3qaAxGfsyKCs0wRbw"

type status struct {
   Exteneded_Entities struct {
      Media []struct {
         Video_Info struct {
            Variants []struct {
               URL string
            }
         }
      }
   }
}

type activate struct {
   Guest_Token string
}

func (a activate) status() (*status, error) {
   req, err := http.NewRequest(
      "GET", root + "/statuses/show/1470124083547418624.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + authorization},
      "X-Guest-Token": {a.Guest_Token},
   }
   req.URL.RawQuery = url.Values{
      "cards_platform":[]string{"Web-12"},
      "include_cards":[]string{"1"},
      "include_reply_count":[]string{"1"},
      "include_user_entities":[]string{"0"},
      "tweet_mode":[]string{"extended"},
   }.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stat := new(status)
   if err := json.NewDecoder(res.Body).Decode(stat); err != nil {
      return nil, err
   }
   return stat, nil
}

func newActivate() (*activate, error) {
   req, err := http.NewRequest("POST", root + "/guest/activate.json", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + authorization)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   act := new(activate)
   if err := json.NewDecoder(res.Body).Decode(act); err != nil {
      return nil, err
   }
   return act, nil
}
