package main

import (
   "bufio"
   "strings"
)

const ww = "Wonder.Woman.1984.2020.IMAX.1080p.WEBRip.x265-RARBG"

func dot(data []byte, eof bool) (int, []byte, error) {
   var tok []byte
   for i, t := range data {
      if t != '.' {
         tok = append(tok, t)
      } else if tok != nil {
         return i+1, tok, nil
      }
   }
   return 0, nil, nil
}

func main() {
   r := strings.NewReader(ww)
   s := bufio.NewScanner(r)
   s.Split(dot)
   s.Scan()
   println(s.Text())
   r.Reset(ww)
   s.Scan()
   println(s.Text())
}
