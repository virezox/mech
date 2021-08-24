package main

import (
   "encoding/base64"
   p "google.golang.org/protobuf/testing/protopack"
)

var params = map[string][]param{
   "TYPE": {
      {
         "Video", "EgIQAQ==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Channel", "EgIQAg==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(2),
            },
         },
      }, {
         "Playlist", "EgIQAw==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(3),
            },
         },
      }, {
         "Movie", "EgIQBA==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(4),
            },
         },
      },
   },
   "DURATION": {
      {
         "Under 4 minutes", "EgIYAQ==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Over 20 minutes", "EgIYAg==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(2),
            },
         },
      }, {
         "4 - 20 minutes", "EgIYAw==", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(3),
            },
         },
      },
   },
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
   "SORT BY": {
      {"Relevance", "CAASAhAB", nil},
      {"Upload date", "CAISAhAB", nil},
      {"View count", "CAMSAhAB", nil},
      {"Rating", "CAESAhAB", nil},
   },
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

func encode(m p.Message) string {
   return "youtube.com/results?search_query=autechre&sp=" +
   base64.StdEncoding.EncodeToString(m.Marshal())
}

type param struct {
   human string
   encode string
   decode p.Message
}
