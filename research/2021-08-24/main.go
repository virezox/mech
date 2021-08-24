package main
import "fmt"

func main() {
   for _, p := range params["FEATURES"] {
      s := encode(p.decode)
      fmt.Println(s)
   }
}
