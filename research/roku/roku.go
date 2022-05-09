package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "getwvkeys.cc"
   req.URL.Path = "/api"
   req.URL.Scheme = "https"
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

var body = strings.NewReader(`{
   "buildInfo": "",
   "license": "https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKZ0AJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIC63N2hKChJVVAEieJRdasQTDLxWpWxGIczvZHdpGu2FYj9dY5Yu2Nm148TQigP5OYYg3VYpZ6k7nbLR7UhYB28CRzzr57UkY7uV42YIpil7OqAWD_8dfW5l99IdN8QuYr0bHjwIizxMJOASXIqdBZuRXv85GlPGeIMX2wHu9YKHyZ7_wD4bU1WO3aymDA1vlZx5oqQ",
   "pssh": "AAAAW3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADsIARIQ62dqu8s0Xpa7z2FmMPGj2hoNd2lkZXZpbmVfdGVzdCIQZmtqM2xqYVNkZmFsa3IzaioCSEQyAA=="
}
`)
