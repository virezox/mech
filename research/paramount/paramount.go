package paramount

import (
   "bytes"
   "encoding/hex"
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "os"
   "strings"
)

var LogLevel format.LogLevel

func (m Media) kID() ([]byte, error) {
   for _, ada := range m.Period.AdaptationSet {
      for _, con := range ada.ContentProtection {
         con.Default_KID = strings.ReplaceAll(con.Default_KID, "-", "")
         return hex.DecodeString(con.Default_KID)
      }
   }
   return nil, errors.New("no init data found")
}

type Media struct {
   Period struct {
      AdaptationSet []struct {
         ContentProtection []struct {
            Default_KID string `xml:"default_KID,attr"`
         }
      }
   }
}

func NewMedia(name string) (*Media, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   med := new(Media)
   if err := xml.NewDecoder(file).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

func (m Media) Keys(contentID, bearer string) ([]KeyContainer, error) {
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      return nil, err
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      return nil, err
   }
   kID, err := m.kID()
   if err != nil {
      return nil, err
   }
   mod, err := NewModule(privateKey, clientID, kID)
   if err != nil {
      return nil, err
   }
   signedLicense, err := mod.SignedLicenseRequest()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://cbsi.live.ott.irdeto.com/widevine/getlicense",
      bytes.NewReader(signedLicense),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   req.URL.RawQuery = url.Values{
      "AccountId": {"cbsi"},
      "ContentId": {contentID},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return mod.Keys(res.Body)
}

