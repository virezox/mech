package paramount

import (
   "bytes"
   "crypto/rsa"
   "crypto/x509"
   "encoding/base64"
   "encoding/pem"
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "io"
   "github.com/89z/format/protobuf"
   "net/http"
   "os"
)

// Creates a new CDM object with the specified device information.
func newCDM(privateKey, clientID, initData []byte) (*decryptionModule, error) {
   if len(initData) < 32 {
      return nil, errors.New("initData not long enough")
   }
   block, _ := pem.Decode(privateKey)
   if block == nil || (block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY") {
      return nil, errors.New("failed to decode device private key")
   }
   keyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      // if PCKS1 doesn't work, try PCKS8
      pcks8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
      if err != nil {
         return nil, err
      }
      keyParsed = pcks8Key.(*rsa.PrivateKey)
   }
   var dec decryptionModule
   dec.clientID = clientID
   dec.privateKey = keyParsed
   for i := range dec.sessionID {
      if i == 17 {
         dec.sessionID[i] = '1'
      } else {
         dec.sessionID[i] = '0'
      }
   }
   mes, err := protobuf.Unmarshal(initData[32:])
   if err != nil {
      return nil, err
   }
   dec.cencHeader.KeyId, err = mes.GetBytes(2)
   if err != nil {
      return nil, err
   }
   return &dec, nil
}

func newKeys(contentID, bearer string) ([]licenseKey, error) {
   file, err := os.Open("ignore/stream.mpd")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   initData, err := initDataFromMPD(file)
   if err != nil {
      return nil, err
   }
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      return nil, err
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      return nil, err
   }
   cdm, err := newCDM(privateKey, clientID, initData)
   if err != nil {
      return nil, err
   }
   licenseRequest, err := cdm.getLicenseRequest()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST",
      "https://cbsi.live.ott.irdeto.com/widevine/getlicense?AccountId=cbsi&ContentId=" + contentID,
      bytes.NewReader(licenseRequest),
   )
   if err != nil {
      return nil, err
   }
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return cdm.getLicenseKeys(licenseRequest, licenseResponse)
}

// pks padding is designed so that the value of all the padding bytes is the
// number of padding bytes repeated so to figure out how many padding bytes
// there are we can just look at the value of the last byte. i.e if there are 6
// padding bytes then it will look at like <data> 0x6 0x6 0x6 0x6 0x6 0x6
func unpad(b []byte) []byte {
   if len(b) == 0 {
      return b
   }
   count := int(b[len(b)-1])
   return b[0 : len(b)-count]
}

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}

type mpd struct {
   Period                    struct {
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
            } `xml:"ContentProtection"`
         } `xml:"Representation"`
      } `xml:"AdaptationSet"`
   } `xml:"Period"`
}

var LogLevel format.LogLevel

type errorString string

func (e errorString) Error() string {
   return string(e)
}

// This function retrieves the PSSH/Init Data from a given MPD file reader.
// Example file: https://bitmovin-a.akamaihd.net/content/art-of-motion_drm/mpds/11331.mpd
func initDataFromMPD(r io.Reader) ([]byte, error) {
   var mpdPlaylist mpd
   if err := xml.NewDecoder(r).Decode(&mpdPlaylist); err != nil {
      return nil, err
   }
   const widevineSchemeIdURI = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
   for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
      for _, protection := range adaptionSet.ContentProtection {
         if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh) > 0 {
            return base64.StdEncoding.DecodeString(protection.Pssh)
         }
      }
   }
   for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
      for _, representation := range adaptionSet.Representation {
         for _, protection := range representation.ContentProtection {
            if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh.Text) > 0 {
               return base64.StdEncoding.DecodeString(protection.Pssh.Text)
            }
         }
      }
   }
   return nil, errors.New("no init data found")
}
