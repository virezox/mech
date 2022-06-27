package paramount

import (
   "bytes"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
)

func (s Session) Content(c widevine.Client) (*widevine.Content, error) {
   module, err := c.Key_ID()
   if err != nil {
      return nil, err
   }
   buf, err := module.Marshal()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", s.URL, bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + s.LS_Session)
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
