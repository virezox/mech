package youtube

import (
   "bufio"
   "fmt"
   "net/http"
   "regexp"
   "strconv"
   "strings"
)

const (
   /*
   \/player\/bce81a70\/
   */
   player = `/player\\/(\w+)\\`
   split = `.split("");`
)

func swap(sig []byte, index int) {
   c := sig[0]
   sig[0] = sig[index % len(sig)]
   sig[index % len(sig)] = c
}

func decrypt(sig []byte) error {
   body, err := readAll("https://www.youtube.com/iframe_api")
   if err != nil { return err }
   id := regexp.MustCompile(player).FindSubmatch(body)
   base := fmt.Sprintf(
      "https://www.youtube.com/s/player/%s/player_ias.vflset/en_US/base.js",
      id[1],
   )
   println("Get", base)
   res, err := http.Get(base)
   if err != nil { return err }
   defer res.Body.Close()
   scan := bufio.NewScanner(res.Body)
   for scan.Scan() {
      if ! strings.Contains(scan.Text(), split) { continue }
      for _, match := range regexp.MustCompile(`\d+`).FindAllString(scan.Text(), -1) {
         index, err := strconv.Atoi(match)
         if err != nil { return err }
         swap(sig, index)
      }
      return nil
   }
   return fmt.Errorf("%q not found", split)
}
