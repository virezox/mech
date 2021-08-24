package main

import (
   "encoding/base64"
   p "google.golang.org/protobuf/testing/protopack"
)

type param struct {
   human string
   decode p.Message
}

var params = map[string][]param{
   "SORT BY": {
      {
         "Relevance", p.Message{
            p.Tag{1, p.VarintType}, p.Varint(0),
         },
      }, {
         "Rating", p.Message{
            p.Tag{1, p.VarintType}, p.Varint(1),
         },
      }, {
         "Upload date", p.Message{
            p.Tag{1, p.VarintType}, p.Varint(2),
         },
      }, {
         "View count", p.Message{
            p.Tag{1, p.VarintType}, p.Varint(3),
         },
      },
   },
   "UPLOAD DATE": {
      {
         "Last hour", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{1, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Today", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{1, p.VarintType}, p.Varint(2),
            },
         },
      }, {
         "This week", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{1, p.VarintType}, p.Varint(3),
            },
         },
      }, {
         "This month", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{1, p.VarintType}, p.Varint(4),
            },
         },
      }, {
         "This year", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{1, p.VarintType}, p.Varint(5),
            },
         },
      },
   },
   "TYPE": {
      {
         "Video", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Channel", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(2),
            },
         },
      }, {
         "Playlist", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(3),
            },
         },
      }, {
         "Movie", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{2, p.VarintType}, p.Varint(4),
            },
         },
      },
   },
   "DURATION": {
      {
         "Under 4 minutes", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Over 20 minutes", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(2),
            },
         },
      }, {
         "4 - 20 minutes", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{3, p.VarintType}, p.Varint(3),
            },
         },
      },
   },
   "FEATURES": {
      {
         "HD", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{4, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Subtitles/CC", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{5, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Creative Commons", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{6, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "3D", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{7, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Live", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{8, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "4K", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{14, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "360°", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{15, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "Location", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{23, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "HDR", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{25, p.VarintType}, p.Varint(1),
            },
         },
      }, {
         "VR180", p.Message{
            p.Tag{2, p.BytesType}, p.LengthPrefix{
               p.Tag{26, p.VarintType}, p.Varint(1),
            },
         },
      },
   },
}

func encode(m p.Message) string {
   return "youtube.com/results?search_query=autechre&sp=" +
   base64.StdEncoding.EncodeToString(m.Marshal())
}
