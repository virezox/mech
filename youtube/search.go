package youtube

import (
   "encoding/json"
)

type Item struct {
   CompactVideoRenderer *struct {
      LengthText text
      Title text
      VideoID string
   }
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
   body.Query = query
   par := NewParams()
   fil := NewFilter()
   fil.Type(TypeVideo)
   par.Filter(fil)
   body.Params = par.Encode()
   res, err := post(origin + "/youtubei/v1/search", Key, body)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sea := new(Search)
   if err := json.NewDecoder(res.Body).Decode(sea); err != nil {
      return nil, err
   }
   return sea, nil
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

type text struct {
   Runs []struct {
      Text string
   }
}
