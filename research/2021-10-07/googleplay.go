package main

import (
   "fmt"
   "github.com/89z/mech"
   "google.golang.org/protobuf/testing/protopack"
   "io"
   "net/http"
   "net/url"
)

const origin = "https://android.clients.google.com"

func details(device, app string) ([]byte, error) {
   req, err := http.NewRequest("GET", origin + "/fdfe/details", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {bearer},
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

func main() {
   mech.Verbose(true)
   data, err := details("38B5418D8683ADBB", "com.google.android.youtube")
   if err != nil {
      panic(err)
   }
   var mes protopack.Message
   mes.UnmarshalAbductive(data, nil)
   fmt.Printf("%+v\n", mes)
}
