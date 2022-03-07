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
   Video_DASH_Manifest string
   Image_Versions2 ImageVersion
   Video_Versions []VideoVersion
   Carousel_Media []struct {
      Video_DASH_Manifest string
      Image_Versions2 ImageVersion
      Video_Versions []VideoVersion
   }
}
