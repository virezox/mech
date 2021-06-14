package deezer

import (
   "bytes"
   "crypto/aes"
   "crypto/cipher"
   "crypto/md5"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "golang.org/x/crypto/blowfish"
   "net/http"
)

const (
   GatewayWWW = "http://www.deezer.com/ajax/gw-light.php"
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

const (
   FLAC = '9'
   MP3_320 = '3'
)

var (
   iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
   keyAES = []byte("jo6aey6haid2Teih")
   keyBlowfish = []byte("g4el58wc0zvf9na1")
)

// Given SNG_ID and byte slice, decrypt byte slice in place.
func Decrypt(sngID string, data []byte) error {
   hash := md5Hash(sngID)
   for n := range keyBlowfish {
      keyBlowfish[n] ^= hash[n] ^ hash[n + len(keyBlowfish)]
   }
   block, err := blowfish.NewCipher(keyBlowfish)
   if err != nil {
      return err
   }
   size := 2048
   for pos := 0; len(data) - pos >= size; pos += size {
      if pos / size % 3 == 0 {
         text := data[pos : pos + size]
         cipher.NewCBCDecrypter(block, iv).CryptBlocks(text, text)
      }
   }
   return nil
}

func md5Hash(s string) string {
   sum := md5.Sum([]byte(s))
   return hex.EncodeToString(sum[:])
}

type UserData struct {
   Results struct {
      CheckForm string
      User struct {
         Options struct {
            License_Token string
         }
      }
   }
   SID string
}

func NewUserData(arl string) (UserData, error) {
   req, err := http.NewRequest("GET", GatewayWWW, nil)
   if err != nil {
      return UserData{}, err
   }
   val := req.URL.Query()
   val.Set("api_version", "1.0")
   val.Set("api_token", "")
   val.Set("input", "3")
   val.Set("method", "deezer.getUserData")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Cookie", "arl=" + arl)
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return UserData{}, err
   }
   defer res.Body.Close()
   var user UserData
   for _, c := range res.Cookies() {
      if c.Name == "sid" {
         user.SID = c.Value
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
      return UserData{}, err
   }
   return user, nil
}

type Track struct {
   ART_NAME string
   MD5_ORIGIN string
   MEDIA_VERSION string
   SNG_TITLE string
}

// Given a SNG_ID string, make a "deezer.pageTrack" request and return the
// result.
func NewTrack(sngID, arl, sID string) (Track, error) {
   in, out := map[string]string{"SNG_ID": sngID}, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := http.NewRequest("POST", GatewayWWW, out)
   if err != nil {
      return Track{}, err
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
      return Track{}, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return Track{}, fmt.Errorf("status %v", res.Status)
   }
   var page struct {
      Results struct {
         Data Track
      }
   }
   json.NewDecoder(res.Body).Decode(&page)
   return page.Results.Data, nil
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

type ecbEncrypter struct {
   cipher.Block
}

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
   return ecbEncrypter{b}
}

func (x ecbEncrypter) BlockSize() int {
   return x.Block.BlockSize()
}

func (x ecbEncrypter) CryptBlocks(dst, src []byte) {
   size := x.BlockSize()
   if len(src) % size != 0 {
      panic("crypto/cipher: input not full blocks")
   }
   if len(dst) < len(src) {
      panic("crypto/cipher: output smaller than input")
   }
   for len(src) > 0 {
      x.Encrypt(dst, src)
      src = src[size:]
      dst = dst[size:]
   }
}
