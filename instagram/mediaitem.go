package instagram

import (
   "bytes"
   "encoding/base64"
   "encoding/binary"
   "encoding/json"
   "net/http"
   "os"
   "path/filepath"
   "strconv"
)

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

type Login struct {
   Authorization string
}

func NewLogin(username, password string) (*Login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   sig := map[string]string{
      "device_id": userAgent,
      "enc_password": "#PWD_INSTAGRAM:0:0:" + password,
      "username": username,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/api/v1/accounts/login/", buf)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {userAgent},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var login Login
   login.Authorization = res.Header.Get("Ig-Set-Authorization")
   return &login, nil
}

func OpenLogin(name string) (*Login, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   log := new(Login)
   if err := json.NewDecoder(file).Decode(log); err != nil {
      return nil, err
   }
   return log, nil
}

func (l Login) Create(name string) error {
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   return enc.Encode(l)
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
   buf = append(buf, "\nURLs: "...)
   for i, info := range m.Infos() {
      ver, err := info.Version()
      if err != nil {
         return "", err
      }
      if i >= 1 {
         buf = append(buf, "\n---\n"...)
      }
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
