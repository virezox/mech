package main

import (
   "encoding/base64"
   "fmt"
   p "google.golang.org/protobuf/testing/protopack"
)

var searchParams = map[string]map[string]string{
   "TYPE": {
      "Video" : "EgIQAQ==", // 1
      "Channel": "EgIQAg==",
      "Playlist": "EgIQAw==",
      "Movie": "EgIQBA==",
   },
   "UPLOAD DATE": {
      "Last hour": "EgIIAQ==",
      "Today": "EgQIAhAB",
      "This week": "EgQIAxAB",
      "This month": "EgQIBBAB",
      "This year": "EgQIBRAB",
   },
   "DURATION": {
      "Under 4 minutes": "EgIYAQ==",
      "4 - 20 minutes": "EgIYAw==",
      "Over 20 minutes": "EgIYAg==",
   },
   "SORT BY": {
      "Relevance": "CAASAhAB",
      "Upload date": "CAISAhAB",
      "View count": "CAMSAhAB",
      "Rating": "CAESAhAB",
   },
   "FEATURES": {
      "360°": "EgJ4AQ%253D%253D",
      "3D": "EgI4AQ%253D%253D",
      "4K": "EgJwAQ%253D%253D",
      "Creative Commons": "EgIwAQ%253D%253D",
      "HD": "EgIgAQ%253D%253D",
      "HDR": "EgPIAQE%253D",
      "Live": "EgJAAQ==",
      "Location": "EgO4AQE%253D",
      "Subtitles/CC": "EgIoAQ%253D%253D",
      "VR180": "EgPQAQE%253D",
   },
}

func main() {
   // byte 1
   params := "EgIQAQ==" // type video
   b, err := base64.StdEncoding.DecodeString(params)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", b)
   // byte 2
   m := p.Message{
      p.Tag{2, p.BytesType}, p.LengthPrefix{
         p.Tag{2, p.VarintType}, p.Varint(1),
      },
   }
   b = m.Marshal()
   fmt.Printf("%q\n", b)
}
