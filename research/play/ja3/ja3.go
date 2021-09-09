package main

import (
   "bufio"
   "github.com/refraction-networking/utls"
   "net"
   "net/http"
   "os"
   "play"
)

func main() {
   req, err := http.NewRequest("GET", "https://ja3er.com/json", nil)
   if err != nil {
      panic(err)
   }
   tcpConn, err := net.Dial("tcp", req.URL.Host + ":" + req.URL.Scheme)
   if err != nil {
      panic(err)
   }
   config := &tls.Config{ServerName: req.URL.Host}
   tlsConn := tls.UClient(tcpConn, config, tls.HelloCustom)
   if err := tlsConn.ApplyPreset(play.Preset); err != nil {
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
   os.Stdout.ReadFrom(res.Body)
}
