package github

import (
   "encoding/json"
   "fmt"
   "net/http"
)

const Origin = "https://api.github.com"

var Verbose bool

type RepoSearch struct {
   Items []struct {
      HTML_URL string
      Language string
      Stargazers_Count int
   }
}

func (x Exchange) RepoSearch(query, page string) (*RepoSearch, error) {
   req, err := http.NewRequest("GET", Origin + "/search/repositories", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("per_page", "100")
   q.Set("q", query)
   q.Set("page", page)
   req.URL.RawQuery = q.Encode()
   if x.Access_Token != "" {
      req.Header.Set("Authorization", "Bearer " + x.Access_Token)
   }
   if Verbose {
      fmt.Println(req.Method, req.URL)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   rs := new(RepoSearch)
   if err := json.NewDecoder(res.Body).Decode(rs); err != nil {
      return nil, err
   }
   return rs, nil
}
