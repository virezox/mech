package youtube

import (
   "encoding/json"
   "fmt"
   "testing"
)

func TestAndroid(t *testing.T) {
   client := Client{"ANDROID_EMBEDDED_PLAYER", "16.20"}
   for _, id := range []string{"HtVdAasjOgU", "SkRSXFQerZs"} {
      res, err := client.newPlayer(id).post()
      if err != nil {
         t.Fatal(err)
      }
      defer res.Body.Close()
      m := make(map[string]interface{})
      if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
         t.Fatal(err)
      }
      if _, ok := m["streamingData"]; ok {
         fmt.Printf("pass %+v\n", client)
      } else {
         fmt.Printf("fail %+v\n", client)
      }
   }
}
