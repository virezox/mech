package main

import (
   "encoding/base64"
   "fmt"
   "github.com/89z/mech/widevine"
   "net/http"
   "os"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   privateKey, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      panic(err)
   }
   clientID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      panic(err)
   }
   kID, err := base64.StdEncoding.DecodeString("62dqu8s0Xpa7z2FmMPGj2g==")
   if err != nil {
      panic(err)
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
   if err != nil {
      panic(err)
   }
   keys, err := mod.Post("https://shield-twoproxy.imggaming.com/proxy", http.Header{
      "Authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJjNzg0N2U3MC1hNTUzLTNkZTctYWM4OS0zNmNjN2YwMDAwMDEiLCJhdWQiOiJuYmEiLCJpc3MiOiJuYmFhcGkubmV1bGlvbi5jb20iLCJlaWQiOiJEN0U4QUQ0MC1GNkIzLTNERTUtQjQ4Mi1ERTQzN0YwMDAwMDEiLCJhaWQiOiJsaW9uLTM1MC1jaWFtLTM2MGE0MDUyY2FkMjY0YjAwOWE5OGY2MzJiOTZhY2ZlMTRiYTE1OWM1ODA1OGViZDA5ZmY5NjQyMmI5MTFjODAyZWZhZDQzMjBmMjNjY2FkNDVkOWE4M2M3M2M4NTk0MzAxNTE3ZTYxNjlhZDJmMWI0MWY3OGQ2ODYwYzAyMjJhNzAwNjY4ZWY1ZDVjY2ZjYSIsImRpZCI6IlVua25vdyIsInBsYyI6ZmFsc2UsImRlZiI6ImhkIiwiaWF0IjoxNjU0NDYwMDAxLCJleHAiOjE2NTQ0NjcyMDEsImlzcyI6Im5iYWFwaS5uZXVsaW9uLmNvbSJ9.TgNP6R1jOl9baOoYy59EET3hRh_z7GMvZaT5an9Khps"},
   })
   fmt.Printf("%+v\n", keys)
}
