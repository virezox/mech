package paramount

import (
   "bytes"
   "errors"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
)

var LogLevel format.LogLevel

func KeyContainers(contentID, bearer string) ([]KeyContainer, error) {
   file, err := os.Open("ignore/stream.mpd")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   initData, err := initDataFromMPD(file)
   if err != nil {
      return nil, err
   }
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      return nil, err
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      return nil, err
   }
   mod, err := NewModule(privateKey, clientID, initData)
   if err != nil {
      return nil, err
   }
   licenseRequest, err := mod.LicenseRequest()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST",
      "https://cbsi.live.ott.irdeto.com/widevine/getlicense?AccountId=cbsi&ContentId=" + contentID,
      bytes.NewReader(licenseRequest),
   )
   if err != nil {
      return nil, err
   }
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return mod.KeyContainers(licenseRequest, licenseResponse)
}
