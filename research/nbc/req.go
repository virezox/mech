package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "bytes"
   "encoding/json"
   "os"
)

type friendship struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"`
      MpxGuid string `json:"mpxGuid"`
      OneApp bool `json:"oneApp"`
      Platform string `json:"platform"`
      TimeZone string `json:"timeZone"`
      Type string `json:"type"`
      UserId string `json:"userId"`
   } `json:"variables"`
}

func main() {
   query, err := os.ReadFile("query.js")
   if err != nil {
      panic(err)
   }
   var f friendship
   f.Query = string(query)
   f.Variables.App="nbc"
   f.Variables.MpxGuid="9000245869"
   f.Variables.OneApp=true
   f.Variables.Platform="android"
   f.Variables.TimeZone="America/Chicago"
   f.Variables.Type="segment"
   f.Variables.UserId="8292284999374523746"
   body := new(bytes.Buffer)
   if err := json.NewEncoder(body).Encode(f); err != nil {
      panic(err)
   }
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
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
