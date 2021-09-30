package instagram

import (
   "bytes"
   "fmt"
   "github.com/89z/parse/html"
   "github.com/89z/parse/json"
   "net/http"
)

const Origin = "https://www.instagram.com"

var Verbose bool

// instagram.com/p/CT-cnxGhvvO
func ValidID(id string) error {
   if len(id) == 11 {
      return nil
   }
   return fmt.Errorf("%q invalid as ID", id)
}

type Edge struct {
   Node struct {
      Display_URL string
   }
}

type Sidecar struct {
   Shortcode_Media struct {
      Edge_Sidecar_To_Children struct {
         Edges []Edge
      }
   }
}

func NewSidecar(id string) (*Sidecar, error) {
   req, err := http.NewRequest("GET", Origin + "/p/" + id + "/embed/", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "Mozilla")
   fmt.Println(req.Method, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   lex := html.NewLexer(res.Body)
   media := []byte(`"shortcode_media"`)
   for lex.NextTag("script") {
      extra := lex.Bytes()
      if bytes.Contains(extra, media) {
         car := new(Sidecar)
         err := json.UnmarshalObject(extra, car)
         if err != nil {
            return nil, err
         }
         return car, nil
      }
   }
   return nil, fmt.Errorf("%s not found", media)
}

func (s Sidecar) Edges() []Edge {
   return s.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}
