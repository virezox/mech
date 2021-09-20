package github

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "strconv"
)

const Origin = "https://api.github.com"

var Verbose bool

type Repos struct {
   Items []struct {
      HTML_URL string
      Language string
      Stargazers_Count int
   }
}

type Search struct {
   url.Values
}

func NewSearch(q string) Search {
   return Search{
      url.Values{
         "q": {q},
      },
   }
}

// default 1
func (s Search) Page(value int) {
   val := strconv.Itoa(value)
   s.Set("page", val)
}

// default 30, max 100
func (s Search) PerPage(value int) {
   val := strconv.Itoa(value)
   s.Set("per_page", val)
}

// Set "x" to "nil" for no authentication.
func (s Search) Repos(x *Exchange) (*Repos, error) {
   req, err := http.NewRequest("GET", Origin + "/search/repositories", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = s.Encode()
   if x != nil {
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
   if Verbose {
      key := "X-Ratelimit-Remaining"
      fmt.Println(key, res.Header.Get(key))
   }
   rep := new(Repos)
   if err := json.NewDecoder(res.Body).Decode(rep); err != nil {
      return nil, err
   }
   return rep, nil
}
