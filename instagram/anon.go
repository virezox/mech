package instagram

import (
   "github.com/89z/format"
)

var LogLevel format.LogLevel

// instagram.com/p/CT-cnxGhvvO
// instagram.com/p/yza2PAPSx2
func Valid(shortcode string) bool {
   switch len(shortcode) {
   case 10, 11:
      return true
   }
   return false
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
