package paramount

import (
   "bytes"
   "encoding/base64"
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/url"
   "os"
)

const widevineSchemeIdURI = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"

var LogLevel format.LogLevel

func KeyContainers(contentID, bearer string) ([]KeyContainer, error) {
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      return nil, err
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      return nil, err
   }
   media, err := NewMedia("ignore/stream.mpd")
   if err != nil {
      return nil, err
   }
   pssh, err := media.PSSH()
   if err != nil {
      return nil, err
   }
   mod, err := NewModule(privateKey, clientID, pssh)
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
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return mod.Keys(licenseResponse)
}

type Media struct {
   Period struct {
      AdaptationSet []struct {
         ContentProtection []struct {
            SchemeIdUri string `xml:"schemeIdUri,attr"`
            Pssh        string `xml:"pssh"`
         } `xml:"ContentProtection"`
         Representation []struct {
            ContentProtection []struct {
               SchemeIdUri string `xml:"schemeIdUri,attr"`
               Pssh        struct {
                  Text string `xml:",chardata"`
               } `xml:"pssh"`
            }
         }
      }
   }
}

func NewMedia(name string) (*Media, error) {
   file, err := os.Open("ignore/stream.mpd")
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

func (m Media) PSSH() ([]byte, error) {
   for _, adaptionSet := range m.Period.AdaptationSet {
      for _, protection := range adaptionSet.ContentProtection {
         if protection.SchemeIdUri == widevineSchemeIdURI {
            return base64.StdEncoding.DecodeString(protection.Pssh)
         }
      }
   }
   for _, adaptionSet := range m.Period.AdaptationSet {
      for _, representation := range adaptionSet.Representation {
         for _, protection := range representation.ContentProtection {
            if protection.SchemeIdUri == widevineSchemeIdURI {
               return base64.StdEncoding.DecodeString(protection.Pssh.Text)
            }
         }
      }
   }
   return nil, errors.New("no init data found")
}
