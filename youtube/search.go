package youtube

import (
   "encoding/json"
   "time"
   p "google.golang.org/protobuf/testing/protopack"
)

var Params = map[string]map[string]p.Message{
   "SORT BY": {
      "Relevance": p.Message{
         p.Tag{1, p.VarintType}, p.Varint(0),
      },
      "Rating": p.Message{
         p.Tag{1, p.VarintType}, p.Varint(1),
      },
      "Upload date": p.Message{
         p.Tag{1, p.VarintType}, p.Varint(2),
      },
      "View count": p.Message{
         p.Tag{1, p.VarintType}, p.Varint(3),
      },
   },
   "UPLOAD DATE": {
      "Last hour": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{1, p.VarintType}, p.Varint(1),
         },
      },
      "Today": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{1, p.VarintType}, p.Varint(2),
         },
      },
      "This week": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{1, p.VarintType}, p.Varint(3),
         },
      },
      "This month": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{1, p.VarintType}, p.Varint(4),
         },
      },
      "This year": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{1, p.VarintType}, p.Varint(5),
         },
      },
   },
   "TYPE": {
      "Video": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{2, p.VarintType}, p.Varint(1),
         },
      },
      "Channel": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{2, p.VarintType}, p.Varint(2),
         },
      },
      "Playlist": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{2, p.VarintType}, p.Varint(3),
         },
      },
      "Movie": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{2, p.VarintType}, p.Varint(4),
         },
      },
   },
   "DURATION": {
      "Under 4 minutes": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{3, p.VarintType}, p.Varint(1),
         },
      },
      "Over 20 minutes": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{3, p.VarintType}, p.Varint(2),
         },
      },
      "4 - 20 minutes": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{3, p.VarintType}, p.Varint(3),
         },
      },
   },
   "FEATURES": {
      "HD": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{4, p.VarintType}, p.Varint(1),
         },
      },
      "Subtitles/CC": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{5, p.VarintType}, p.Varint(1),
         },
      },
      "Creative Commons": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{6, p.VarintType}, p.Varint(1),
         },
      },
      "3D": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{7, p.VarintType}, p.Varint(1),
         },
      },
      "Live": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{8, p.VarintType}, p.Varint(1),
         },
      },
      "Purchased": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{9, p.VarintType}, p.Varint(1),
         },
      },
      "4K": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{14, p.VarintType}, p.Varint(1),
         },
      },
      "360°": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{15, p.VarintType}, p.Varint(1),
         },
      },
      "Location": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{23, p.VarintType}, p.Varint(1),
         },
      },
      "HDR": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{25, p.VarintType}, p.Varint(1),
         },
      },
      "VR180": p.Message{
         p.Tag{2, p.BytesType}, p.LengthPrefix{
            p.Tag{26, p.VarintType}, p.Varint(1),
         },
      },
   },
}

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
