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
   req, err := http.NewRequest(
      "GET", "https://client.tlsfingerprint.io:8443", nil,
   )
   if err != nil {
      panic(err)
   }
   tcpConn, err := net.Dial("tcp", "client.tlsfingerprint.io:8443")
   if err != nil {
      panic(err)
   }
   config := &tls.Config{
      ServerName: req.URL.Hostname(),
   }
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
