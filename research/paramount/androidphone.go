package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func newToken() (string, error) {
   src := []byte{'|'}
   src = append(src, tv_secret...)
   key, err := hex.DecodeString(aes_key)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   iv := []byte("0123456789ABCDEF")
   src = pad(src)
   cipher.NewCBCEncrypter(block, iv).CryptBlocks(src, src)
   dst := []byte{0, 16}
   dst = append(dst, iv...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

func pad(c []byte) []byte {
   cLen := aes.BlockSize - len(c) % aes.BlockSize
   for high := byte(cLen); cLen >= 1; cLen-- {
      c = append(c, high)
   }
   return c
}

func androidPhone() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "www.paramountplus.com"
   req.URL.Scheme = "https"
   val := make(url.Values)
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json"
   token, err := newToken()
   if err != nil {
      panic(err)
   }
   val["at"] = []string{token}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
