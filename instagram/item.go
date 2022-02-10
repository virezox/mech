package instagram

import (
   "bytes"
   "encoding/json"
   "encoding/xml"
   "net/http"
   "os"
   "path/filepath"
   "strconv"
   "strings"
)

type Item struct {
   Media
   Carousel_Media []Media
}

func (i Item) Format() (string, error) {
   var buf []byte
   for j, med := range i.Medias() {
      addrs, err := med.URLs()
      if err != nil {
         return "", err
      }
      if j >= 1 {
         buf = append(buf, "\n---\n"...)
      }
      for k, addr := range addrs {
         if k >= 1 {
            buf = append(buf, "\n---\n"...)
         }
         buf = append(buf, addr...)
      }
   }
   return string(buf), nil
}

func (i Item) Medias() []Media {
   if i.Media_Type == 8 {
      return i.Carousel_Media
   }
   return []Media{i.Media}
}

type Login struct {
   Authorization string
}

func NewLogin(username, password string) (*Login, error) {
   buf := bytes.NewBufferString("signed_body=SIGNATURE.")
   sig := map[string]string{
      "device_id": "device_id",
      "enc_password": "#PWD_INSTAGRAM:0:0:" + password,
      "username": username,
   }
   if err := json.NewEncoder(buf).Encode(sig); err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://i.instagram.com/api/v1/accounts/login/", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "User-Agent": {Android.String()},
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

func (l Login) Items(shortcode string) ([]Item, error) {
   var buf strings.Builder
   buf.WriteString("https://www.instagram.com/p/")
   buf.WriteString(shortcode)
   buf.WriteByte('/')
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {l.Authorization},
      "User-Agent": {Android.String()},
   }
   req.URL.RawQuery = "__a=1"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var post struct {
      Items []Item
   }
   if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
      return nil, err
   }
   return post.Items, nil
}

type Media struct {
   Image_Versions2 struct {
      Candidates []struct {
         Width int
         Height int
         URL string
      }
   }
   Media_Type int
   Video_DASH_Manifest string
   Video_Versions []struct {
      Type int
      Width int
      Height int
      URL string
   }
}

func (m Media) URLs() ([]string, error) {
   var addrs []string
   switch m.Media_Type {
   case 1:
      var max int
      for _, can := range m.Image_Versions2.Candidates {
         if can.Height > max {
            addrs = []string{can.URL}
            max = can.Height
         }
      }
   case 2:
      if m.Video_DASH_Manifest != "" {
         var manifest mpd
         err := xml.Unmarshal([]byte(m.Video_DASH_Manifest), &manifest)
         if err != nil {
            return nil, err
         }
         for _, ada := range manifest.Period.AdaptationSet {
            var (
               addr string
               max int
            )
            for _, rep := range ada.Representation {
               if rep.Bandwidth > max {
                  addr = rep.BaseURL
                  max = rep.Bandwidth
               }
            }
            addrs = append(addrs, addr)
         }
      } else {
         // Type:101 Bandwidth:211,754
         // Type:102 Bandwidth:541,145
         // Type:103 Bandwidth:541,145
         var max int
         for _, ver := range m.Video_Versions {
            if ver.Type > max {
               addrs = []string{ver.URL}
               max = ver.Type
            }
         }
      }
   }
   return addrs, nil
}

// I noticed that even with the posts that have `video_dash_manifest`, you have
// to request with a correct User-Agent. If you use wrong agent, you will get a
// normal response, but the `video_dash_manifest` will be missing.
type UserAgent struct {
   API int64
   Brand string
   Density string
   Device string
   Instagram string
   Model string
   Platform string
   Release int64
   Resolution string
}

var Android = UserAgent{
   API: 99,
   Brand: "brand",
   Density: "density",
   Device: "device",
   Instagram: "220.0.0.16.115",
   Model: "model",
   Platform: "platform",
   Release: 9,
   Resolution: "9999x9999",
}

func (u UserAgent) String() string {
   buf := []byte("Instagram ")
   buf = append(buf, u.Instagram...)
   buf = append(buf, " Android ("...)
   buf = strconv.AppendInt(buf, u.API, 10)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.Release, 10)
   buf = append(buf, "; "...)
   buf = append(buf, u.Density...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Resolution...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Brand...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Model...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Device...)
   buf = append(buf, "; "...)
   buf = append(buf, u.Platform...)
   return string(buf)
}

type mpd struct {
   Period struct {
      AdaptationSet []struct {
         Representation []struct {
            Width int `xml:"width,attr"`
            Height int `xml:"height,attr"`
            Bandwidth int `xml:"bandwidth,attr"`
            BaseURL string
         }
      }
   }
}
