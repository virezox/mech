package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Preview struct {
   Title string
   SeasonNumber int64 `json:"cbs$SeasonNumber"`
   EpisodeNumber string `json:"cbs$EpisodeNumber"`
   GUID string
}

type Media struct {
   GUID string
   aid int64
   sid string
}

func (p Preview) Base() string {
   var buf []byte
   buf = append(buf, p.Title...)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, p.SeasonNumber, 10)
   buf = append(buf, '-')
   buf = append(buf, p.EpisodeNumber...)
   return mech.Clean(string(buf))
}

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

var LogLevel format.LogLevel

func newToken() (string, error) {
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
   bLen := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(bLen); bLen >= 1; bLen-- {
      b = append(b, high)
   }
   return b
}

func NewMedia(guid string) Media {
   return Media{sid: "dJ5BDC", aid: 2198311517, GUID: guid}
}

func (m Media) DASH() (*url.URL, error) {
   return m.location("MPEG-DASH", "DASH_CENC")
}

func (m Media) HLS() (*url.URL, error) {
   return m.location("MPEG4,M3U", "StreamPack")
}

func (m Media) String() string {
   var buf []byte
   buf = append(buf, "http://link.theplatform.com/s/"...)
   buf = append(buf, m.sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, m.aid, 10)
   buf = append(buf, '/')
   buf = append(buf, m.GUID...)
   return string(buf)
}

func (m Media) Preview() (*Preview, error) {
   req, err := http.NewRequest("GET", m.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "format=preview"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   return res.Location()
}

type Session struct {
   URL string
   LS_Session string
}

func NewSession(contentID string) (*Session, error) {
   token, err := newToken()
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sess := new(Session)
   if err := json.NewDecoder(res.Body).Decode(sess); err != nil {
      return nil, err
   }
   sess.URL += contentID
   return sess, nil
}

func (s Session) Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + s.LS_Session)
   return head
}
