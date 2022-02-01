package instagram

import (
   "encoding/base64"
   "encoding/binary"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

func GetID(shortcode string) (uint64, error) {
   for len(shortcode) <= 11 {
      shortcode = "A" + shortcode
   }
   buf, err := base64.URLEncoding.DecodeString(shortcode)
   if err != nil {
      return 0, err
   }
   return binary.BigEndian.Uint64(buf[1:]), nil
}

type Info struct {
   Media_Type int
   Image_Versions2 struct {
      Candidates []Version
   }
   Video_Versions []Version
}

func (i Info) Version() (*Version, error) {
   var dst Version
   switch i.Media_Type {
   case 1:
      for _, src := range i.Image_Versions2.Candidates {
         if src.Height > dst.Height {
            dst = src
         }
      }
   case 2:
      done := make(map[string]bool)
      var length int64
      for _, src := range i.Video_Versions {
         if !done[src.URL] {
            done[src.URL] = true
            if src.Height > dst.Height {
               dst = src
            } else if src.Height == dst.Height {
               req, err := http.NewRequest("HEAD", src.URL, nil)
               if err != nil {
                  return nil, err
               }
               LogLevel.Dump(req)
               res, err := new(http.Transport).RoundTrip(req)
               if err != nil {
                  return nil, err
               }
               if res.ContentLength > length {
                  dst = src
               }
            }
         }
      }
   }
   return &dst, nil
}

// This gets us image 1241 by 1241, but requires authentication.
func (l Login) MediaItems(id uint64) ([]MediaItem, error) {
   buf := []byte("https://i.instagram.com/api/v1/media/")
   buf = strconv.AppendUint(buf, id, 10)
   buf = append(buf, "/info/"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Authorization},
      "User-Agent": {userAgent},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var info struct {
      Items []MediaItem
   }
   if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
      return nil, err
   }
   return info.Items, nil
}

type MediaItem struct {
   Info
   Carousel_Media []Info
   Like_Count int64
}

func (m MediaItem) Format() (string, error) {
   buf := []byte("Like_Count: ")
   buf = strconv.AppendInt(buf, m.Like_Count, 10)
   buf = append(buf, "\nURLs:"...)
   for _, info := range m.Infos() {
      ver, err := info.Version()
      if err != nil {
         return "", err
      }
      buf = append(buf, "\n- "...)
      buf = append(buf, ver.URL...)
   }
   return string(buf), nil
}

func (m MediaItem) Infos() []Info {
   if m.Media_Type == 8 {
      return m.Carousel_Media
   }
   return []Info{m.Info}
}

type Version struct {
   Width int
   Height int
   URL string
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
