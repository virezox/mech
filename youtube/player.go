package youtube

import (
   "strconv"
   "strings"
   "time"
)

func (p Player) MarshalText() ([]byte, error) {
   var buf []byte
   buf = append(buf, p.PlayabilityStatus.String()...)
   buf = append(buf, "\nVideo ID: "...)
   buf = append(buf, p.VideoDetails.VideoId...)
   buf = append(buf, "\nDuration: "...)
   buf = append(buf, p.Duration().String()...)
   buf = append(buf, "\nView Count: "...)
   buf = strconv.AppendInt(buf, p.VideoDetails.ViewCount, 10)
   buf = append(buf, "\nAuthor: "...)
   buf = append(buf, p.VideoDetails.Author...)
   buf = append(buf, "\nTitle: "...)
   buf = append(buf, p.VideoDetails.Title...)
   if p.PublishDate() != "" {
      buf = append(buf, "\nPublish Date: "...)
      buf = append(buf, p.PublishDate()...)
   }
   buf = append(buf, '\n')
   for _, form := range p.StreamingData.AdaptiveFormats {
      t, err := form.MarshalText()
      if err != nil {
         return nil, err
      }
      buf = append(buf, t...)
   }
   return buf, nil
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

func (p Status) String() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.Status)
   if p.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.Reason)
   }
   return buf.String()
}
