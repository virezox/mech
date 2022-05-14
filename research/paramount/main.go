package main

import (
   "bytes"
   "crypto/md5"
   "encoding/hex"
   "fmt"
   "github.com/chris124567/hulu/widevine"
   "github.com/google/uuid"
   "io"
   "lukechampine.com/flagg"
   "net/http"
   "os"
   "strings"
   "time"
   hulu "github.com/chris124567/hulu/client"
)

func main() {
   // parse init data/PSSH from XML
   initData, err := widevine.InitDataFromMPD(response.Body)
   if err != nil {
      panic(err)
   }
   cdm, err := widevine.NewDefaultCDM(initData)
   if err != nil {
      panic(err)
   }
   licenseRequest, err := cdm.GetLicenseRequest()
   if err != nil {
      panic(err)
   }
   request, err := http.NewRequest(http.MethodPost, playlist.WvServer, bytes.NewReader(licenseRequest))
   if err != nil {
      panic(err)
   }
   // hulu actually checks for headers here so this is necessary
   request.Header = hulu.StandardHeaders()
   request.Close = true
   // send license request to license server
   response, err = client.Do(request)
   if err != nil {
      panic(err)
   }
   defer response.Body.Close()
   licenseResponse, err := io.ReadAll(response.Body)
   if err != nil {
      panic(err)
   }
   // parse keys from response
   keys, err := cdm.GetLicenseKeys(licenseRequest, licenseResponse)
   if err != nil {
      panic(err)
   }
   for _, key := range keys {
      if key.Type == widevine.License_KeyContainer_CONTENT {
         command += " --key " + hex.EncodeToString(key.ID) + ":" + hex.EncodeToString(key.Value)
      }
   }
}
