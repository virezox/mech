package main
import "fmt"

func main() {
   for _, p := range params["UPLOAD DATE"] {
      m, err := decode(p.encode)
      if err != nil {
         panic(err)
      }
      fmt.Printf("%+v\n", m)
      s := encode(p.decode)
      fmt.Println(s)
   }
}
