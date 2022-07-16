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

func (p Playback) Request_URL() string {
   return p.DRM.Widevine.LicenseServer
}

func (Playback) Request_Header() http.Header {
   return nil
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
