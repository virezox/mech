package main

import (
   "encoding/base64"
   "fmt"
   p "google.golang.org/protobuf/testing/protopack"
)


func decode(param string) (p.Message, error) {
   b, err := base64.StdEncoding.DecodeString(param)
   if err != nil {
      return nil, err
   }
   var m p.Message
   m.UnmarshalAbductive(b, nil)
   return m, nil
}

var params = map[string][]param{
   "UPLOAD DATE": {
      {"Last hour", "EgIIAQ==", nil},
      {"Today", "EgQIAhAB", nil},
      {"This week", "EgQIAxAB", nil},
      {"This month", "EgQIBBAB", nil},
      {"This year", "EgQIBRAB", nil},
   },
   "FEATURES": {
      {"360°", "EgJ4AQ==", nil},
      {"3D", "EgI4AQ==", nil},
      {"4K", "EgJwAQ==", nil},
      {"Creative Commons", "EgIwAQ==", nil},
      {"HD", "EgIgAQ==", nil},
      {"HDR", "EgPIAQE=", nil},
      {"Live", "EgJAAQ==", nil},
      {"Location", "EgO4AQE=", nil},
      {"Subtitles/CC", "EgIoAQ==", nil},
      {"VR180", "EgPQAQE=", nil},
   },
   "TYPE": {
      {"Video", "EgIQAQ==", nil}, // 2 1
      {"Channel", "EgIQAg==", nil}, // 2 2
      {"Playlist", "EgIQAw==", nil}, // 2 3
      {"Movie", "EgIQBA==", nil}, // 2 4
   },
   "DURATION": {
      {"Under 4 minutes", "EgIYAQ==", nil}, // 3 1
      {"Over 20 minutes", "EgIYAg==", nil}, // 3 2
      {"4 - 20 minutes", "EgIYAw==", nil}, // 3 3
   },
   "SORT BY": {
      {"Relevance", "CAASAhAB", nil},
      {"Upload date", "CAISAhAB", nil},
      {"View count", "CAMSAhAB", nil},
      {"Rating", "CAESAhAB", nil},
   },
}

type param struct {
   human string
   encode string
   decode p.Message
}

func main() {
   m, err := decode(params["SORT BY"][0].encode)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", m)
}

func encode(param int) string {
   m := p.Message{
      p.Tag{2, p.BytesType}, p.LengthPrefix{
         p.Tag{2, p.VarintType}, p.Varint(param),
      },
   }
   return "youtube.com/results?search_query=autechre&sp=" +
   base64.StdEncoding.EncodeToString(m.Marshal())
}
