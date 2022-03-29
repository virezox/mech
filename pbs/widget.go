package pbs

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

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
   var (
      sep = []byte("\twindow.videoBridge = ")
      wid = new(Widget)
   )
   if err := json.Decode(res.Body, sep, wid); err != nil {
      return nil, err
   }
   return wid, nil
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
      return nova.Episode().Asset(), nil
   case hasPrefix("/widget/"):
      return NewWidget(parse)
   case hasPrefix("/video/"):
      return NewVideo(addr)
   }
   return nil, notFound{parse.Path}
}

type Widget struct {
   Encodings []string
   Slug string
}

func (w Widget) HLS() string {
   for _, enc := range w.Encodings {
      return enc
   }
   return ""
}
