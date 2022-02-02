package main

import (
   "bytes"
   "strconv"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

type UserAgent struct {
   API int64
   Brand string
   Density int64
   Device string
   Instagram string
   Model string
   Platform string
   Release int64
   Resolution string
}

func (u UserAgent) String() string {
   buf := []byte("Instagram ")
   buf = append(buf, u.Instagram...)
   buf = append(buf, " Android ("...)
   buf = strconv.AppendInt(buf, u.API, 10)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.Release, 10)
   buf = append(buf, "; "...)
   buf = strconv.AppendInt(buf, u.Density, 10)
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

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/api/v1/media/2762134734241678695_12223257/info/"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   req.Header["User-Agent"] = []string{
      "Instagram 219.0.0.12.117 Android (24/7.0; 560dpi; 1440x2872; Google/google; Android SDK built for x86; generic_x86; ranchu; en_US; 346138365)",
   }
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(buf, []byte(`"height":1241,`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
