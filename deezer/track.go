package deezer

import (
   "bytes"
   "crypto/aes"
   "encoding/json"
   "fmt"
   "net/http"
)

type Track struct {
   ART_NAME string
   MD5_ORIGIN string
   MEDIA_VERSION string
   SNG_TITLE string
}

// Given a SNG_ID string, make a "deezer.pageTrack" request and return the
// result.
func NewTrack(sngID, arl, sID string) (*Track, error) {
   in, out := map[string]string{"SNG_ID": sngID}, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := http.NewRequest("POST", GatewayWWW, out)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("method", "deezer.pageTrack")
   val.Set("api_version", "1.0")
   val.Set("api_token", arl)
   req.URL.RawQuery = val.Encode()
   req.AddCookie(&http.Cookie{Name: "sid", Value: sID})
   fmt.Println(invert, "POST", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   var page struct {
      Results struct {
         Data Track
      }
   }
   json.NewDecoder(res.Body).Decode(&page)
   return &page.Results.Data, nil
}

// Given SNG_ID and file format, return audio URL.
func (t Track) Source(sngID string, format rune) (string, error) {
   block, err := aes.NewCipher(keyAES)
   if err != nil { return "", err }
   text :=
      t.MD5_ORIGIN +
      "\xa4" +
      string(format) +
      "\xa4" +
      sngID +
      "\xa4" +
      t.MEDIA_VERSION
   var pt bytes.Buffer
   pt.WriteString(md5Hash(text))
   pt.WriteByte(0xA4)
   pt.WriteString(text)
   for pt.Len() % aes.BlockSize > 0 {
      pt.WriteByte(0)
   }
   ciphertext := make([]byte, pt.Len())
   newECBEncrypter(block).CryptBlocks(ciphertext, pt.Bytes())
   return fmt.Sprintf(
      "http://e-cdn-proxy-%c.deezer.com/mobile/1/%x",
      t.MD5_ORIGIN[0], ciphertext,
   ), nil
}
