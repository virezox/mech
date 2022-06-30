package youtube

import (
   "strconv"
   "strings"
   "time"
)

func (p Player) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, p.PlayabilityStatus.String()...)
   b = append(b, "\nVideo ID: "...)
   b = append(b, p.VideoDetails.VideoId...)
   b = append(b, "\nDuration: "...)
   b = append(b, p.Duration().String()...)
   b = append(b, "\nView Count: "...)
   b = strconv.AppendInt(b, p.VideoDetails.ViewCount, 10)
   b = append(b, "\nAuthor: "...)
   b = append(b, p.VideoDetails.Author...)
   b = append(b, "\nTitle: "...)
   b = append(b, p.VideoDetails.Title...)
   if p.PublishDate() != "" {
      b = append(b, "\nPublish Date: "...)
      b = append(b, p.PublishDate()...)
   }
   b = append(b, '\n')
   for _, form := range p.StreamingData.AdaptiveFormats {
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

func (p Player) Duration() time.Duration {
   return time.Duration(p.VideoDetails.LengthSeconds) * time.Second
}

func (p Player) PublishDate() string {
   return p.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (p Player) Time() (time.Time, error) {
   return time.Parse("2006-01-02", p.PublishDate())
}

func (p Player) Base() string {
   var buf strings.Builder
   buf.WriteString(p.VideoDetails.Author)
   buf.WriteByte('-')
   buf.WriteString(p.VideoDetails.Title)
   return buf.String()
}

func (s Status) String() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(s.Status)
   if s.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(s.Reason)
   }
   return buf.String()
}
