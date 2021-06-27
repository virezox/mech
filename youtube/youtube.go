package youtube

var (
   ClientAndroid = Client{"ANDROID", "15.01"}
   ClientMWeb = Client{"MWEB", "2.19700101"}
   ClientWeb = Client{"WEB", "1.19700101"}
)

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
}

type Search struct {
   Context `json:"context"`
   Query string `json:"query"`
}

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

type Result struct {
   Contents struct {
      TwoColumnSearchResultsRenderer `json:"twoColumnSearchResultsRenderer"`
   }
}

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

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   VideoDetails `json:"videoDetails"`
}

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}

type Image struct {
   Height int64
   Frame int64
   Format int64
   Base string
}

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}

type MWeb struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}
