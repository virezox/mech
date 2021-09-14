package main

import (
   "fmt"
   "net/http"
   "net/url"
   "time"
)

// these all work
var ids = map[string]string{
   // 2 results
   "win10_app":"QObXsyIJdGdhklZlSmjL2OtWhZK1X1IL",
   // 2 results
   "win10_app_beta":"fokD8ahgluWvi0FjspkwaSGAxQ1hdH2j",
   // 4 results
   "pc_browser":"B31E7OJEB3BxbSbJBHarCQOhvKZUY09J",
   // 6 results
   "android":"dbdsA8b6V6Lw7wzu1x0T4CLxt58yd4Bf",
   // 9 results
   "widget":"Iy5e1Ri4GTNgrafaXe4mLpmJLXbXEfBR",
   // 12 results
   "ios":"Fiy8xlRI0xJNNGDLbPmGUjTpPRESPx8C",
   // 91 results
   "widget2":"LBCcHmRB8XSStWL6wKH2HPACspQlXg2P",
   // 334 results
   "mobi_browser":"iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX",
}

func main() {
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/resolve", nil,
   )
   if err != nil {
      panic(err)
   }
   val := url.Values{
      "url": {"https://soundcloud.com/pdis_inpartmaint/harold-budd-perhaps-moss"},
   }
   for k, v := range ids {
      val.Set("client_id", v)
      req.URL.RawQuery = val.Encode()
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      fmt.Println(res.Status, k)
      time.Sleep(100 * time.Millisecond)
   }
}
