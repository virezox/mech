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
   doc, err := mech.Parse(res.Body)
   if err != nil {
      return nil, err
   }
   var results []Result
   tr := doc.ByAttr("class", "lista2")
   for tr.Scan() {
      var r Result
      td := tr.ByTag("td")
      // Cat.
      td.Scan()
      // File
      td.Scan()
      // genre
      span := td.ByTag("span")
      span.Scan()
      r.Genre = span.Text()
      // title
      a := td.ByTag("a")
      a.Scan()
      r.Title = a.Text()
      // image
      r.Image = a.Attr("onmouseover")
      // torrent
      r.Torrent = Origin + a.Attr("href")
      // Added
      td.Scan()
      // Size
      td.Scan()
      r.Size = td.Text()
      // S.
      td.Scan()
      font := td.ByTag("font")
      font.Scan()
      r.Seeders = font.Text()
      // append
      r.Image = regexp.MustCompile(`\d+`).FindString(r.Image)
      r.Image = fmt.Sprintf(
         "https://dyncdn.me/mimages/%v/poster_opt.jpg", r.Image,
      )
      results = append(results, r)
   }
   return results, nil
}
