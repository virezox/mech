package main

import (
   "bufio"
   "fmt"
   "github.com/89z/ja3"
   "github.com/refraction-networking/utls"
   "net"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   pass := os.Getenv("ENCRYPTEDPASS")
   if pass == "" {
      fmt.Println("missing pass")
      return
   }
   val := url.Values{
      "Email": {"srpen6@gmail.com"},
      "sdk_version": {"17"},
      "EncryptedPasswd": {pass},
   }
   fu, err := os.Open("getAllUasJson.json")
   if err != nil {
      panic(err)
   }
   defer fu.Close()
   us, err := ja3.NewUsers(fu)
   if err != nil {
      panic(err)
   }
   us.Sort()
   fh, err := os.Open("getAllHashesJson.json")
   if err != nil {
      panic(err)
   }
   defer fh.Close()
   hs, err := ja3.NewHashes(fh)
   if err != nil {
      panic(err)
   }
   for _, u := range us {
      fmt.Println(u)
      spec, err := ja3.Parse(hs.JA3(u.MD5))
      if err != nil {
         panic(err)
      }
      tr := ja3.NewTransport(spec)
      // FIXME
      req, err := http.NewRequest(
         "POST", "https://android.clients.google.com/auth",
         strings.NewReader(val.Encode()),
      )
      if err != nil {
         panic(err)
      }
      req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
      tcpConn, err := net.Dial("tcp", req.URL.Host + ":" + req.URL.Scheme)
      if err != nil {
         panic(err)
      }
      config := &tls.Config{ServerName: req.URL.Host}
      tlsConn := tls.UClient(tcpConn, config, tls.HelloCustom)
      if err := tlsConn.ApplyPreset(preset); err != nil {
         panic(err)
      }
      if err := req.Write(tlsConn); err != nil {
         panic(err)
      }
      res, err := http.ReadResponse(bufio.NewReader(tlsConn), req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      time.Sleep(time.Second)
   }
}






