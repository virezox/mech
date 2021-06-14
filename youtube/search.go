package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

type ytInitialData struct {
   Contents struct {
      TwoColumnSearchResultsRenderer struct {
         PrimaryContents primaryContents
      }
   }
}

func newYTInitialData(query string) (ytInitialData, error) {
   req, err := http.NewRequest("GET", "https://www.youtube.com/results", nil)
   if err != nil {
      return ytInitialData{}, err
   }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   fmt.Println("GET", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return ytInitialData{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return ytInitialData{}, err
   }
   re := regexp.MustCompile(" ytInitialData = ([^;]+)")
   find := re.FindSubmatch(body)
   if find == nil {
      return ytInitialData{}, fmt.Errorf("FindSubmatch %v", re)
   }
   var data ytInitialData
   if err := json.Unmarshal(find[1], &data); err != nil {
      return ytInitialData{}, err
   }
   return data, nil
}

type primaryContents struct {
   SectionListRenderer struct {
      Contents []struct {
         ItemSectionRenderer struct {
            Contents []struct {
               VideoRenderer videoRenderer
            }
         }
      }
   }
}

func (d ytInitialData) primaryContents() primaryContents {
   return d.Contents.TwoColumnSearchResultsRenderer.PrimaryContents
}

type videoRenderer struct {
   VideoID string
}

func (p primaryContents) videoRenderers() []videoRenderer {
   var video []videoRenderer
   for _, section := range p.SectionListRenderer.Contents {
      for _, item := range section.ItemSectionRenderer.Contents {
         video = append(video, item.VideoRenderer)
      }
   }
   return video
}
