package main

import (
   "bytes"
   "fmt"
   "github.com/89z/mech"
   "google.golang.org/protobuf/testing/protopack"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   txt, err := os.ReadFile("ac2dm.txt")
   if err != nil {
      panic(err)
   }
   val, err := url.ParseQuery(string(txt))
   if err != nil {
      panic(err)
   }
   ac2 := Ac2dm{val}
   mech.Verbose(true)
   auth, err := ac2.OAuth2()
   if err != nil {
      panic(err)
   }
   data, err := auth.Details("38B5418D8683ADBB", "com.google.android.youtube")
   if err != nil {
      panic(err)
   }
   var mes protopack.Message
   mes.UnmarshalAbductive(data, nil)
   fmt.Printf("%+v\n", mes)
}

type Ac2dm struct {
   url.Values
}

// Exchange refresh token (aas_et) for access token (Auth).
func (a Ac2dm) OAuth2() (*OAuth2, error) {
   val := url.Values{
      "Token": {
         a.Get("Token"),
      },
      "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/auth", strings.NewReader(val.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   query, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &OAuth2{
      ParseQuery(query),
   }, nil
}

type OAuth2 struct {
   url.Values
}

// device is Google Service Framework.
func (o OAuth2) Details(device, app string) ([]byte, error) {
   req, err := http.NewRequest("GET", Origin + "/fdfe/details", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {
         "Bearer " + o.Get("Auth"),
      },
      "X-DFE-Device-Id": {device},
   }
   val := url.Values{
      "doc": {app},
   }
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

const Origin = "https://android.clients.google.com"

func ParseQuery(query []byte) url.Values {
   res := make(url.Values)
   for _, pair := range bytes.Split(query, []byte{'\n'}) {
      nv := bytes.SplitN(pair, []byte{'='}, 2)
      res.Add(string(nv[0]), string(nv[1]))
   }
   return res
}
