package main

import (
   "bufio"
   "strings"
)

func comma(data []byte, eof bool) (int, []byte, error) {
   var tok []byte
   for adv, b := range data {
      if b != ',' {
         tok = append(tok, b)
      } else if tok != nil {
         return adv+1, tok, nil
      }
   }
   return 0, nil, nil
}

func main() {
   r := strings.NewReader(",north,,south,")
   s := bufio.NewScanner(r)
   s.Split(comma)
   for s.Scan() {
      println(s.Text())
   }
}
