package tiktok

import (
   "crypto/md5"
   "encoding/hex"
   "io"
   "strconv"
)

def initialize(data):
   myhex = 0
   byteTable2 = byteTable1.split(" ")
   for i in range(len(data)):
       hex1 = 0
       if i == 0:
           hex1 = int(byteTable2[int(byteTable2[0], 16) - 1], 16)
           byteTable2[i] = hex(hex1)
       elif i == 1:
           temp = int("D6", 16) + int("28", 16)
           if temp > 256:
               temp -= 256
           hex1 = int(byteTable2[temp - 1], 16)
           myhex = temp
           byteTable2[i] = hex(hex1)
       else:
           temp = myhex + int(byteTable2[i], 16)
           if temp > 256:
               temp -= 256
           hex1 = int(byteTable2[temp - 1], 16)
           myhex = temp
           byteTable2[i] = hex(hex1)
       if hex1 * 2 > 256:
           hex1 = hex1 * 2 - 256
       else:
           hex1 = hex1 * 2
       hex2 = byteTable2[hex1 - 1]
       result = int(hex2, 16) ^ int(data[i], 16)
       data[i] = hex(result)
   for i in range(len(data)):
       data[i] = data[i].replace("0x", "")
   return data

func input(timeMillis int64, inputBytes []byte) []string {
   var result []string
   for i := range [4]struct{}{} {
      temp := strconv.FormatInt(
         int64(inputBytes[i]), 16,
      )
      result = append(result, temp)
   }
   result = append(result, "0", "0", "0", "0")
   for i := range [4]struct{}{} {
      temp := strconv.FormatInt(
         int64(inputBytes[i+32]), 16,
      )
      result = append(result, temp)
   }
   result = append(result, "0", "0", "0", "0")
   tempByte := strconv.FormatInt(timeMillis, 16)
   for i := range [4]struct{}{} {
      result = append(result, tempByte[i*2 : 2*i+2])
   }
   return result
}

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

func str2hex(s string) (int64, error) {
   return strconv.ParseInt(s, 16, 64)
}

func strToByte(str string) ([]byte, error) {
   return hex.DecodeString(str)
}
