package instagram

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

const (
   agent = "Instagram 214.1.0.29.120 Android"
   queryHash = "7d4d42b121a214d23bd43206e5142c8c"
)

var logLevel format.LogLevel

type media struct {
   Data struct {
      Shortcode_Media struct {
         Display_URL string
         Edge_Media_Preview_Like struct {
            Count int64
         }
         Edge_Media_To_Comment struct {
            Edges []struct {
               Node struct {
                  Text string
               }
            }
         }
         Edge_Sidecar_To_Children *struct {
            Edges []struct {
               Node struct {
                  Display_URL string
                  Video_URL string
               }
            }
         }
         Video_URL string
      }
   }
}

type mediaRequest struct {
   Query_Hash string `json:"query_hash"`
   Variables struct {
      Shortcode string `json:"shortcode"`
      Fetch_Comment_Count int `json:"fetch_comment_count"`
   } `json:"variables"`
}

func newMedia(shortcode string) (*media, error) {
   var body mediaRequest
   body.Query_Hash = queryHash
   body.Variables.Fetch_Comment_Count = 9
   body.Variables.Shortcode = shortcode
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://i.instagram.com/graphql/query/", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {auth},
      "Content-Type": {"application/json"},
      "User-Agent": {agent},
   }
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}
