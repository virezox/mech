package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            mpxAccountId
            mpxGuid
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"`
      Name string `json:"name"` // String cannot represent a non string value
      Platform string `json:"platform"`
      Type string `json:"type"`
      User_ID string `json:"userId"` // can be empty
   } `json:"variables"`
}

func graphQL_compact(s string) string {
   old_new := []string{
      "\n", "",
      strings.Repeat(" ", 12), " ",
      strings.Repeat(" ", 9), " ",
      strings.Repeat(" ", 6), " ",
      strings.Repeat(" ", 3), " ",
   }
   return strings.NewReplacer(old_new...).Replace(s)
}

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   var p page_request
   p.Variables.App = "nbc"
   p.Variables.Name = "NBCE680517903"
   p.Variables.Platform = "android"
   p.Variables.Type = "VIDEO"
   p.Query = graphQL_compact(query)
   req_body := new(bytes.Buffer)
   enc := json.NewEncoder(req_body)
   enc.SetIndent("", " ")
   enc.Encode(p)
   req.Body = io.NopCloser(req_body)
   buf, err := httputil.DumpRequest(req, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      buf, err := httputil.DumpResponse(res, true)
      if err != nil {
         panic(err)
      }
      os.Stdout.Write(buf)
   }
   raw_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   var dst bytes.Buffer
   json.Indent(&dst, raw_body, "", " ")
   fmt.Println(dst.String())
}
