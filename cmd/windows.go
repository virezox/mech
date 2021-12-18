package main

import (
   "fmt"
   "os"
   "os/exec"
)

func main() {
   if len(os.Args) == 1 {
      fmt.Println("cmd [folder]...")
      return
   }
   arg := []string{"build", "-ldflags", "-s", "-o", "."}
   for _, dir := range os.Args[1:] {
      arg = append(arg, dir)
   }
   if err := exec.Command("go", arg...).Run(); err != nil {
      panic(err)
   }
}
