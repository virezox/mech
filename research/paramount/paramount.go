package paramount

import (
   "bytes"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/base64"
   "encoding/pem"
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "github.com/89z/format/protobuf"
   "github.com/chmike/cmac-go"
   "io"
   "net/http"
   "os"
)

var LogLevel format.LogLevel

func initDataFromMPD(src io.Reader) ([]byte, error) {
   var mpdPlaylist mpd
   err := xml.NewDecoder(src).Decode(&mpdPlaylist)
   if err != nil {
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

func unpad(buf []byte) []byte {
   if len(buf) >= 1 {
      pad := buf[len(buf)-1]
      if len(buf) >= int(pad) {
         buf = buf[:len(buf)-int(pad)]
      }
   }
   return buf
}

type KeyContainer struct {
   Key []byte
   Type  uint64
}

func KeyContainers(contentID, bearer string) ([]KeyContainer, error) {
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
   mod, err := NewModule(privateKey, clientID, initData)
   if err != nil {
      return nil, err
   }
   licenseRequest, err := mod.LicenseRequest()
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
      return nil, errors.New(res.Status)
   }
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return mod.KeyContainers(licenseRequest, licenseResponse)
}

type Module struct {
   *rsa.PrivateKey
   clientID []byte
   keyID []byte
}

func NewModule(privateKey, clientID, initData []byte) (*Module, error) {
   var mod Module
   // clientID
   mod.clientID = clientID
   // keyID
   widevineCencHeader, err := protobuf.Unmarshal(initData[32:])
   if err != nil {
      return nil, err
   }
   mod.keyID, err = widevineCencHeader.GetBytes(2)
   if err != nil {
      return nil, err
   }
   // PrivateKey
   block, _ := pem.Decode(privateKey)
   mod.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   return &mod, nil
}

func (c *Module) KeyContainers(bRequest, bResponse []byte) ([]KeyContainer, error) {
   // message
   signedLicenseRequest, err := protobuf.Unmarshal(bRequest)
   if err != nil {
      return nil, err
   }
   licenseRequest := signedLicenseRequest.Get(2).Marshal()
   var message []byte
   message = append(message, 1)
   message = append(message, "ENCRYPTION"...)
   message = append(message, 0)
   message = append(message, licenseRequest...)
   message = append(message, 0, 0, 0, 0x80)
   // key
   signedLicense, err := protobuf.Unmarshal(bResponse)
   if err != nil {
      return nil, err
   }
   sessionKey, err := signedLicense.GetBytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(sha1.New(), nil, c.PrivateKey, sessionKey, nil)
   if err != nil {
      return nil, err
   }
   // CMAC
   mac, err := cmac.New(aes.NewCipher, key)
   if err != nil {
      return nil, err
   }
   mac.Write(message)
   block, err := aes.NewCipher(mac.Sum(nil))
   if err != nil {
      return nil, err
   }
   var containers []KeyContainer
   // .Msg.Key
   for _, message := range signedLicense.Get(2).GetMessages(3) {
      iv, err := message.GetBytes(2)
      if err != nil {
         return nil, err
      }
      key, err := message.GetBytes(3)
      if err != nil {
         return nil, err
      }
      typ, err := message.GetVarint(4)
      if err != nil {
         return nil, err
      }
      cipher.NewCBCDecrypter(block, iv).CryptBlocks(key, key)
      var container KeyContainer
      container.Key = unpad(key)
      container.Type = uint64(typ)
      containers = append(containers, container)
   }
   return containers, nil
}

func (c *Module) LicenseRequest() ([]byte, error) {
   // licenseRequest
   licenseRequest := protobuf.Message{
      1: protobuf.Bytes(c.clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(c.keyID),
            },
         },
      },
   }
   // signature
   digest := sha1.Sum(licenseRequest.Marshal())
   signature, err := rsa.SignPSS(
      nopSource{},
      c.PrivateKey,
      crypto.SHA1,
      digest[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signedLicenseRequest := protobuf.Message{
      2: licenseRequest,
      3: protobuf.Bytes(signature),
   }
   return signedLicenseRequest.Marshal(), nil
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

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}
