package main

import (
   "os"
   "os/exec"
)

func main() {
   dirs, err := os.ReadDir(".")
   if err != nil {
      panic(err)
   }
   args := []string{"build", "-ldflags", "-s", "-o", "darwin"}
   for _, dir := range dirs {
      if dir.IsDir() {
         args = append(args, "./" + dir.Name())
      }
   }
   os.Mkdir("darwin", os.ModeDir)
   cmd := exec.Command("go", args...)
   cmd.Env = []string{
      "GOOS=darwin",
      `GOCACHE=C:\Users\Steven\AppData\Local\go-build`,
      `GOPATH=C:\Users\Steven\go`,
      `TMP=C:\Windows\TEMP`,
   }
   if err := cmd.Run(); err != nil {
      panic(err)
   }
}
