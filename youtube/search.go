package youtube
import "encoding/json"

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

type CompactVideoRenderer struct {
   VideoID string
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct{
            ItemSectionRenderer struct {
               Contents	[]struct{
                  CompactVideoRenderer CompactVideoRenderer
               }
            }
         }
      }
   }
}

func (i I) Search(query string) (*Search, error) {
   i.Query = query
   res, err := i.post("/youtubei/v1/search")
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

func (s Search) Videos() []CompactVideoRenderer {
   var vids []CompactVideoRenderer
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.CompactVideoRenderer)
      }
   }
   return vids
}
