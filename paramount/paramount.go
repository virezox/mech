package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
)

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

var Log format.Log

func (p Preview) Base() string {
   var buf []byte
   buf = append(buf, p.Title...)
   if p.SeasonNumber >= 1 {
      buf = append(buf, '-')
      buf = strconv.AppendInt(buf, p.SeasonNumber, 10)
      buf = append(buf, '-')
      buf = append(buf, p.EpisodeNumber...)
   }
   return mech.Clean(string(buf))
}

type Preview struct {
   GUID string
   Title string
   SeasonNumber int64 `json:"cbs$SeasonNumber"`
   EpisodeNumber string `json:"cbs$EpisodeNumber"`
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
   bLen := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(bLen); bLen >= 1; bLen-- {
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
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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
   Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   return res.Location()
}
