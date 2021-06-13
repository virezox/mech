package rarbg

import (
   "encoding/json"
   "fmt"
   "golang.org/x/net/html"
   "net/http"
   "os"
   "regexp"
)

func queryAll(n *html.Node, f func(n *html.Node) bool) []*html.Node {
   var (
      in = []*html.Node{n}
      out []*html.Node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
      }
      for n = n.FirstChild; n != nil; n = n.NextSibling {
         in = append(in, n)
      }
      in = in[1:]
   }
   return out
}

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
   doc, err := html.Parse(res.Body)
   if err != nil {
      return nil, err
   }
   rows := queryAll(doc, func(n *html.Node) bool {
      if n.Data != "tr" {
         return false
      }
      if len(n.Attr) == 0 {
         return false
      }
      return n.Attr[0].Val == "lista2"
   })
   var results []Result
   for _, row := range rows {
      cols := queryAll(row, func(n *html.Node) bool {
         return n.Data == "td"
      })
      var r Result
      //        <td>    <a>        #text
      r.Title = cols[1].FirstChild.FirstChild.Data
      //        <td>    <a>        onmouseover
      r.Image = cols[1].FirstChild.Attr[0].Val
      r.Image = regexp.MustCompile(`\d+`).FindString(r.Image)
      r.Image = fmt.Sprintf(
         "https://dyncdn.me/mimages/%v/poster_opt.jpg", r.Image,
      )
      //                   <td>    <a>        href
      r.Torrent = Origin + cols[1].FirstChild.Attr[2].Val
      //              <td>    #text     <span>      #text
      r.Genre = cols[1].LastChild.PrevSibling.FirstChild.Data
      //       <td>    #text
      r.Size = cols[3].FirstChild.Data
      //          <td>    <font>     #text
      r.Seeders = cols[4].FirstChild.FirstChild.Data
      // append
      results = append(results, r)
   }
   return results, nil
}
