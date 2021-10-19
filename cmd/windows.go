package main

import (
   "os"
   "os/exec"
)

func main() {
   args := []string{"build", "-ldflags", "-s", "-o", "."}
   dirs, err := os.ReadDir(".")
   if err != nil {
      panic(err)
   }
   for _, dir := range dirs {
      if dir.IsDir() {
         args = append(args, "./" + dir.Name())
      }
   }
   if err := exec.Command("go", args...).Run(); err != nil {
      panic(err)
   }
}
