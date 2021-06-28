package youtube

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

////////////////////////////////////////////////////////////////////////////////

type Image struct {
   Height int64
   Frame int64
   Format int64
   Base string
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

type MWeb struct {
   Microformat `json:"microformat"`
   VideoDetails `json:"videoDetails"`
}

type Microformat struct {
   PlayerMicroformatRenderer `json:"playerMicroformatRenderer"`
}

type PlayerMicroformatRenderer struct {
   AvailableCountries []string
   PublishDate string
}

////////////////////////////////////////////////////////////////////////////////

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}

func (c Client) newPlayer(id string) player {
   var p player
   p.Client = c
   p.VideoID = id
   return p
}

////////////////////////////////////////////////////////////////////////////////

type Search struct {
   Context `json:"context"`
   Query string `json:"query"`
}

type Context struct {
   Client `json:"client"`
}

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

var (
   ClientAndroid = Client{"ANDROID", "15.01"}
   ClientMWeb = Client{"MWEB", "2.19700101"}
   ClientWeb = Client{"WEB", "1.19700101"}
)

func NewSearch(query string) Search {
   var s Search
   s.Client = ClientWeb
   s.Query = query
   return s
}
