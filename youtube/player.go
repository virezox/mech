package youtube

import (
   "strconv"
   "strings"
   "time"
)

func (self Player) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, self.PlayabilityStatus.String()...)
   b = append(b, "\nVideo ID: "...)
   b = append(b, self.VideoDetails.VideoId...)
   b = append(b, "\nDuration: "...)
   b = append(b, self.Duration().String()...)
   b = append(b, "\nView Count: "...)
   b = strconv.AppendInt(b, self.VideoDetails.ViewCount, 10)
   b = append(b, "\nAuthor: "...)
   b = append(b, self.VideoDetails.Author...)
   b = append(b, "\nTitle: "...)
   b = append(b, self.VideoDetails.Title...)
   if self.PublishDate() != "" {
      b = append(b, "\nPublish Date: "...)
      b = append(b, self.PublishDate()...)
   }
   b = append(b, '\n')
   for _, form := range self.StreamingData.AdaptiveFormats {
      t, err := form.MarshalText()
      if err != nil {
         return nil, err
      }
      b = append(b, t...)
   }
   return b, nil
}

type Player struct {
   VideoDetails struct {
      Author string
      LengthSeconds int64 `json:"lengthSeconds,string"`
      ShortDescription string
      Title string
      VideoId string
      ViewCount int64 `json:"viewCount,string"`
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   StreamingData struct {
      AdaptiveFormats Formats
   }
   PlayabilityStatus Status
}

type Status struct {
   Status string
   Reason string
}

func (self Player) Duration() time.Duration {
   return time.Duration(self.VideoDetails.LengthSeconds) * time.Second
}

func (self Player) PublishDate() string {
   return self.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (self Player) Time() (time.Time, error) {
   return time.Parse("2006-01-02", self.PublishDate())
}

func (self Player) Base() string {
   var buf strings.Builder
   buf.WriteString(self.VideoDetails.Author)
   buf.WriteByte('-')
   buf.WriteString(self.VideoDetails.Title)
   return buf.String()
}

func (self Status) String() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(self.Status)
   if self.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(self.Reason)
   }
   return buf.String()
}
