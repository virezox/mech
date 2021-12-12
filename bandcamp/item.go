package bandcamp

import (
   "github.com/89z/mech"
   "net/http"
   "regexp"
   "strconv"
)

// URL to Item. Request is anonymous.
func NewItem(addr string) (*Item, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return nil, err
   }
   if req.URL.Path == "" {
      req.URL.Path = "/music"
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   // [nilZ0t2809477874x t 2809477874]
   reg := regexp.MustCompile(`nilZ0([ait])(\d+)x`)
   for _, c := range res.Cookies() {
      if c.Name == "session" {
         find := reg.FindStringSubmatch(c.Value)
         if find != nil {
            id, err := strconv.Atoi(find[2])
            if err == nil {
               var item Item
               item.Item_Type = find[1]
               item.Item_ID = id
               return &item, nil
            }
         }
      }
   }
   return nil, mech.NotFound{"session"}
}
