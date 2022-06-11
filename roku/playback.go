package roku
// github.com/89z

import (
   "bytes"
   "errors"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
)

func (p Playback) Key(privateKey, clientID, keyID []byte) ([]byte, error) {
   mod, err := widevine.NewModule(privateKey, clientID, keyID)
   if err != nil {
      return nil, err
   }
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
   keys, err := mod.Unmarshal(out)
   if err != nil {
      return nil, err
   }
   return keys.Content().Key, nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}
