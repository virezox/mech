package tiktok

import (
   "crypto/md5"
   "encoding/binary"
   "encoding/hex"
   "io"
   "strconv"
)

func xGorgon(timeMillis uint32, inputBytes []byte) {
   data1 = []
   data1.append("3")
   data1.append("61")
   data1.append("41")
   data1.append("10")
   data1.append("80")
   data1.append("0")
   data2 = input(timeMillis, inputBytes)
   data2 = initialize(data2)
   data2 = handle(data2)
   for i in range(len(data2)):
       data1.append(data2[i])
   xGorgonStr = ""
   for i in range(len(data1)):
       temp = data1[i] + ""
       if len(temp) > 1:
           xGorgonStr += temp
       else:
           xGorgonStr += "0"
           xGorgonStr += temp
   return xGorgonStr
}

const byteTable1 =
   "D6283B717076BE1BA4FE19575E6CBC21B214377D8CA2FA67556A95E3FA6778ED" +
   "8E553389A8CE36B35CD6B26F96C434B96AEC3495C4FA72FFB8428DFBEC70F085" +
   "46D8B2A1E0CEAE4B7DAEA487CEE3AC5155C436ADFCC4EA97706A85376AC868FA" +
   "FEB033B9677ECEE3CC86D69F767489E9DA9C78C595AAB034B3F27DB2A2EDE0B5" +
   "B68895D151D69E7DD1C8F9B770CC9CB692C5FADD9F28DAC7E0CA95B2DA3497CE" +
   "74FA37E97DC4A237FBFAF1CFAA897D55AE87BCF5E96AC468C7FA768514D0D0E5" +
   "CEFF19D6E5D6CCF1F46CE9E789B2B7AE2889BE5EDC876CF751F26778AEB34BA2" +
   "B3213B55F8B376B2CFB3B3FFB35E717DFAFCFFA87DFED89C1BC46AF988B5E5"

func getXGon(url string) string {
   null_md5_string := "00000000000000000000000000000000"
   obj := md5.New()
   io.WriteString(obj, url)
   sb := hex.EncodeToString(obj.Sum(nil))
   sb += null_md5_string
   sb += null_md5_string
   sb += null_md5_string
   return sb
}

func initialize(data []byte) ([]byte, error) {
   byteTable2, err := hex.DecodeString(byteTable1)
   if err != nil {
      return nil, err
   }
   var myhex byte
   for i := range data {
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
      data[i] = hex2 ^ data[i]
   }
   return data, nil
}

func input(timeMillis uint32, inputBytes []byte) []byte {
   var result []byte
   for i := range [4]struct{}{} {
      temp := inputBytes[i]
      result = append(result, temp)
   }
   result = append(result, 0, 0, 0, 0)
   for i := range [4]struct{}{} {
      temp := inputBytes[i+32]
      result = append(result, temp)
   }
   result = append(result, 0, 0, 0, 0)
   var tempByte [4]byte
   binary.BigEndian.PutUint32(tempByte[:], timeMillis)
   return append(result, tempByte[:]...)
}

func str2hex(s string) (int64, error) {
   return strconv.ParseInt(s, 16, 64)
}

func strToByte(str string) ([]byte, error) {
   return hex.DecodeString(str)
}
