package instagram

import (
   "bufio"
   "strings"
)

func bufioSplit(address string) string {
   buf := bufio.NewScanner(strings.NewReader(address))
   for buf.Scan() {
      if buf.Text() == "p" {
         buf.Scan()
         return buf.Text()
      }
   }
   return ""
}
