package pbs

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func NewWidget(addr *url.URL) (*Widget, error) {
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Cookie", "pbsol.station=KERA")
   LogLevel.Dump(req)
   // this can redirect
   res, err := new(http.Client).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"availability"`)
   scan.Scan()
   wid := new(Widget)
   if err := scan.Decode(wid); err != nil {
      return nil, err
   }
   return wid, nil
}

type Widget struct {
   Slug string
   Program struct {
      Title string
   }
   Title string
   Duration int64
   Encodings []string
}

var LogLevel format.LogLevel

func (w Widget) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "Slug:", w.Slug)
   fmt.Fprintln(f, "Program:", w.Program.Title)
   fmt.Fprintln(f, "Title:", w.Title)
   fmt.Fprint(f, "Duration: ", w.GetDuration())
   if verb == 'a' {
      for _, enc := range w.Encodings {
         fmt.Fprint(f, "\nEncoding: ", enc)
      }
   }
}

func (w Widget) GetDuration() time.Duration {
   return time.Duration(w.Duration) * time.Second
}

func (w Widget) HLS() string {
   for _, enc := range w.Encodings {
      return enc
   }
   return ""
}

func (w Widget) Widget() (*Widget, error) {
   return &w, nil
}

type Widgeter interface {
   Widget() (*Widget, error)
}

func NewWidgeter(addr string) (Widgeter, error) {
   parse, err := url.Parse(addr)
   if err != nil {
      return nil, err
   }
   hasPrefix := func(prefix string) bool {
      return strings.HasPrefix(parse.Path, prefix)
   }
   switch {
   case hasPrefix("/wgbh/frontline/"):
      return NewFrontline(addr)
   case hasPrefix("/wgbh/masterpiece/"):
      return NewMasterpiece(addr)
   case hasPrefix("/wnet/nature/"):
      return NewNature(addr)
   case hasPrefix("/wgbh/nova/"):
      nova, err := NewNova(addr)
      if err != nil {
         return nil, err
      }
      return nova.Asset(), nil
   case hasPrefix("/widget/"):
      return NewWidget(parse)
   case hasPrefix("/video/"):
      return NewVideo(addr)
   }
   return nil, notFound{parse.Path}
}
