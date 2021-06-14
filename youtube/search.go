package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

type Search struct {
   Contents struct {
      TwoColumnSearchResultsRenderer struct {
         PrimaryContents primaryContents
      }
   }
}

func NewSearch(query string) (Search, error) {
   req, err := http.NewRequest("GET", "https://www.youtube.com/results", nil)
   if err != nil {
      return Search{}, err
   }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Search{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return Search{}, err
   }
   re := regexp.MustCompile(">var ytInitialData = (.+);<")
   find := re.FindSubmatch(body)
   if find == nil {
      return Search{}, fmt.Errorf("FindSubmatch %v", re)
   }
   var s Search
   if err := json.Unmarshal(find[1], &s); err != nil {
      return Search{}, err
   }
   return s, nil
}

func (s Search) VideoRenderers() []VideoRenderer {
   var vids []VideoRenderer
   for _, sect := range s.primaryContents().SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.VideoRenderer)
      }
   }
   return vids
}

func (s Search) primaryContents() primaryContents {
   return s.Contents.TwoColumnSearchResultsRenderer.PrimaryContents
}

type VideoRenderer struct {
   VideoID string
}

type primaryContents struct {
   SectionListRenderer struct {
      Contents []struct {
         ItemSectionRenderer struct {
            Contents []struct {
               VideoRenderer VideoRenderer
            }
         }
      }
   }
}
