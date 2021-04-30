package youtube

type player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
         ViewCount string
         Description struct {
            SimpleText string
         }
         Title struct {
            SimpleText string
         }
      }
   }
}

type Video struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         AvailableCountries []string
         Category           string
         Description struct {
            SimpleText string
         }
         Embed struct {
            IframeURL      string
            FlashURL       string
            Width          int
            Height         int
            FlashSecureURL string
         }
         ExternalChannelID  string
         LengthSeconds      string
         OwnerChannelName   string
         OwnerProfileURL    string
         PublishDate        string
         Title struct {
            SimpleText string
         }
         UploadDate         string
         ViewCount          string
      }
   }
   StreamingData struct {
      DashManifestURL  string
   }
}
