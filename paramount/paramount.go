package paramount

import (
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

var LogLevel format.LogLevel

func (m Media) Video() (*Video, error) {
   for _, vid := range m.Body.Seq.Video {
      return &vid, nil
   }
   // No AssetType/ProtectionScheme/Format Matches
   return nil, errors.New(m.Body.Seq.Ref.Title)
}

type Media struct {
   Body struct {
      Seq  struct {
         Ref struct {
            Title string `xml:"title,attr"`
         } `xml:"ref"`
         Video []Video `xml:"video"`
      } `xml:"seq"`
   } `xml:"body"`
}

type Video struct {
   Title string `xml:"title,attr"`
   Src string `xml:"src,attr"`
   Param []struct {
      Name string `xml:"name,attr"`
      Value string `xml:"value,attr"`
   } `xml:"param"`
}

func (v Video) Base() string {
   var buf strings.Builder
   buf.WriteString(v.Title)
   buf.WriteByte('-')
   buf.WriteString(v.SeasonNumber())
   buf.WriteByte('-')
   buf.WriteString(v.EpisodeNumber())
   return mech.Clean(buf.String())
}

func (v Video) EpisodeNumber() string {
   for _, par := range v.Param {
      if par.Name == "EpisodeNumber" {
         return par.Value
      }
   }
   return ""
}

func (v Video) SeasonNumber() string {
   for _, par := range v.Param {
      if par.Name == "SeasonNumber" {
         return par.Value
      }
   }
   return ""
}

func HLS(guid string) (*Media, error) {
   return newMedia(guid, "MPEG4,M3U", "StreamPack")
}

func DASH(guid string) (*Media, error) {
   return newMedia(guid, "MPEG-DASH", "DASH_CENC")
}

func newMedia(guid, formats, asset string) (*Media, error) {
   buf := []byte("https://link.theplatform.com/s/")
   buf = append(buf, sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, aid, 10)
   buf = append(buf, '/')
   buf = append(buf, guid...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "assetTypes": {asset},
      "format": {"SMIL"},
      "formats": {formats},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := xml.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}
