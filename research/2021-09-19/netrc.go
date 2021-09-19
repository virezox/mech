package github

import (
   "fmt"
   "os"
)

func netrc() (string, string, error) {
   home, err := os.UserHomeDir()
   if err != nil { return "", "", err }
   f, err := os.Open(home + "/_netrc")
   if err != nil { return "", "", err }
   defer f.Close()
   var login, pass string
   fmt.Fscanf(f, "default login %v password %v", &login, &pass)
   return login, pass, nil
}
