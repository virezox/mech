package main

import (
   "encoding/json"
   "fmt"
   "strings"
)

const jsonStream = `
= {"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`

func main() {
   r := strings.NewReader(jsonStream)
   dec := json.NewDecoder(r)
   // this consumes the whole reader
   fmt.Println(dec.Token())
   fmt.Println(r.ReadByte())
}
