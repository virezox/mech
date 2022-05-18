package bearer

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "github.com/89z/format"
   "net/http"
   "net/url"
)

var LogLevel format.LogLevel

func pad(b []byte) []byte {
   bLen := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(bLen); bLen >= 1; bLen-- {
      b = append(b, high)
   }
   return b
}

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

func newToken() (string, error) {
   key, err := hex.DecodeString(aes_key)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   iv := []byte("0123456789ABCDEF")
   var (
      dst []byte
      src []byte
   )
   src = append(src, '|')
   src = append(src, tv_secret...)
   src = pad(src)
   cipher.NewCBCEncrypter(block, iv).CryptBlocks(src, src)
   dst = append(dst, 0, 16)
   dst = append(dst, iv...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

func newBearer() (*http.Response, error) {
   token, err := newToken()
   if err != nil {
      return nil, err
   }
   req := new(http.Request)
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["at"] = []string{token}
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
