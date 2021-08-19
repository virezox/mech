package vimeo

import (
   "encoding/json"
   "github.com/89z/mech/html"
   "github.com/89z/mech/js"
   "net/http"
)

type segment struct {
   VideoID string `json:"video_id"`
}

func newSegment(videoID string) segment {
   return segment{videoID}
}

type callable struct {
   App int `json:"app"`
   Group string `json:"group"`
   OperationID string `json:"operation_id"`
   PayloadParams string `json:"payload_params"`
   Ptoken string `json:"ptoken"`
   QueryParams string `json:"query_params"`
   Segments string `json:"segments"`
}

func (s segment) callable(r *http.Response) (*callable, error) {
   lex := html.NewLexer(r.Body)
   lex.NextAttr("class", "footers")
   var script []byte
   for lex.NextTag("script") {
      script = append(script, lex.Bytes()...)
   }
   val := js.NewLexer(script).Values()
   call := new(callable)
   // app
   var apps []app
   if err := json.Unmarshal(val["apps"], &apps); err != nil {
      return nil, err
   }
   call.App = apps[0].ID
   // group
   call.Group = "videos"
   // operation_id
   call.OperationID = "get_video"
   // payload_params
   call.PayloadParams = "{}"
   // ptoken
   var ref reference
   if err := json.Unmarshal(val["reference"], &ref); err != nil {
      return nil, err
   }
   call.Ptoken = ref.token()
   // query_params
   call.QueryParams = "{}"
   // segments
   seg, err := json.Marshal(s)
   if err != nil {
      return nil, err
   }
   call.Segments = string(seg)
   // return
   return call, nil
}
