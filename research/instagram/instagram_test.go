package instagram

import (
   "os"
   "testing"
)

func TestGraphQL(t *testing.T) {
   res, err := GraphQL()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

func TestMedia(t *testing.T) {
   res, err := Media()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
