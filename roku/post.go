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

func (self Playback) Request_URL() string {
   return self.DRM.Widevine.LicenseServer
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
