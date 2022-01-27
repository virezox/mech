package ted

import (
   "net/http"
   "strings"
)

func get(slug string) (*http.Response, error) {
   var addr strings.Builder
   addr.WriteString("https://devices.ted.com/api/v2/videos/")
   addr.WriteString(slug)
   addr.WriteString("/react_native_v2.json")
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   return new(http.Transport).RoundTrip(req)
}
