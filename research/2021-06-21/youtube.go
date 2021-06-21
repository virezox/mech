package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "os"
   "strings"
)

type m map[string]interface{}

func request() error {
   body := fmt.Sprintf(`
   {
      "query": %q, "context": {
         "client": {"clientName": "WEB", "clientVersion": %q}
      }
   }
   `, "nelly furtado say it right", youtube.VersionWeb)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8")
   req.URL.RawQuery = val.Encode()
   fmt.Println("POST", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   f, err := os.Create("file.json")
   if err != nil {
      return err
   }
   defer f.Close()
   f.ReadFrom(res.Body)
   return nil
}

func main() {
   err := request()
   if err != nil {
      panic(err)
   }
}
