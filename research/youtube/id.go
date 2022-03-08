package youtube

import (
   "net/url"
)

// https://youtube.com/shorts/9Vsdft81Q6w
// https://youtube.com/watch?v=XY-hOqcPGCY
func VideoID(address string) (string, error) {
   addr, err := url.Parse(address)
   if err != nil {
      return "", err
   }
   return addr.Query().Get("v"), nil
}
