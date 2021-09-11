package main

import (
   "crypto/md5"
   "fmt"
)

func two() []byte {
   var a []byte
   a = append(a, 0xA4)
   a = append(a, "MD5_ORIGIN"...)
   a = append(a, 0xA4)
   a = append(a, '9')
   a = append(a, 0xA4)
   a = append(a, "SNG_ID"...)
   a = append(a, 0xA4)
   a = append(a, "MEDIA_VERSION"...)
   b := md5.Sum(a[1:])
   a = append(b[:], a...)
   return append(a, 0)
}

func main() {
   fmt.Printf("%q\n", two())
}
