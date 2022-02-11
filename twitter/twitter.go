package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
)

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel format.LogLevel

type Guest struct {
   Guest_Token string
}

func NewGuest() (*Guest, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/guest/activate.json", nil,
   )
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
   guest := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(guest); err != nil {
      return nil, err
   }
   return guest, nil
}

type Status struct {
   Extended_Entities struct {
      Media []struct {
         Media_URL string
         Video_Info struct {
            Variants []Variant
         }
      }
   }
   User struct {
      Name string
   }
}

func NewStatus(guest *Guest, id int64) (*Status, error) {
   buf := []byte("https://api.twitter.com/1.1/statuses/show/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, ".json?tweet_mode=extended"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {guest.Guest_Token},
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

func (s Status) Variants() []Variant {
   var varis []Variant
   for _, med := range s.Extended_Entities.Media {
      for _, vari := range med.Video_Info.Variants {
         if vari.Content_Type != "application/x-mpegURL"{
            varis = append(varis, vari)
         }
      }
   }
   return varis
}

type Stream struct {
   Source struct {
      Location string
   }
}

type URL string

func (u URL) String() string {
   address := string(u)
   addr, err := url.Parse(address)
   if err != nil {
      return address
   }
   addr.RawQuery = ""
   return addr.String()
}

type Variant struct {
   Content_Type string
   URL URL
}
