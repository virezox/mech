package twitter

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "net/url"
   "path"
   "strconv"
   "strings"
)

type Stream struct {
   Source struct {
      Location string
   }
}

func (s Stream) Chunks() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", s.Source.Location, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(s.Source.Location)
   return m3u.Decode(res.Body, dir)
}

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

type Variant struct {
   Content_Type string
   URL string
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

type spaceRequest struct {
   ID string `json:"id"`
   IsMetatagsQuery bool `json:"isMetatagsQuery"`
   WithBirdwatchPivots bool `json:"withBirdwatchPivots"`
   WithDownvotePerspective bool `json:"withDownvotePerspective"`
   WithReactionsMetadata bool `json:"withReactionsMetadata"`
   WithReactionsPerspective bool `json:"withReactionsPerspective"`
   WithReplays bool `json:"withReplays"`
   WithScheduledSpaces bool `json:"withScheduledSpaces"`
   WithSuperFollowsTweetFields bool `json:"withSuperFollowsTweetFields"`
   WithSuperFollowsUserFields bool `json:"withSuperFollowsUserFields"`
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

func NewSpace(guest *Guest, id string) (*Space, error) {
   req, err := http.NewRequest(
      "GET", graphQL + "/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {guest.Guest_Token},
   }
   buf, err := json.Marshal(spaceRequest{ID: id})
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "variables=" + url.QueryEscape(string(buf))
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   space := new(Space)
   if err := json.NewDecoder(res.Body).Decode(space); err != nil {
      return nil, err
   }
   return space, nil
}

const (
   API = "https://api.twitter.com/1.1"
   graphQL = "https://twitter.com/i/api/graphql"
   iAPI = "https://twitter.com/i/api/1.1"
)

type Space struct {
   Data struct {
      AudioSpace struct {
         Metadata struct {
            Media_Key string
         }
      }
   }
}

func (s Space) Stream(guest *Guest) (*Stream, error) {
   var addr strings.Builder
   addr.WriteString(iAPI)
   addr.WriteString("/live_video_stream/status/")
   addr.WriteString(s.Data.AudioSpace.Metadata.Media_Key)
   req, err := http.NewRequest("GET", addr.String(), nil)
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
   stream := new(Stream)
   if err := json.NewDecoder(res.Body).Decode(stream); err != nil {
      return nil, err
   }
   return stream, nil
}
