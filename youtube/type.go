package youtube

type Android struct {
   PlayabilityStatus struct {
      Reason string
   }
   StreamingData `json:"streamingData"`
   VideoDetails `json:"videoDetails"`
}

type Client struct {
   ClientName string `json:"clientName"`
   ClientVersion string `json:"clientVersion"`
}

type Context struct {
   Client `json:"client"`
}

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

type StreamingData struct {
   AdaptiveFormats Formats
}

type VideoDetails struct {
   Author string
   ShortDescription string
   Title string
   ViewCount int `json:"viewCount,string"`
}

type player struct {
   Context `json:"context"`
   VideoID string `json:"videoId"`
}


type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
}

type Formats []Format

type Image struct {
   Height int
   Frame int
   Format int
   Base string
}

type Images []Image
