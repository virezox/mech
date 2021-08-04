package main
import "fmt"

func powerSet(a ...string) [][]string {
   b := make([][]string, 1)
   for _, c := range a {
      for _, d := range b {
         b = append(b, append(d, c))
      }
   }
   return b
}

func main() {
   for _, a := range powerSet("index", "duration", "size", "title") {
      fmt.Println(a)
   }
}
