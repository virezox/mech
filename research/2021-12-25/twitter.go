package twitter

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const root = "https://api.twitter.com/1.1"

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel mech.LogLevel

type Activate struct {
   Guest_Token string
}

func NewActivate() (*Activate, error) {
   req, err := http.NewRequest("POST", root + "/guest/activate.json", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   act := new(Activate)
   if err := json.NewDecoder(res.Body).Decode(act); err != nil {
      return nil, err
   }
   return act, nil
}

func (a Activate) Status(id uint64) (*Status, error) {
   buf := []byte(root)
   buf = append(buf, "/statuses/show/"...)
   buf = strconv.AppendUint(buf, id, 10)
   buf = append(buf, ".json?tweet_mode=extended"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {a.Guest_Token},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stat := new(Status)
   if err := json.NewDecoder(res.Body).Decode(stat); err != nil {
      return nil, err
   }
   return stat, nil
}

type Status struct {
   Entities struct {
      URLs []struct {
         // https://twitter.com/i/spaces/1OdKrBnaEPXKX?s=20
         Expanded_URL string
      }
   }
   Extended_Entities *struct {
      Media []struct {
         Video_Info struct {
            Variants []struct {
               Content_Type string
               URL string
            }
         }
      }
   }
   User struct {
      Name string
   }
}
