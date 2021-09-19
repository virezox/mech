package github

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
)

const clientID = "3f8b8834a91f0caad392"

type OAuth struct {
   url.Values
}

func NewOAuth() (*OAuth, error) {
   data := url.Values{
      "client_id": {clientID},
   }
   res, err := http.PostForm("https://github.com/login/device/code", data)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   val, err := url.ParseQuery(string(body))
   if err != nil {
      return nil, err
   }
   return &OAuth{val}, nil
}
