package vimeo

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
)

const videos = "http://developer.vimeo.com/api/reference/videos"

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

func playground(videoID string) (*http.Request, error) {
   fmt.Println("GET", videos)
   res, err := http.Get(videos)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err := newSegment(videoID).callable(res)
   if err != nil {
      return nil, err
   }
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(body)
   req, err := http.NewRequest(
      "POST", "http://developer.vimeo.com/api/playground/callable", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   for _, c := range res.Cookies() {
      switch c.Name {
      case "XSRF-TOKEN":
         req.Header.Set("X-CSRF-TOKEN", c.Value)
      case "session":
         req.AddCookie(c)
      }
   }
   return req, nil
}
