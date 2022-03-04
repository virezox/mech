package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
   "time"
)

type AudioSpace struct {
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

func (a AudioSpace) Admins() string {
   var buf strings.Builder
   for i, admin := range a.Participants.Admins {
      if i >= 1 {
         buf.WriteByte(',')
      }
      buf.WriteString(admin.Display_Name)
   }
   return buf.String()
}

func (a AudioSpace) Base() string {
   var buf strings.Builder
   buf.WriteString(a.Admins())
   buf.WriteByte('-')
   buf.WriteString(a.Metadata.Title)
   return format.Clean(buf.String())
}

func (a AudioSpace) Duration() time.Duration {
   dur := a.Metadata.Ended_At - a.Metadata.Started_At
   return time.Duration(dur) * time.Millisecond
}

func (g Guest) AudioSpace(id string) (*AudioSpace, error) {
   var str strings.Builder
   str.WriteString("https://twitter.com")
   str.WriteString("/i/api/graphql/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById")
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
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
   var space struct {
      Data struct {
         AudioSpace AudioSpace
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&space); err != nil {
      return nil, err
   }
   return &space.Data.AudioSpace, nil
}

func (g Guest) Stream(space *AudioSpace) (*Stream, error) {
   var str strings.Builder
   str.WriteString("https://twitter.com/i/api/1.1/live_video_stream/status/")
   str.WriteString(space.Metadata.Media_Key)
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
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
