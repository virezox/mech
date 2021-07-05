package youtube

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client Client
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         PublishDate string
      }
   }
   PlayabilityStatus struct {
      Reason string
   }
   StreamingData struct {
      AdaptiveFormats Formats
   }
   VideoDetails VideoDetails
}

type PlayerRequest struct {
   Context Context
   VideoID string `json:"videoId"`
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
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

type SearchRequest struct {
   Context Context
   Query string `json:"query"`
}
