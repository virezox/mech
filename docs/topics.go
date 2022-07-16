package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

const topics = "https://api.github.com/repos/89z/mech/topics"

var names = []string{
   "youtube",
   "soundcloud",
   "roku",
   "vimeo",
   "paramount",
   "nbc",
   "abc",
   ///////////
   "bandcamp",
   "cbc-gem",
   "amc",
}

func userinfo() (*url.Userinfo, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   buf, err := os.ReadFile(home + "/.git-credentials")
   if err != nil {
      return nil, err
   }
   var ref url.URL
   if err := ref.UnmarshalBinary(bytes.TrimSpace(buf)); err != nil {
      return nil, err
   }
   return ref.User, nil
}

func main() {
   buf, err := json.Marshal(map[string][]string{
      "names": names,
   })
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest("PUT", topics, bytes.NewReader(buf))
   if err != nil {
      panic(err)
   }
   info, err := userinfo()
   if err != nil {
      panic(err)
   }
   password, ok := info.Password()
   if ok {
      req.SetBasicAuth(info.Username(), password)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
}
