package instagram

type ImageVersion struct {
   Candidates []struct {
      Width int
      Height int
      URL string
   }
}

type VideoVersion struct {
   Type int
   Width int
   Height int
   URL string
}

type dashManifest struct {
   Period struct {
      AdaptationSet []struct { // one video one audio
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}

type Item struct {
   Caption struct {
      Text string
   }
   Carousel_Media []struct {
      Media_Type int
      Video_DASH_Manifest string
      Video_Versions []VideoVersion
      Image_Versions2 ImageVersion
   }
   Image_Versions2 ImageVersion
   Media_Type int
   Taken_At int64
   User struct {
      Username string
   }
   Video_DASH_Manifest string
   Video_Versions []VideoVersion
}


