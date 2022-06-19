package youtube

import (
   "fmt"
   "github.com/89z/mech"
   "strings"
   "time"
)

func (p Player) Duration() time.Duration {
   return time.Duration(p.VideoDetails.LengthSeconds) * time.Second
}

func (p Player) Date() string {
   return p.Microformat.PlayerMicroformatRenderer.PublishDate
}

func (p Player) Time() (time.Time, error) {
   return time.Parse("2006-01-02", p.Date())
}

func (p Player) Base() string {
   var buf strings.Builder
   buf.WriteString(p.VideoDetails.Author)
   buf.WriteByte('-')
   buf.WriteString(p.VideoDetails.Title)
   return mech.Clean(buf.String())
}

type Player struct {
   VideoDetails struct {
      VideoId string
      LengthSeconds int64 `json:"lengthSeconds,string"`
      ViewCount int64 `json:"viewCount,string"`
      Author string
      Title string
      ShortDescription string
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

func (p Player) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, p.PlayabilityStatus)
   fmt.Fprintln(f, "VideoId:", p.VideoDetails.VideoId)
   fmt.Fprintln(f, "Duration:", p.Duration())
   fmt.Fprintln(f, "ViewCount:", p.VideoDetails.ViewCount)
   fmt.Fprintln(f, "Author:", p.VideoDetails.Author)
   fmt.Fprintln(f, "Title:", p.VideoDetails.Title)
   if p.Date() != "" {
      fmt.Fprintln(f, "Date:", p.Date())
   }
   for _, form := range p.StreamingData.AdaptiveFormats {
      fmt.Fprintln(f)
      form.Format(f, verb)
   }
}
