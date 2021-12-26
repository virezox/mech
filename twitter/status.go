package twitter

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
)

const API = "https://api.twitter.com/1.1"

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel mech.LogLevel

type Guest struct {
   Guest_Token string
}

func NewGuest() (*Guest, error) {
   req, err := http.NewRequest("POST", API + "/guest/activate.json", nil)
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
         Video_Info struct {
            Variants []Variant
         }
      }
   }
   User struct {
      Name string
   }
}

func NewStatus(guest *Guest, id uint64) (*Status, error) {
   buf := []byte(API)
   buf = append(buf, "/statuses/show/"...)
   buf = strconv.AppendUint(buf, id, 10)
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

type Variant struct {
   Content_Type string
   URL string
}
