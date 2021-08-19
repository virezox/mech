package vimeo

import (
   "net/http"
)

type callable struct {
   App string `json:"app"`
   Group string `json:"group"`
   OperationID string `json:"operation_id"`
   PayloadParams string `json:"payload_params"`
   Ptoken string `json:"ptoken"`
   QueryParams string `json:"query_params"`
   Segments string `json:"segments"`
}

func newCallable(r *http.Response) (*callable, error) {
   call := new(callable)
   call.Group = "videos"
   call.OperationID = "get_video"
   call.PayloadParams = "{}"
   call.QueryParams = "{}"
   defer r.Body.Close()
   lex := html.NewLexer(r.Body)
   lex.NextAttr("class", "footers")
   var script []byte
   for lex.NextTag("script") {
      script = append(script, lex.Bytes()...)
   }
   val := js.NewLexer(script).Values()
   // app
   var apps []app
   if err := json.Unmarshal(val["apps"], &apps); err != nil {
      panic(err)
   }
   id := apps[0].ID
   // print ptoken
   var ref reference
   if err := json.Unmarshal(val["reference"], &ref); err != nil {
      panic(err)
   }
   ptoken := ref.OpenAPI.
   Paths["_essentials"].
   Paths["/videos/{video_id}"].Get.Token
}
