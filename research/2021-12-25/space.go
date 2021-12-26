package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
)

const graphQL =
   "https://twitter.com/i/api/graphql/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById"

type Space struct {
   Data struct {
      AudioSpace struct {
         Metadata struct {
            Media_Key string
         }
      }
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

func (g Guest) Space(id string) (*Space, error) {
   req, err := http.NewRequest("GET", graphQL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   sReq := spaceRequest{ID: id}
   buf, err := json.Marshal(sReq)
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
