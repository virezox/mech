package tls

import (
   "bufio"
   "fmt"
   "net/http"
   "strings"
)

const param =
   "https://www.iana.org/assignments/tls-parameters/tls-parameters.txt"

type iana map[string]string

func newIana() (*iana, error) {
   fmt.Println("GET", param)
   r, err := http.Get(param)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   s := bufio.NewScanner(r.Body)
   m := make(iana)
   for s.Scan() {
      field := strings.Fields(s.Text())
      if len(field) >= 2 {
         m[field[0]] = field[1]
      }
   }
   return &m, nil
}

func (i iana) get(key uint16) string {
   hi := byte(key >> 8)
   lo := byte(key)
   return i[fmt.Sprintf("0x%02X,0x%02X", hi, lo)]
}
