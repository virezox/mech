package main

import (
   "bytes"
   "fmt"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
   "research/paramount/widevine"
)

var LogLevel format.LogLevel

func main() {
   file, err := os.Open("ignore/stream.mpd")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   initData, err := widevine.InitDataFromMPD(file)
   if err != nil {
      panic(err)
   }
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      panic(err)
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      panic(err)
   }
   cdm, err := widevine.NewCDM(privateKey, clientID, initData)
   if err != nil {
      panic(err)
   }
   licenseRequest, err := cdm.GetLicenseRequest()
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST",
      "https://cbsi.live.ott.irdeto.com/widevine/getlicense?AccountId=cbsi&ContentId=eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
      bytes.NewReader(licenseRequest),
   )
   if err != nil {
      panic(err)
   }
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjY3NTI5NSwiZXhwIjoxNjUyNjgyNDk1LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI1ZDAwMzRjNy1mZGY1LTQ5MmUtOTQzNS02NzQ4NzU0ZjEyMDMifQ.8TJfoE-JTMSjL0Nq7nevN_QJR0GEaKmF5FXhJNM6ksc"}
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status)
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   keys, err := cdm.GetLicenseKeys(licenseRequest, licenseResponse)
   if err != nil {
      panic(err)
   }
   for _, key := range keys {
      fmt.Printf("%x\n", key.Value)
   }
}
