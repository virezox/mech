package instagram

import (
   "github.com/89z/format"
)

type EdgeMedia struct {
   Edges []struct {
      Node struct {
         Text string
      }
   }
}

type GraphMedia struct {
   Display_URL string
   Edge_Media_To_Caption EdgeMedia
   Edge_Media_To_Parent_Comment EdgeMedia
   Edge_Sidecar_To_Children struct {
      Edges []struct {
         Node struct {
            Display_URL string
            Video_URL string
         }
      }
   }
   Owner struct {
      Username string
   }
   Taken_At_Timestamp int64
   Video_URL string
}

var LogLevel format.LogLevel

type User struct {
   Edge_Followed_By struct {
      Count int64
   }
   Edge_Follow struct {
      Count int64
   }
}

type errorString string

type Item struct {
   Caption struct {
      Text string
   }
   Carousel_Media []Media
   Media
   Taken_At int64
   User struct {
      Username string
   }
}

type Login struct {
   Authorization string
}

type Media struct {
   Image_Versions2 struct {
      Candidates []struct {
         Width int
         Height int
         URL string
      }
   }
   Media_Type int
   Video_DASH_Manifest string
   Video_Versions []struct {
      Type int
      Width int
      Height int
      URL string
   }
}

// I noticed that even with the posts that have `video_dash_manifest`, you have
// to request with a correct User-Agent. If you use wrong agent, you will get a
// normal response, but the `video_dash_manifest` will be missing.
type UserAgent struct {
   API int64
   Brand string
   Density string
   Device string
   Instagram string
   Model string
   Platform string
   Release int64
   Resolution string
}

var Android = UserAgent{
   API: 99,
   Brand: "brand",
   Density: "density",
   Device: "device",
   Instagram: "222.0.0.15.114",
   Model: "model",
   Platform: "platform",
   Release: 9,
   Resolution: "9999x9999",
}

type mpd struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}
