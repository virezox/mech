package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
   "os"
)

func request1() (*http.Response, error) {
   return http.Get("http://developer.vimeo.com/api/reference/videos")
}

var _ = fmt.Print

type app struct {
   ID int
}

type reference struct {
   OpenAPI struct {
      Paths map[string]struct {
         Paths map[string]struct {
            Get struct {
               Token string `json:"x-playground-token"`
            }
         }
      }
   }
}

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   lex := html.NewLexer(f)
   lex.NextAttr("class", "footers")
   var script []byte
   for lex.NextTag("script") {
      script = append(script, lex.Bytes()...)
   }
   val := js.NewLexer(script).Values()
   // print ptoken
   var ref reference
   if err := json.Unmarshal(val["reference"], &ref); err != nil {
      panic(err)
   }
   ptoken := ref.OpenAPI.
      Paths["_essentials"].
      Paths["/videos/{video_id}"].Get.Token
   fmt.Println(ptoken)
   // print app
   var apps []app
   if err := json.Unmarshal(val["apps"], &apps); err != nil {
      panic(err)
   }
   id := apps[0].ID
   fmt.Println(id)
}
