package main
import "fmt"

func main() {
   for k, v := range params["FEATURES"] {
      val := encode(v)
      fmt.Print(k, "\n", val, "\n\n")
   }
}
