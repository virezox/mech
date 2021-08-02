package youtube

import (
   "encoding/json"
   "time"
)

type Item struct {
   CompactVideoRenderer *struct {
      LengthText Text
      Title Text
      VideoID string
   }
}

func (i Item) Duration() (time.Duration, error) {
   l := "00:00:00"
   r := i.CompactVideoRenderer.LengthText.join()
   l = l[:len(l) - len(r)]
   u, err := time.Parse("15:04:05", l + r)
   if err != nil {
      return 0, err
   }
   var t time.Time
   return u.AddDate(1, 0, 0).Sub(t), nil
}

func (i Item) Title() string {
   return i.CompactVideoRenderer.Title.join()
}

func (i Item) VideoID() string {
   return i.CompactVideoRenderer.VideoID
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer *struct {
               Contents []Item
            }
         }
      }
   }
}

func NewSearch(query string) (*Search, error) {
   var body youTubeI
   body.Context.Client = Mweb
   body.Params = "EgIQAQ" // type video
   body.Query = query
   res, err := post(origin + "/youtubei/v1/search", Key, body)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(Search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            if item.CompactVideoRenderer != nil {
               items = append(items, item)
            }
         }
      }
   }
   return items
}

type Text struct {
   Runs []struct {
      Text string
   }
}

func (t Text) join() string {
   var s string
   for _, r := range t.Runs {
      s += r.Text
   }
   return s
}
