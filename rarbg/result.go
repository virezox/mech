package rarbg

import (
   "encoding/json"
   "fmt"
   "golang.org/x/net/html"
   "io"
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
   doc, err := newNode(res.Body)
   if err != nil {
      return nil, err
   }
   var results []Result
   for _, tr := range doc.byAttrAll("class", "lista2") {
      var r Result
      td := tr.byTagAll("td")
      // title
      r.Title = td[1].byTag("a").text()
      // image
      r.Image = td[1].byTag("a").attr("onmouseover")
      r.Image = regexp.MustCompile(`\d+`).FindString(r.Image)
      r.Image = fmt.Sprintf(
         "https://dyncdn.me/mimages/%v/poster_opt.jpg", r.Image,
      )
      // torrent
      r.Torrent = Origin + td[1].byTag("a").attr("href")
      // genre
      r.Genre = td[1].byTag("span").text()
      // size
      r.Size = td[3].text()
      // seeders
      r.Seeders = td[4].byTag("font").text()
      // append
      results = append(results, r)
   }
   return results, nil
}

func tag(t string) func(node) bool {
   return func(n node) bool {
      return n.Data == t
   }
}

type node struct {
   *html.Node
}

func newNode(r io.Reader) (node, error) {
   d, err := html.Parse(r)
   if err != nil {
      return node{}, err
   }
   return node{d}, nil
}

func (n node) attr(key string) string {
   for _, a := range n.Attr {
      if a.Key == key {
         return a.Val
      }
   }
   return ""
}

func (n node) byAttrAll(key, val string) []node {
   return n.selector(true, func(n node) bool {
      for _, a := range n.Attr {
         if a.Key == key && a.Val == val {
            return true
         }
      }
      return false
   })
}

func (n node) byTag(name string) node {
   for _, n := range n.selector(false, tag(name)) {
      return n
   }
   return node{}
}

func (n node) byTagAll(name string) []node {
   return n.selector(true, tag(name))
}

func (n node) selector(all bool, f func(n node) bool) []node {
   var (
      in = []node{n}
      out []node
   )
   for len(in) > 0 {
      n := in[0]
      if f(n) {
         out = append(out, n)
         if ! all {
            break
         }
      }
      for c := n.FirstChild; c != nil; c = c.NextSibling {
         in = append(in, node{c})
      }
      in = in[1:]
   }
   return out
}

func (n node) text() string {
   for _, n := range n.selector(false, func(n node) bool {
      return n.Type == html.TextNode
   }) {
      return n.Data
   }
   return ""
}
