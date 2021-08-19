package vimeo

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
)

func request1() (*http.Response, error) {
   return http.Get("http://developer.vimeo.com/api/reference/videos")
}
