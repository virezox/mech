package bandcamp

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "text/scanner"
)

type Item struct {
   Item_Type string
   Item_ID int
}

func NewItem(addr string) (*Item, error) {
   req, err := http.NewRequest("HEAD", addr, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   var (
      scan scanner.Scanner
      item Item
   )
   for _, cook := range res.Cookies() {
      if cook.Name == "session" {
         sess, err := url.QueryUnescape(cook.Value)
         if err != nil {
            return nil, err
         }
         scan.Init(strings.NewReader(sess))
         scan.IsIdentRune = func(r rune, i int) bool {
            return r >= 'A'
         }
         scan.Mode = scanner.ScanIdents | scanner.ScanInts
         for scan.Scan() != scanner.EOF {
            if scan.TokenText() == "nilZ" {
               scan.Scan()
               scan.Scan()
               item.Item_Type = scan.TokenText()
               scan.Scan()
               item.Item_ID, err = strconv.Atoi(scan.TokenText())
               if err != nil {
                  return nil, err
               }
               return &item, nil
            }
         }
      }
   }
   return nil, notFound{"nilZ"}
}

func (i Item) Band() (*Band, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/band_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "band_id=" + strconv.Itoa(i.Item_ID)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   band := new(Band)
   if err := json.NewDecoder(res.Body).Decode(band); err != nil {
      return nil, err
   }
   return band, nil
}

func (i Item) Tralbum() (*Tralbum, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "band_id": {"1"},
      "tralbum_id": {strconv.Itoa(i.Item_ID)},
      "tralbum_type": {i.Type()},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tralb := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tralb); err != nil {
      return nil, err
   }
   return tralb, nil
}

func (i Item) Type() string {
   for _, char := range i.Item_Type {
      return string(char)
   }
   return ""
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
