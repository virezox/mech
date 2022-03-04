package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "time"
)

type Space struct {
   Data struct {
      AudioSpace struct {
         Metadata struct {
            Media_Key string
            Started_At int64
            Ended_At int64 `json:"ended_at,string"`
            Title string
         }
         Participants struct {
            Admins []struct {
               Display_Name string
            }
         }
      }
   }
}

func NewSpace(guest *Guest, id string) (*Space, error) {
   var addr strings.Builder
   addr.WriteString("https://twitter.com")
   addr.WriteString("/i/api/graphql/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById")
   req, err := http.NewRequest("GET", addr.String(), nil)
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

func (s Space) Admins() string {
   var admins strings.Builder
   for i, admin := range s.Data.AudioSpace.Participants.Admins {
      if i >= 1 {
         admins.WriteByte(',')
      }
      admins.WriteString(admin.Display_Name)
   }
   return admins.String()
}

func (s Space) Duration() time.Duration {
   meta := s.Data.AudioSpace.Metadata
   return time.Duration(meta.Ended_At - meta.Started_At) * time.Millisecond
}

func (s Space) Name() string {
   var buf strings.Builder
   buf.WriteString(s.Admins())
   buf.WriteByte('-')
   buf.WriteString(s.Data.AudioSpace.Metadata.Title)
   buf.WriteString(".aac")
   return format.Clean(buf.String())
}

func (s Space) Stream(guest *Guest) (*Stream, error) {
   var addr strings.Builder
   addr.WriteString("https://twitter.com/i/api/1.1/live_video_stream/status/")
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

type Stream struct {
   Source struct {
      Location string // Segment
   }
}

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

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel format.LogLevel

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
