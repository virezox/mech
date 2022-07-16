package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/rosso/http"
   "net/url"
   "strconv"
   "strings"
)

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

var Client = http.Default_Client

func (self Preview) Base() string {
   var b []byte
   b = append(b, self.Title...)
   if self.Season_Number >= 1 {
      b = append(b, '-')
      b = strconv.AppendInt(b, self.Season_Number, 10)
      b = append(b, '-')
      b = append(b, self.Episode_Number...)
   }
   return string(b)
}

type Preview struct {
   Episode_Number string `json:"cbs$EpisodeNumber"`
   GUID string
   Season_Number int64 `json:"cbs$SeasonNumber"`
   Title string
}

func new_token() (string, error) {
   key, err := hex.DecodeString(aes_key)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   var (
      dst []byte
      iv [aes.BlockSize]byte
      src []byte
   )
   src = append(src, '|')
   src = append(src, tv_secret...)
   src = pad(src)
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

func pad(b []byte) []byte {
   b_len := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(b_len); b_len >= 1; b_len-- {
      b = append(b, high)
   }
   return b
}

func New_Session(guid string) (*Session, error) {
   token, err := new_token()
   if err != nil {
      return nil, err
   }
   var b strings.Builder
   b.WriteString("https://www.paramountplus.com/apps-api/v3.0/androidphone")
   b.WriteString("/irdeto-control/anonymous-session-token.json")
   req, err := http.NewRequest("GET", b.String(), nil)
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
   sess.URL += guid
   return sess, nil
}

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

func media(guid string) string {
   var b []byte
   b = append(b, "http://link.theplatform.com/s/"...)
   b = append(b, sid...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, aid, 10)
   b = append(b, '/')
   b = append(b, guid...)
   return string(b)
}

func New_Preview(guid string) (*Preview, error) {
   req, err := http.NewRequest("GET", media(guid), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "format=preview"
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prev := new(Preview)
   if err := json.NewDecoder(res.Body).Decode(prev); err != nil {
      return nil, err
   }
   return prev, nil
}

func DASH(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=DASH_CENC")
   b.WriteByte('&')
   b.WriteString("formats=MPEG-DASH")
   return b.String()
}

func HLS(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=StreamPack")
   b.WriteByte('&')
   b.WriteString("formats=MPEG4,M3U")
   return b.String()
}
