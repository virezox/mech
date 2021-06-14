package main

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const gatewayWWW = "https://www.deezer.com/ajax/gw-light.php"

func auth() error {
   in, out := map[string]string{
      "mail": "srpen6@gmail.com", "password": "encryptedPassword",
   }, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := http.NewRequest("POST", gatewayWWW, out)
   if err != nil {
      return err
   }
   println(req)
   return nil
}

func main() {
   auth()
}
