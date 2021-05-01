package main

import (
   "fmt"
   "github.com/89z/youtube"
   "strings"
)

var ids = []struct{in, out string}{
   {"9HzQvow8zF8", "f47cd0be8c3ccc5f"},
   {"BnEn7X3Pr7o", "067127ed7dcfafba"},
}

func main() {
   v, e := youtube.NewVideo(ids[0].in)
   if e != nil {
      panic(e)
   }
   front := "https://manifest.googlevideo.com/api/manifest/dash/"
   var back string
   fmt.Sscanf(v.StreamingData.DashManifestURL, front + "%v", &back)
   query := strings.Split(back, "/")
   out := new(strings.Builder)
   for n := 0; n < len(query); n += 2 {
      fmt.Fprint(out, query[n], "/", query[n+1], "/\n")
   }
   fmt.Println(front)
   fmt.Print(out.String())
}
