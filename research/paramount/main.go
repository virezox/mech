package main

import (
   "bufio"
   "encoding/json"
   "github.com/89z/format/protobuf"
   "github.com/chris124567/hulu/widevine"
   "io"
   "net/http"
   "os"
)

func main() {
   key, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      panic(err)
   }
   id, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      panic(err)
   }
   file, err := os.Open("ignore/stream.mpd")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   initData, err := widevine.InitDataFromMPD(file)
   if err != nil {
      panic(err)
   }
   cdm, err := widevine.NewCDM(string(key), id, initData)
   if err != nil {
      panic(err)
   }
   _ = cdm
   // piece of shit server sends more data than the Content-Length, as least
   // with the response. So we need to look at the headers too.
   licenseRequest, err := readRequest("ignore/req.txt")
   if err != nil {
      panic(err)
   }
   licenseResponse, err := readResponse("ignore/res.txt")
   if err != nil {
      panic(err)
   }
   _ = licenseResponse
   //keys, err := cdm.GetLicenseKeys(licenseRequest, licenseResponse)
   licenseRequestParsed, err := protobuf.Unmarshal(licenseRequest)
   if err != nil {
      panic(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(licenseRequestParsed)
}

func readResponse(s string) ([]byte, error) {
   file, err := os.Open(s)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   res, err := http.ReadResponse(bufio.NewReader(file), nil)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

func readRequest(s string) ([]byte, error) {
   file, err := os.Open(s)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   req, err := http.ReadRequest(bufio.NewReader(file))
   if err != nil {
      return nil, err
   }
   defer req.Body.Close()
   return io.ReadAll(req.Body)
}
