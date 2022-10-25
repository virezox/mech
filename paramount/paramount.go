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

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

var app_secrets = map[string]string{
   "12.0.28": "439ba2e3622c344a",
   "12.0.27": "79b7e56e442e65ed",
   "12.0.26": "f012987182d6f16c",
   "8.1.28": "d0795c0dffebea73",
   "8.1.26": "a75bd3a39bfcbc77",
   "8.1.23": "c0966212aa651e8b",
   "8.1.22": "ddca2f16bfa3d937",
   "8.1.20": "817774cbafb2b797",
   "8.1.18": "1705732089ff4d60",
   "8.1.16": "add603b54be2fc3c",
}

func new_token() (string, error) {
   key, err := hex.DecodeString(secret_key)
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
   src = append(src, app_secrets["12.0.28"]...)
   src = pad(src)
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

var Client = http.Default_Client

func (p Preview) Name() string {
   b := []byte(p.Title)
   if p.Season_Number >= 1 {
      b = append(b, '-')
      b = strconv.AppendInt(b, p.Season_Number, 10)
      b = append(b, '-')
      b = append(b, p.Episode_Number...)
   }
   return string(b)
}

type Preview struct {
   Episode_Number string `json:"cbs$EpisodeNumber"`
   GUID string
   Season_Number int64 `json:"cbs$SeasonNumber"`
   Title string
}

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

func New_Session(guid string) (*Session, error) {
   token, err := new_token()
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("https://www.paramountplus.com/apps-api/v3.0/androidphone")
   buf.WriteString("/irdeto-control/anonymous-session-token.json")
   req, err := http.NewRequest("GET", buf.String(), nil)
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
   b := []byte("http://link.theplatform.com/s/")
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
   var buf strings.Builder
   buf.WriteString(media(guid))
   buf.WriteByte('?')
   buf.WriteString("assetTypes=DASH_CENC")
   buf.WriteByte('&')
   buf.WriteString("formats=MPEG-DASH")
   return buf.String()
}

func HLS(guid string) string {
   var buf strings.Builder
   buf.WriteString(media(guid))
   buf.WriteByte('?')
   buf.WriteString("assetTypes=StreamPack")
   buf.WriteByte('&')
   buf.WriteString("formats=MPEG4,M3U")
   return buf.String()
}
