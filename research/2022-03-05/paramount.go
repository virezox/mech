package paramount

import (
   "bufio"
   "strings"
)

func scanSlash(buf []byte, eof bool) (int, []byte, error) {
   for key, val := range buf {
      if val == '/' {
         return key+1, buf[:key], nil
      }
   }
   return 0, buf, bufio.ErrFinalToken
}

func doBufio(addr string) []string {
   src := bufio.NewScanner(strings.NewReader(addr))
   src.Split(scanSlash)
   var dst []string
   for src.Scan() {
      dst = append(dst, src.Text())
   }
   return dst
}
