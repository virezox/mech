package main

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
)

func main() {
   var addr url.URL
   addr.Scheme = "https"
   addr.Host = "content.sr.roku.com"
   addr.Path = "/content/v1/roku-trc/105c41ea75775968b670fbb26978ed76"
   addr.RawQuery = url.Values{
      //"filter":[]string{"categoryObjects:genreAppropriate%20eq%20true"},
      "filter":[]string{"categoryObjects:genreAppropriate eq true"},
      "featureInclude":[]string{"bookmark,watchlist,linearSchedule"},
      "expand":[]string{strings.Join([]string{
         "categoryObjects",
         "credits",
         "next",
         "seasons",
         "seasons.episodes",
         "series",
         "series.seasons",
         "series.seasons.episodes",
         "viewOptions",
         "viewOptions.providerDetails",
      }, ",")},
      // if you remove this, the size goes way up
      "include": {"series.seasons.episodes.viewOptions"},
   }.Encode()
   req, err := http.NewRequest(
      "GET",
      "https://therokuchannel.roku.com/api/v2/homescreen/content/" + url.PathEscape(addr.String()),
      nil,
   )
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(buf))
   fmt.Println(len(buf))
}
