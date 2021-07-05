package youtube
import "encoding/json"

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func (r Search) Videos() []CompactVideoRenderer {
   var vids []CompactVideoRenderer
   for _, sect := range r.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.CompactVideoRenderer)
      }
   }
   return vids
}

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

func NewSearch(query string) (*Search, error) {
   res, err := Mweb.query(query).post("/youtubei/v1/search")
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
