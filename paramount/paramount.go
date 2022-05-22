package paramount

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
)

type Address struct {
   sid string
   aid int64
   guid string
}

func NewAddress(guid string) Address {
   return Address{sid: "dJ5BDC", aid: 2198311517, guid: guid}
}

func (a Address) String() string {
   var buf []byte
   buf = append(buf, "http://link.theplatform.com/s/"...)
   buf = append(buf, a.sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, a.aid, 10)
   buf = append(buf, '/')
   buf = append(buf, a.guid...)
   return string(buf)
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

type Preview struct {
   Title string
   SeasonNumber int64 `json:"cbs$SeasonNumber"`
   EpisodeNumber string `json:"cbs$EpisodeNumber"`
}

func (a Address) Preview() (*Preview, error) {
   req, err := http.NewRequest("GET", a.String(), nil)
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

func (a Address) HLS(guid string) (*url.URL, error) {
   return a.location("MPEG4,M3U", "StreamPack")
}

func (a Address) DASH(guid string) (*url.URL, error) {
   return a.location("MPEG-DASH", "DASH_CENC")
}

func (a Address) location(formats, asset string) (*url.URL, error) {
   req, err := http.NewRequest("HEAD", a.String(), nil)
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
