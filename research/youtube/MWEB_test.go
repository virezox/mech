package youtube

import (
   "fmt"
   "github.com/89z/format/json"
   "net/http"
   "testing"
)

func TestMweb(t *testing.T) {
   client, err := newInnertubeClient()
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(client.ClientName, client.ClientVersion)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, client.ClientName, client.ClientVersion)
}

type innertubeClient struct {
   ClientName string
   ClientVersion string
}

func newInnertubeClient() (*innertubeClient, error) {
   req, err := http.NewRequest("GET", "https://m.youtube.com", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "iPad")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var (
      client = new(innertubeClient)
      sep = []byte(`"client":`)
   )
   if err := json.Decode(res.Body, sep, client); err != nil {
      return nil, err
   }
   return client, nil
}
