package instagram

import (
   "github.com/89z/mech"
)

const (
   originI = "https://i.instagram.com"
   queryHash = "2efa04f61586458cef44441f474eee7c"
   // com.instagram.android
   userAgent = "Instagram 214.1.0.29.120 Android"
)

var LogLevel mech.LogLevel

// instagram.com/p/CT-cnxGhvvO
// instagram.com/p/yza2PAPSx2
func Valid(shortcode string) bool {
   switch len(shortcode) {
   case 10, 11:
      return true
   }
   return false
}

type Edge struct {
   Node struct {
      Display_URL string
      Video_URL string
   }
}

func (e Edge) URL() string {
   if e.Node.Video_URL != "" {
      return e.Node.Video_URL
   }
   return e.Node.Display_URL
}

type Item struct {
   Media struct {
      Video_Versions []struct {
         URL string
      }
   }
}

type Media struct {
   Shortcode_Media struct {
      Display_URL string
      Edge_Media_Preview_Like struct {
         Count int
      }
      Edge_Media_To_Parent_Comment struct {
         Edges []struct {
            Node struct {
               Text string
            }
         }
      }
      Edge_Sidecar_To_Children *struct {
         Edges []Edge
      }
      Video_URL string
   }
}

func (m Media) Edges() []Edge {
   if m.Shortcode_Media.Edge_Sidecar_To_Children == nil {
      return nil
   }
   return m.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}

type Query struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
   } `json:"variables"`
}

func NewQuery(shortcode string) Query {
   var val Query
   val.Query_Hash = queryHash
   val.Variables.Shortcode = shortcode
   return val
}
