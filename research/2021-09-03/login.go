package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   q := make(url.Values)
   q.Set("device_country", "us")
   q.Set("operatorCountry", "us")
   q.Set("lang", "us")
   q.Set("sdk_version", "25")
   q.Set("accountType", "HOSTED_OR_GOOGLE")
   q.Set("Email", "srpen6@gmail.com")
   q.Set("service", "oauth2:https://www.googleapis.com/auth/googleplay")
   q.Set("source", "android")
   q.Set("androidId", "38B5418D8683ADBB")
   q.Set("app", "com.android.vending")
   q.Set("callerPkg", "com.android.vending")
   q.Set("Passwd", os.Args[1])
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/auth",
      strings.NewReader(q.Encode()),
   )
   if err != nil {
      panic(err)
   }
   req.Header = http.Header{
      "User-Agent": {"GoogleAuth/1.4 (GT-I9100 KTU84Q)"},
      "app": {"com.android.vending"},
      "content-type": {"application/x-www-form-urlencoded"},
      "device": {"38B5418D8683ADBB"},
   }
   dReq, err := httputil.DumpRequest(req, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(append(dReq, '\n'))
   res, err := new(http.Transport).RoundTrip(req)
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
