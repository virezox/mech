package main

import (
   "embed"
   "fmt"
   "github.com/89z/torrent"
   "html/template"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
   "strings"
)

const origin = "https://dyncdn.me"

//go:embed index.html
var content embed.FS

func image(w http.ResponseWriter, r *http.Request) {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "torrent", r.URL.Path)
   if _, err := os.Stat(cache); err != nil {
      fmt.Println("Get", origin + r.URL.Path)
      res, err := http.Get(origin + r.URL.Path)
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      os.MkdirAll(filepath.Dir(cache), os.ModeDir)
      file, err := os.Create(cache)
      if err != nil {
         panic(err)
      }
      file.ReadFrom(res.Body)
      file.Close()
   } else {
      fmt.Println("Exist", cache)
   }
   file, err := os.Open(cache)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   io.Copy(w, file)
}

var done = make(map[string]bool)

func index(w http.ResponseWriter, r *http.Request) {
   val := r.URL.Query()
   search := val.Get("search")
   // favicon.ico
   if search == "" {
      return
   }
   in, err := torrent.NewResults(search, val.Get("page"))
   if err != nil {
      panic(err)
   }
   var (
      out []torrent.Result
      uniq = val.Get("uniq")
   )
   for _, r := range in {
      if uniq != "" {
         if done[r.Image] {
            continue
         } else {
            done[r.Image] = true
         }
      }
      p, err := url.Parse(r.Image)
      if err != nil {
         panic(err)
      }
      r.Image = p.Path
      if n := strings.Index(r.Genre, " IMDB:"); n != -1 {
         r.Genre = r.Genre[:n]
      }
      out = append(out, r)
   }
   t, err := template.ParseFS(content, "index.html")
   if err != nil {
      panic(err)
   }
   t.Execute(w, out)
}

func google(w http.ResponseWriter, r *http.Request) {
   re := regexp.MustCompile(`(.+\.\d{4})\.`)
   find := re.FindStringSubmatch(filepath.Base(r.URL.Path))
   if find == nil {
      fmt.Println(re)
      return
   }
   http.Redirect(
      w, r, "http://google.com/search?q=" + find[1], http.StatusSeeOther,
   )
}

func main() {
   http.HandleFunc("/", index)
   http.HandleFunc("/google/", google)
   http.HandleFunc("/mimages/", image)
   // slash prevents redirect
   fmt.Println(`localhost/?search=2020
localhost/?uniq=1&search=2020&page=2`)
   new(http.Server).ListenAndServe()
}
