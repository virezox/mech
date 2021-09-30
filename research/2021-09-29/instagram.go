package instagram

import (
   "bytes"
   "fmt"
   "github.com/89z/parse/html"
   "github.com/89z/parse/json"
   "net/http"
)

const origin = "https://www.instagram.com"

type sidecar struct {
   Shortcode_Media struct {
      Edge_Sidecar_To_Children struct {
         Edges []struct {
            Node struct {
               Display_URL string
            }
         }
      }
   }
}

func newSidecar(id string) (*sidecar, error) {
   req, err := http.NewRequest("GET", origin + "/p/" + id + "/embed/", nil)
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
   for lex.NextTag("script") {
      extra := lex.Bytes()
      if bytes.Contains(extra, []byte(`"shortcode_media"`)) {
         car := new(sidecar)
         err := json.UnmarshalObject(extra, car)
         if err != nil {
            return nil, err
         }
         return car, nil
      }
   }
}
