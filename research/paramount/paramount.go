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

func (c *decryptionModule) getLicenseKeys(licenseRequest, licenseResponse []byte) ([]licenseKey, error) {
   license, err := protobuf.Unmarshal(licenseResponse)
   if err != nil {
      return nil, err
   }
   requestParsed, err := protobuf.Unmarshal(licenseRequest)
   if err != nil {
      return nil, err
   }
   requestMsg := requestParsed.Get(2).Marshal()
   cipherText, err := license.GetBytes(4)
   if err != nil {
      return nil, err
   }
   sessionKey, err := rsa.DecryptOAEP(
      sha1.New(), nil, c.privateKey, cipherText, nil,
   )
   if err != nil {
      return nil, err
   }
   hash, err := cmac.New(aes.NewCipher, sessionKey)
   if err != nil {
      return nil, err
   }
   var key []byte
   key = append(key, 1)
   key = append(key, "ENCRYPTION"...)
   key = append(key, 0)
   key = append(key, requestMsg...)
   key = append(key, 0, 0, 0, 0x80)
   hash.Write(key)
   block, err := aes.NewCipher(hash.Sum(nil))
   if err != nil {
      return nil, err
   }
   var keys []licenseKey
   for _, con := range license.Get(2).GetMessages(3) {
      iv, err := con.GetBytes(2)
      if err != nil {
         return nil, err
      }
      key, err := con.GetBytes(3)
      if err != nil {
         return nil, err
      }
      typ, err := con.GetVarint(4)
      if err != nil {
         return nil, err
      }
      decrypter := cipher.NewCBCDecrypter(block, iv)
      decryptedKey := make([]byte, len(key))
      decrypter.CryptBlocks(decryptedKey, key)
      keys = append(keys, licenseKey{
         Type2: uint64(typ),
         Value: unpad(decryptedKey),
      })
   }
   return keys, nil
}

type licenseKey struct {
   Type2  uint64
   Value []byte
}

type decryptionModule struct {
   clientID   []byte
   privateKey *rsa.PrivateKey
   cencHeader struct {
      KeyId []byte "2"
   }
}

// Generates the license request data.  This is sent to the license server via
// HTTP POST and the server in turn returns the license response.
func (c *decryptionModule) getLicenseRequest() ([]byte, error) {
   msg := protobuf.Message{
      1: protobuf.Bytes(c.clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(c.cencHeader.KeyId),
            },
         },
      },
   }
   hash := sha1.Sum(msg.Marshal())
   signature, err := rsa.SignPSS(
      nopSource{},
      c.privateKey,
      crypto.SHA1,
      hash[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   licenseRequest := protobuf.Message{
      2: msg,
      3: protobuf.Bytes(signature),
   }
   return licenseRequest.Marshal(), nil
}

// Creates a new CDM object with the specified device information.
func newCDM(privateKey, clientID, initData []byte) (*decryptionModule, error) {
   block, _ := pem.Decode(privateKey)
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
      return nil, errors.New(res.Status)
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
