package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "regexp"
)

type search struct {
   Contents struct {
      TwoColumnSearchResultsRenderer struct {
         PrimaryContents struct {
            SectionListRenderer struct {
               Contents []struct {
                  ItemSectionRenderer struct {
                     Contents []item
                  }
               }
            }
         }
      }
   }
}

func (s search) items() []item {
   var (
      iCons []item
      pCons = s.Contents.TwoColumnSearchResultsRenderer.PrimaryContents
   )
   for _, sCon := range pCons.SectionListRenderer.Contents {
      for _, iCon := range sCon.ItemSectionRenderer.Contents {
         if iCon.VideoRenderer.VideoID != "" {
            iCons = append(iCons, iCon)
         }
      }
   }
   return iCons
}

type item struct {
   VideoRenderer struct {
      VideoID string
   }
}

func newData(query string) (search, error) {
   req, err := http.NewRequest("GET", "https://www.youtube.com/results", nil)
   if err != nil {
      return search{}, err
   }
   val := req.URL.Query()
   val.Set("search_query", query)
   req.URL.RawQuery = val.Encode()
   fmt.Println("GET", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return search{}, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return search{}, err
   }
   re := regexp.MustCompile(" ytInitialData = ([^;]+)")
   find := re.FindSubmatch(body)
   if find == nil {
      return search{}, fmt.Errorf("FindSubmatch %v", re)
   }
   var s search
   if err := json.Unmarshal(find[1], &s); err != nil {
      return search{}, err
   }
   return s, nil
}
