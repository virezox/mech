package tomato

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

type Movie struct {
   RTID string
   Score string
   Title string
   URLID string
}

func NewMovie(addr string) (Movie, error) {
   fmt.Println(invert, "Get", reset, addr)
   res, err := http.Get(addr)
   if err != nil {
      return Movie{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Movie{}, err
   }
   re := regexp.MustCompile(`\.mpscall = (.+);`)
   find := re.FindSubmatch(body)
   if find == nil {
      return Movie{}, fmt.Errorf("FindSubmatch %v", re)
   }
   var mov Movie
   if err := json.Unmarshal(find[1], &mov); err != nil {
      return Movie{}, err
   }
   return mov, nil
}
