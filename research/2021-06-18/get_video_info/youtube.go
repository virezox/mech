package main

import (
   "net/http"
   "os"
)

func main() {
   req, err := http.NewRequest(
      "GET", "https://www.youtube.com/get_video_info", nil,
   )
   if err != nil {
      panic(err)
   }
   val := req.URL.Query()
   val.Set("video_id", "NMYIVsdGfoo")
   val.Set("el", "embedded")
   val.Set("eurl", "https://youtube.googleapis.com/v/NMYIVsdGfoo")
   val.Set("hl", "en")
   val.Set("html5", "1")
   val.Set("c", "TVHTML5")
   val.Set("cver", "6.20180913")
   val.Set("sts", "18795")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   file, err := os.Create("get_video_info")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.ReadFrom(res.Body)
}

send: b'GET /watch?v=UpNXI3_ctAc&bpctr=9999999999&has_verified=1 HTTP/1.1\r\nHost: www.youtube.com\r\nCookie: CONSENT=YES+cb.20210328-17-p0.en+FX+211\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3592.2 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nConnection: close\r\n\r\n'

