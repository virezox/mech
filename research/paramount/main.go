package main

import (
   "bufio"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/json"
   "encoding/pem"
   "errors"
   "github.com/89z/format/protobuf"
   "io"
   "net/http"
   "os"
)

func main() {
   key, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      panic(err)
   }
   cdm, err := NewCDM(key)
   if err != nil {
      panic(err)
   }
   licenseResponse, err := readResponse("ignore/res2.txt")
   if err != nil {
      panic(err)
   }
   license, err := protobuf.Unmarshal(licenseResponse)
   if err != nil {
      panic(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(license)
   licenseSessionKey, err := license.GetBytes(4)
   if err != nil {
      panic(err)
   }
   if _, err := rsa.DecryptOAEP(
      sha1.New(),
      nil,
      cdm.privateKey,
      licenseSessionKey,
      nil,
   ); err != nil {
      panic(err)
   }
}

type CDM struct {
   privateKey *rsa.PrivateKey
}

func NewCDM(privateKey []byte) (*CDM, error) {
   block, _ := pem.Decode(privateKey)
   if block == nil || (block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY") {
      return nil, errors.New("failed to decode device private key")
   }
   keyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      pcks8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
      if err != nil {
         return nil, err
      }
      keyParsed = pcks8Key.(*rsa.PrivateKey)
   }
   return &CDM{
      privateKey: keyParsed,
   }, nil
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
