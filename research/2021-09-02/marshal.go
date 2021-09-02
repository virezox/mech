package proto

import (
   "encoding/base64"
   "github.com/philpearl/plenc"
)

type continuation struct {
   A struct {
      VideoID string `plenc:"2"`
   } `plenc:"2"`
   B uint32 `plenc:"3"`
   C struct {
      A struct {
         VideoID string `plenc:"4"`
      } `plenc:"4"`
   } `plenc:"6"`
}

func newContinuation(videoID string) continuation {
   var con continuation
   con.A.VideoID = videoID
   con.B = 6
   con.C.A.VideoID = videoID
   return con
}

func (c continuation) encode() (string, error) {
   b, err := plenc.Marshal(nil, c)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(b), nil
}
