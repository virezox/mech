// Rotten Tomatoes
package tomato

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

type Audience struct {
   Scoreboard struct {
      AudienceScore string
   }
}

type Review struct {
   AggregateRating struct {
      RatingValue string
   }
}

type Score struct {
   mech.Node
}

func NewScore(addr string) (Score, error) {
   res, err := http.Get(addr)
   if err != nil {
      return Score{}, err
   }
   defer res.Body.Close()
   doc, err := mech.NewNode(res.Body)
   if err != nil {
      return Score{}, err
   }
   return Score{doc}, nil
}

func (s Score) NewAudience() (Audience, error) {
   text := s.ByAttr("type", "application/json").Text()
   var a Audience
   if err := json.Unmarshal([]byte(text), &a); err != nil {
      return Audience{}, err
   }
   return a, nil
}

func (s Score) NewReview() (Review, error) {
   text := s.ByAttr("type", "application/ld+json").Text()
   var r Review
   if err := json.Unmarshal([]byte(text), &r); err != nil {
      return Review{}, err
   }
   return r, nil
}
