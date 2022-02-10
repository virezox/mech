package vimeo

import (
   "strconv"
   "strings"
   "text/scanner"
)

func scanInt(buf scanner.Scanner) (int, error) {
   for {
      switch buf.Scan() {
      case scanner.Int:
         return strconv.Atoi(buf.TokenText())
      case scanner.EOF:
         return 0, nil
      }
   }
}

type clip struct {
   id int
   unlistedHash int
}

// https://vimeo.com/66531465
// https://vimeo.com/477957994/2282452868
// https://vimeo.com/477957994?unlisted_hash=2282452868
func newClip(address string) (*clip, error) {
   var (
      buf scanner.Scanner
      clipPage clip
      err error
   )
   buf.Init(strings.NewReader(address))
   buf.Mode = scanner.ScanInts
   clipPage.id, err = scanInt(buf)
   if err != nil {
      return nil, err
   }
   clipPage.unlistedHash, err = scanInt(buf)
   if err != nil {
      return nil, err
   }
   return &clipPage, nil
}
