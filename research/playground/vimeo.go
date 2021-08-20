package vimeo

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
   "net/http/httputil"
   "os"
)

type callable struct {
   Headers struct {
      Authorization string
   }
}

type playground struct {
   App int `json:"app"`
   Group string `json:"group"`
   OperationID string `json:"operation_id"`
   PayloadParams string `json:"payload_params"`
   Ptoken string `json:"ptoken"`
   QueryParams string `json:"query_params"`
   Segments string `json:"segments"`
}

type segment struct {
   VideoID string `json:"video_id"`
}

type videos struct {
   XSRF_Token string
   session string
   apps []struct {
      ID int
   }
   reference struct {
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
}

func newVideos() (*videos, error) {
   addr := "http://developer.vimeo.com/api/reference/videos"
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   vid := new(videos)
   for _, c := range res.Cookies() {
      switch c.Name {
      case "XSRF-TOKEN":
         vid.XSRF_Token = c.Value
      case "session":
         vid.session = c.Value
      }
   }
   lex := html.NewLexer(res.Body)
   lex.NextAttr("class", "footers")
   var script []byte
   for lex.NextTag("script") {
      script = append(script, lex.Bytes()...)
   }
   val := js.NewLexer(script).Values()
   if err := json.Unmarshal(val["apps"], &vid.apps); err != nil {
      return nil, err
   }
   if err := json.Unmarshal(val["reference"], &vid.reference); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v videos) callable(p *playground) (*callable, error) {
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(p)
   req, err := http.NewRequest(
      "POST", "http://developer.vimeo.com/api/playground/callable", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "Cookie": {"session=" + v.session},
      "X-CSRF-TOKEN": {v.XSRF_Token},
   }
   d, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   call := new(callable)
   if err := json.NewDecoder(res.Body).Decode(call); err != nil {
      return nil, err
   }
   return call, nil
}

func (v videos) playground(videoID string) (*playground, error) {
   tok, err := v.playgroundToken()
   if err != nil {
      return nil, err
   }
   seg, err := json.Marshal(segment{videoID})
   if err != nil {
      return nil, err
   }
   return &playground{
      App: v.apps[0].ID,
      Group: "videos",
      OperationID: "get_video",
      PayloadParams: "{}",
      Ptoken: tok,
      QueryParams: "{}",
      Segments: string(seg),
   }, nil
}

func (v videos) playgroundToken() (string, error) {
   ess, ok := v.reference.OpenAPI.Paths["_essentials"]
   if !ok {
      return "", fmt.Errorf("%v", v.reference.OpenAPI.Paths)
   }
   vid, ok := ess.Paths["/videos/{video_id}"]
   if !ok {
      return "", fmt.Errorf("%v", ess.Paths)
   }
   return vid.Get.Token, nil
}
