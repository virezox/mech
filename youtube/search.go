package youtube

type TwoColumnSearchResultsRenderer struct {
   PrimaryContents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer struct {
               Contents []struct {
                  VideoRenderer `json:"videoRenderer"`
               }
            }
         }
      }
   }
}

type VideoRenderer struct {
   VideoID string
}

type Result struct {
   Contents struct {
      TwoColumnSearchResultsRenderer `json:"twoColumnSearchResultsRenderer"`
   }
}

func (r Result) VideoRenderers() []VideoRenderer {
   var vids []VideoRenderer
   for _, sect := range r.Contents.PrimaryContents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.VideoRenderer)
      }
   }
   return vids
}
