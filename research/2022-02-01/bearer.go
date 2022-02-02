package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
)

var Android = UserAgent{
   API: 99,
   Brand: "!",
   Device: "!",
   Instagram: "195.0.0.31.123",
   Model: "sdk",
   Platform: "!",
   Release: 9,
   Resolution: "9x9",
}

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
   req.URL.Path = "/api/v1/media/2506147657383710114/info/"
   req.URL.Scheme = "https"
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "User-Agent": {Android.String()},
   }
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
