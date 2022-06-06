package roku
// github.com/89z

import (
   "bytes"
   "errors"
   "io"
   "net/http"
   wv "github.com/89z/mech/widevine"
)

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (p Playback) Containers(mod *wv.Module) (wv.Containers, error) {
   in, err := mod.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", p.DRM.Widevine.LicenseServer, bytes.NewReader(in),
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   out, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return mod.Unmarshal(out)
}
