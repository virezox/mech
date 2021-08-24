package main

import (
   "encoding/base64"
   "fmt"
   p "google.golang.org/protobuf/testing/protopack"
)

type param struct {
   human, machine string
}

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
      {"Last hour", "EgIIAQ=="},
      {"Today", "EgQIAhAB"},
      {"This week", "EgQIAxAB"},
      {"This month", "EgQIBBAB"},
      {"This year", "EgQIBRAB"},
   },
   "SORT BY": {
      {"Relevance", "CAASAhAB"},
      {"Upload date", "CAISAhAB"},
      {"View count", "CAMSAhAB"},
      {"Rating", "CAESAhAB"},
   },
   "FEATURES": {
      {"360°", "EgJ4AQ=="},
      {"3D", "EgI4AQ=="},
      {"4K", "EgJwAQ=="},
      {"Creative Commons", "EgIwAQ=="},
      {"HD", "EgIgAQ=="},
      {"HDR", "EgPIAQE="},
      {"Live", "EgJAAQ=="},
      {"Location", "EgO4AQE="},
      {"Subtitles/CC", "EgIoAQ=="},
      {"VR180", "EgPQAQE="},
   },
   "TYPE": {
      {"Video", "EgIQAQ=="}, // 2 1
      {"Channel", "EgIQAg=="}, // 2 2
      {"Playlist", "EgIQAw=="}, // 2 3
      {"Movie", "EgIQBA=="}, // 2 4
   },
   "DURATION": {
      {"Under 4 minutes", "EgIYAQ=="}, // 3 1
      {"Over 20 minutes", "EgIYAg=="}, // 3 2
      {"4 - 20 minutes", "EgIYAw=="}, // 3 3
   },
}

func main() {
   m, err := decode(params["DURATION"][2].machine)
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
