package main

import (
   "crypto/md5"
   "encoding/hex"
   "flag"
   "io"
   "math/bits"
   "net/http"
   "net/url"
   "os"
)

const byteTable1 =
   "D6283B717076BE1BA4FE19575E6CBC21B214377D8CA2FA67556A95E3FA6778ED" +
   "8E553389A8CE36B35CD6B26F96C434B96AEC3495C4FA72FFB8428DFBEC70F085" +
   "46D8B2A1E0CEAE4B7DAEA487CEE3AC5155C436ADFCC4EA97706A85376AC868FA" +
   "FEB033B9677ECEE3CC86D69F767489E9DA9C78C595AAB034B3F27DB2A2EDE0B5" +
   "B68895D151D69E7DD1C8F9B770CC9CB692C5FADD9F28DAC7E0CA95B2DA3497CE" +
   "74FA37E97DC4A237FBFAF1CFAA897D55AE87BCF5E96AC468C7FA768514D0D0E5" +
   "CEFF19D6E5D6CCF1F46CE9E789B2B7AE2889BE5EDC876CF751F26778AEB34BA2" +
   "B3213B55F8B376B2CFB3B3FFB35E717DFAFCFFA87DFED89C1BC46AF988B5E5"

func genXGorgon(query string) ([]byte, error) {
   null_md5 := make([]byte, 16)
   obj := md5.New()
   io.WriteString(obj, query)
   buf := obj.Sum(nil)
   buf = append(buf, null_md5...)
   buf = append(buf, null_md5...)
   buf = append(buf, null_md5...)
   data2 := append(buf[:4], 0, 0, 0, 0)
   data2 = append(data2, buf[32:36]...)
   data2 = append(data2, 0, 0, 0, 0, 0, 0, 0, 0)
   byteTable2, err := hex.DecodeString(byteTable1)
   if err != nil {
      return nil, err
   }
   var myhex byte
   for i := range data2 {
      var hex1 byte
      if i == 0 {
         hex1 = byteTable2[byteTable2[0] - 1]
         byteTable2[i] = hex1
      } else if i == 1 {
         var temp byte = 0xD6 + 0x28
         hex1 = byteTable2[temp - 1]
         myhex = temp
         byteTable2[i] = hex1
      } else {
         temp := myhex + byteTable2[i]
         hex1 = byteTable2[temp - 1]
         myhex = temp
         byteTable2[i] = hex1
      }
      hex2 := byteTable2[hex1*2 - 1]
      data2[i] = hex2 ^ data2[i]
   }
   for i := range data2 {
      byte1 := data2[i]
      byte1 = bits.RotateLeft8(byte1, 4)
      if i == len(data2)-1 {
         byte1 ^= data2[0]
      } else {
         byte1 ^= data2[i+1]
      }
      byte2 := ((byte1 & 0x55) * 2) | ((byte1 & 0xAA) / 2)
      byte2 = ((byte2 & 0x33) * 4) | ((byte2 & 0xCC) / 4)
      byte3 := bits.RotateLeft8(byte2, 4)
      byte3 ^= 0xFF
      data2[i] = byte3 ^ 0x14
   }
   return append([]byte{0x3, 0x61, 0x41, 0x10, 0x80, 0x0}, data2...), nil
}

func main() {
   var (
      device_id string
      install_id string
      device_model string
   )
   flag.StringVar(&device_id, "d", "7044933115287176709", "device_id")
   flag.StringVar(&install_id, "i", "7044933274012026629", "install_id")
   flag.StringVar(&device_model, "m", "ONEPLUS A3010", "device_model")
   flag.Parse()
   req, err := http.NewRequest(
      "GET", "https://api-h2.tiktokv.com/aweme/v1/aweme/detail/", nil,
   )
   if err != nil {
      panic(err)
   }
   deviceParams := url.Values{
      "aid": {"1233"},
      "app_name": {"musical_ly"},
      "aweme_id": {"7038818332270808325"},
      "channel": {"googleplay"},
      "device_id": {device_id},
      "device_platform": {"android"},
      "device_type": {device_model},
      "iid": {install_id},
      "os_version": {"12"},
      "version_code": {"220405"},
   }
   req.URL.RawQuery = deviceParams.Encode()
   gorgon, err := genXGorgon(req.URL.RawQuery)
   if err != nil {
      panic(err)
   }
   req.Header = http.Header{
      "user-agent": {"okhttp/3.10.0.1"},
      "x-gorgon": {hex.EncodeToString(gorgon)},
      "x-khronos": {"0"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
