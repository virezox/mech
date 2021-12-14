package twitter

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
)

const root = "https://api.twitter.com/1.1"

const bearer =
   "AAAAAAAAAAAAAAAAAAAAAPYXBAAAAAAACLXUNDekMxqa8h/" +
   "40K4moUkGsoc=TYfbDKbT3jJPCEVnMYqilB28NHfOPqkca3qaAxGfsyKCs0wRbw"

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

type URL string

func (u URL) String() string {
   str := string(u)
   addr, err := url.Parse(str)
   if err != nil {
      return str
   }
   addr.RawQuery = ""
   return addr.String()
}

type Status struct {
   Extended_Entities struct {
      Media []struct {
         Video_Info struct {
            Variants []struct {
               Content_Type string
               URL URL
            }
         }
      }
   }
   User struct {
      Name string
   }
}
