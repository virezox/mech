package vimeo

import (
   "fmt"
   "net/http"
)

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

func (r reference) token() string {
   return r.OpenAPI.Paths["_essentials"].Paths["/videos/{video_id}"].Get.Token
}

func main() {
   r, err := http.Get("http://developer.vimeo.com/api/reference/videos")
   if err != nil {
      panic(err)
   }
   defer r.Body.Close()
   c, err := newSegment("66531465").callable(r)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", c)
}
