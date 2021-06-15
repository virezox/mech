// Rotten Tomatoes
package tomato

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
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
   println(invert, "Get", reset, addr)
   res, err := http.Get(addr)
   if err != nil {
      return Score{}, err
   }
   defer res.Body.Close()
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return Score{}, err
   }
   return Score{doc}, nil
}

func (s Score) NewAudience() (Audience, error) {
   s.Node = s.ByAttr("type", "application/json")
   s.Scan()
   data := []byte(s.Text())
   var aud Audience
   if err := json.Unmarshal(data, &aud); err != nil {
      return Audience{}, err
   }
   return aud, nil
}

func (s Score) NewReview() (Review, error) {
   s.Node = s.ByAttr("type", "application/ld+json")
   s.Scan()
   data := []byte(s.Text())
   var rev Review
   if err := json.Unmarshal(data, &rev); err != nil {
      return Review{}, err
   }
   return rev, nil
}
