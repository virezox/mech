package mech

import (
   "path/filepath"
   "strings"
)

type Cleaner struct {
   name string
}

func Clean(dir, file string) Cleaner {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return -1
      }
      return r
   }
   var c Cleaner
   c.name = strings.Map(mapping, file)
   c.name = filepath.Join(dir, c.name)
   return c
}
