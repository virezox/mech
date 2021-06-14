package youtube

import (
   "fmt"
   "io"
   "net/http"
   "os"
   "path/filepath"
   "regexp"
)

type BaseJS struct {
   Cache string
   Create string
}

func NewBaseJS() (BaseJS, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return BaseJS{}, err
   }
   cache = filepath.Join(cache, "mech")
   return BaseJS{
      cache, filepath.Join(cache, "base.js"),
   }, nil
}

func (b BaseJS) Get() error {
   fmt.Println(invert, "GET", reset, Origin + "/iframe_api")
   res, err := http.Get(Origin + "/iframe_api")
   if err != nil { return err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return err }
   re := regexp.MustCompile(`/player\\/(\w+)`)
   id := re.FindSubmatch(body)
   if id == nil {
      return fmt.Errorf("FindSubmatch %v", re)
   }
   os.Mkdir(b.Cache, os.ModeDir)
   file, err := os.Create(b.Create)
   if err != nil { return err }
   defer file.Close()
   get := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[1])
   fmt.Println(invert, "GET", reset, Origin + get)
   if res, err := http.Get(Origin + get); err != nil {
      return err
   } else {
      defer res.Body.Close()
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
   }
   return nil
}
