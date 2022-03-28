package pbs

import (
   "net/url"
   "strconv"
   "strings"
)

func PartnerPlayer(s string) (*url.URL, error) {
   for _, split := range strings.Split(s, "'") {
      if strings.Contains(split, "/partnerplayer/") {
         addr, err := url.Parse(split)
         if err != nil {
            return nil, err
         }
         addr.Scheme = "https"
         return addr, nil
      }
   }
   return nil, notFound{"/partnerplayer/"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
