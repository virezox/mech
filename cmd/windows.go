package main

import (
   "os"
   "os/exec"
)

func main() {
   arg := []string{"build", "-ldflags", "-s", "-o", "."}
   dirs, err := os.ReadDir(".")
   if err != nil {
      panic(err)
   }
   for _, dir := range dirs {
      if dir.IsDir() {
         arg = append(arg, "./" + dir.Name())
      }
   }
   if err := exec.Command("go", arg...).Run(); err != nil {
      panic(err)
   }
}
