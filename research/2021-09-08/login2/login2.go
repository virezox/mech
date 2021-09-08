package main

import (
   "bufio"
   "fmt"
   "github.com/refraction-networking/utls"
   "net"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

var hello = []byte{
   1,0,1,39,3,3,15,133,173,125,240,166,91,143,109,219,170,61,218,252,140,101,245,
   228,231,97,45,116,163,191,20,195,182,189,1,69,174,180,32,179,214,41,214,85,21,
   184,89,149,62,130,157,225,147,209,116,244,102,203,96,54,114,147,218,79,211,226,
   10,146,130,3,161,0,86,19,2,19,3,19,1,192,44,192,48,192,43,192,47,204,169,204,
   168,0,159,0,158,204,170,192,175,192,173,192,174,192,172,192,36,192,40,192,35,
   192,39,192,10,192,20,192,9,192,19,192,163,192,159,192,162,192,158,0,107,0,103,0,
   57,0,51,0,157,0,156,192,161,192,157,192,160,192,156,0,61,0,60,0,53,0,47,0,255,1,
   0,0,136,0,0,0,31,0,29,0,0,26,97,110,100,114,111,105,100,46,99,108,105,101,110,
   116,115,46,103,111,111,103,108,101,46,99,111,109,0,11,0,4,3,0,1,2,0,10,0,12,0,
   10,0,29,0,23,0,30,0,25,0,24,0,35,0,0,0,22,0,0,0,23,0,0,0,13,0,6,0,4,4,3,4,1,0,
   43,0,3,2,3,3,0,45,0,2,1,1,0,51,0,38,0,36,0,29,0,32,98,220,53,111,22,63,122,51,
   236,227,56,80,137,160,168,176,132,132,98,194,37,227,158,176,120,246,109,186,131,
   55,33,72,
}

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
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/auth",
      strings.NewReader(val.Encode()),
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   dReq, err := httputil.DumpRequest(req, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(append(dReq, '\n'))
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
   dRes, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(dRes)
}
