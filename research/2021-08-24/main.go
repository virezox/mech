package main
import "fmt"

func main() {
   m, err := decode(params["SORT BY"][0].encode)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", m)
}
