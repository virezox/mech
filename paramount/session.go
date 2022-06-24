package paramount

import (
   "bytes"
   "encoding/json"
   "github.com/89z/mech/widevine"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (s Session) Content(c widevine.Client) (*widevine.Content, error) {
   mod, err := c.Module()
   if err != nil {
      return nil, err
   }
   buf, err := mod.Marshal()
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
   keys, err := mod.Unmarshal(buf)
   if err != nil {
      return nil, err
   }
   return keys.Content(), nil
}

type Session struct {
   URL string
   LS_Session string
}

func New_Session(content_id string) (*Session, error) {
   token, err := new_token()
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("https://www.paramountplus.com/apps-api/v3.0/androidphone")
   buf.WriteString("/irdeto-control/anonymous-session-token.json")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "at=" + url.QueryEscape(token)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sess := new(Session)
   if err := json.NewDecoder(res.Body).Decode(sess); err != nil {
      return nil, err
   }
   sess.URL += content_id
   return sess, nil
}
