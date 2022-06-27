package roku

import (
   "net/http"
)

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (p Playback) URL() string {
   return p.DRM.Widevine.LicenseServer
}

func (p Playback) Header() http.Header {
   return nil
}

func (p Playback) Body(buf []byte) []byte {
   return buf
}
