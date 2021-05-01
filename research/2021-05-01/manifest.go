package main

import (
   "fmt"
   "github.com/89z/youtube"
   "os"
   "strings"
)

func main() {
   if len(os.Args) != 2 {
      println("manifest id")
      os.Exit(1)
   }
   v, e := youtube.NewVideo(os.Args[1])
   if e != nil {
      panic(e)
   }
   if v.StreamingData.DashManifestURL == "" {
      println("missing dashManifestUrl")
      os.Exit(1)
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
