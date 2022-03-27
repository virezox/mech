package youtube

import (
   "fmt"
   "github.com/89z/mech"
   "strings"
   "time"
)

func (p Player) Base() string {
   var buf strings.Builder
   buf.WriteString(p.VideoDetails.Author)
   buf.WriteByte('-')
   buf.WriteString(p.VideoDetails.Title)
   buf.WriteByte('-')
   buf.WriteString(p.VideoDetails.VideoID)
   return mech.Clean(buf.String())
}

type Player struct {
   PlayabilityStatus struct {
      Status string // "OK", "LOGIN_REQUIRED"
      Reason string // "", "Sign in to confirm your age"
   }
   VideoDetails struct {
      VideoID string
      LengthSeconds int64 `json:"lengthSeconds,string"`
      ViewCount int64 `json:"viewCount,string"`
      Author string
      Title string
      ShortDescription string
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string // 2013-06-11
      }
   }
   StreamingData struct {
      AdaptiveFormats Formats
      Formats Formats
   }
}

func (p Player) Date() (time.Time, error) {
   value := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", value)
}

func (p Player) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, p.Status())
   fmt.Fprintln(f, "VideoID:", p.VideoDetails.VideoID)
   fmt.Fprintln(f, "Length:", p.VideoDetails.LengthSeconds)
   fmt.Fprintln(f, "ViewCount:", p.VideoDetails.ViewCount)
   fmt.Fprintln(f, "Author:", p.VideoDetails.Author)
   fmt.Fprintln(f, "Title:", p.VideoDetails.Title)
   date := p.Microformat.PlayerMicroformatRenderer.PublishDate
   if date != "" {
      fmt.Fprintln(f, "Date:", date)
   }
   for _, form := range p.StreamingData.AdaptiveFormats {
      fmt.Fprintln(f)
      form.Format(f, verb)
   }
}

func (p Player) Status() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.PlayabilityStatus.Status)
   if p.PlayabilityStatus.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.PlayabilityStatus.Reason)
   }
   return buf.String()
}
