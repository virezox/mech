package main

import (
   "encoding/json"
   "fmt"
   "net/url"
   "os"
)

func iterate() ([]string, error) {
   f, e := os.Open(`D:\Git\umber\umber.json`)
   if e != nil { return nil, e }
   defer f.Close()
   var a []struct { Q string }
   json.NewDecoder(f).Decode(&a)
   out := make([]string, len(a))
   for n, vid := range a {
      val, e := url.ParseQuery(vid.Q)
      if e != nil { return nil, e }
      if val.Get("p") != "y" {
         continue
      }
      out[n] = val.Get("b")
   }
   return out, nil
}

type image struct {
   vi string
   left string
   right string
   ext string
}

var images = []image{
   {right:"0", ext:"jpg"},
   {right:"1", ext:"jpg"},
   {right:"2", ext:"jpg"},
   {right:"3", ext:"jpg"},
   {right:"default", ext:"jpg"},
   {left:"hq", right:"1", ext:"jpg"},
   {left:"hq", right:"2", ext:"jpg"},
   {left:"hq", right:"3", ext:"jpg"},
   {left:"hq", right:"720", ext:"jpg"},
   {left:"hq", right:"default", ext:"jpg"},
   {left:"maxres", right:"1", ext:"jpg"},
   {left:"maxres", right:"2", ext:"jpg"},
   {left:"maxres", right:"3", ext:"jpg"},
   {left:"maxres", right:"default", ext:"jpg"},
   {left:"mq", right:"1", ext:"jpg"},
   {left:"mq", right:"2", ext:"jpg"},
   {left:"mq", right:"3", ext:"jpg"},
   {left:"mq", right:"default", ext:"jpg"},
   {left:"sd", right:"1", ext:"jpg"},
   {left:"sd", right:"2", ext:"jpg"},
   {left:"sd", right:"3", ext:"jpg"},
   {left:"sd", right:"default", ext:"jpg"},
   {vi:"_webp", right:"0", ext:"webp"},
   {vi:"_webp", right:"1", ext:"webp"},
   {vi:"_webp", right:"2", ext:"webp"},
   {vi:"_webp", right:"3", ext:"webp"},
   {vi:"_webp", right:"default", ext:"webp"},
   {vi:"_webp", left:"hq", right:"1", ext:"webp"},
   {vi:"_webp", left:"hq", right:"2", ext:"webp"},
   {vi:"_webp", left:"hq", right:"3", ext:"webp"},
   {vi:"_webp", left:"hq", right:"720", ext:"webp"},
   {vi:"_webp", left:"hq", right:"default", ext:"webp"},
   {vi:"_webp", left:"maxres", right:"1", ext:"webp"},
   {vi:"_webp", left:"maxres", right:"2", ext:"webp"},
   {vi:"_webp", left:"maxres", right:"3", ext:"webp"},
   {vi:"_webp", left:"maxres", right:"default", ext:"webp"},
   {vi:"_webp", left:"mq", right:"1", ext:"webp"},
   {vi:"_webp", left:"mq", right:"2", ext:"webp"},
   {vi:"_webp", left:"mq", right:"3", ext:"webp"},
   {vi:"_webp", left:"mq", right:"default", ext:"webp"},
   {vi:"_webp", left:"sd", right:"1", ext:"webp"},
   {vi:"_webp", left:"sd", right:"2", ext:"webp"},
   {vi:"_webp", left:"sd", right:"3", ext:"webp"},
   {vi:"_webp", left:"sd", right:"default", ext:"webp"},
}

func main() {
   for _, i := range images {
      fmt.Printf(
         "http://i.ytimg.com/vi%v/UpNXI3_ctAc/%v%v.%v\n",
         i.vi, i.left, i.right, i.ext,
      )
   }
   ids, err := iterate()
   if err != nil {
      panic(err)
   }
   for _, id := range ids {
      fmt.Println(id)
   }
}
