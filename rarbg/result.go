package rarbg

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "os"
   "regexp"
)

type Result struct {
   Genre string
   Image string
   Seeders string
   Size string
   Title string
   Torrent string
}

// Perform a search with the given value. Note you will need to have previously
// saved the SKT cookie to the Cache folder using IamHuman.
func NewResults(search, page string) ([]Result, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   skt, err := os.Open(cache + "/mech/skt.json")
   if err != nil {
      return nil, err
   }
   defer skt.Close()
   cookie := new(http.Cookie)
   json.NewDecoder(skt).Decode(cookie)
   req, err := http.NewRequest("GET", Origin + "/torrents.php", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("by", "DESC")
   val.Set("category", "movies")
   val.Set("order", "seeders")
   val.Set("search", search)
   if page != "" {
      val.Set("page", page)
   }
   req.URL.RawQuery = val.Encode()
   req.AddCookie(cookie)
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   doc, err := mech.NewNode(res.Body)
   if err != nil {
      return nil, err
   }
   var results []Result
   for _, tr := range doc.ByAttrAll("class", "lista2") {
      var r Result
      td := tr.ByTagAll("td")
      // title
      r.Title = td[1].ByTag("a").Text()
      // image
      r.Image = td[1].ByTag("a").Attr("onmouseover")
      r.Image = regexp.MustCompile(`\d+`).FindString(r.Image)
      r.Image = fmt.Sprintf(
         "https://dyncdn.me/mimages/%v/poster_opt.jpg", r.Image,
      )
      // torrent
      r.Torrent = Origin + td[1].ByTag("a").Attr("href")
      // genre
      r.Genre = td[1].ByTag("span").Text()
      // size
      r.Size = td[3].Text()
      // seeders
      r.Seeders = td[4].ByTag("font").Text()
      // append
      results = append(results, r)
   }
   return results, nil
}
