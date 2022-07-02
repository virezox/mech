package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/std/http"
   "net/url"
   "strconv"
   "strings"
)

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
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

var Client = http.Default_Client

func (p Preview) Base() string {
   var b []byte
   b = append(b, p.Title...)
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

type Media struct {
   GUID string
   aid int64
   sid string
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

func New_Media(guid string) Media {
   return Media{sid: "dJ5BDC", aid: 2198311517, GUID: guid}
}

func (m Media) DASH() (*url.URL, error) {
   return m.location("MPEG-DASH", "DASH_CENC")
}

func (m Media) HLS() (*url.URL, error) {
   return m.location("MPEG4,M3U", "StreamPack")
}

func (m Media) String() string {
   var b []byte
   b = append(b, "http://link.theplatform.com/s/"...)
   b = append(b, m.sid...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, m.aid, 10)
   b = append(b, '/')
   b = append(b, m.GUID...)
   return string(b)
}

func (m Media) Preview() (*Preview, error) {
   req, err := http.NewRequest("GET", m.String(), nil)
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

func (m Media) location(formats, asset string) (*url.URL, error) {
   req, err := http.NewRequest("HEAD", m.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "assetTypes": {asset},
      "formats": {formats},
   }.Encode()
   res, err := Client.Status(302).Do(req)
   if err != nil {
      return nil, err
   }
   return res.Location()
}
