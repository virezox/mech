package roku

import (
   "bytes"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
)

func (p Playback) Content(c widevine.Client) (*widevine.Content, error) {
   module, err := c.Key_ID()
   if err != nil {
      return nil, err
   }
   buf, err := module.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", p.DRM.Widevine.LicenseServer, bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   keys, err := module.Unmarshal(buf)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}
