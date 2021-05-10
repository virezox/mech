package youtube

import (
   "fmt"
   "os"
   "path/filepath"
   "regexp"
)

type BaseJS struct {
   Cache string
   Create string
}

func NewBaseJS() (BaseJS, error) {
   var (
      b BaseJS
      err error
   )
   b.Cache, err = os.UserCacheDir()
   if err != nil {
      return BaseJS{}, err
   }
   b.Cache = filepath.Join(b.Cache, "youtube")
   b.Create = filepath.Join(b.Cache, "base.js")
   return b, nil
}

func (b BaseJS) Get() error {
   buf, err := httpGet(Origin + "/iframe_api")
   if err != nil { return err }
   re := regexp.MustCompile(`/player\\/\w+`)
   id := re.Find(buf.Bytes())
   if id == nil {
      return fmt.Errorf("Find %v", re)
   }
   get := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[9:])
   buf, err = httpGet(Origin + get)
   if err != nil { return err }
   os.Mkdir(b.Cache, os.ModeDir)
   file, err := os.Create(b.Create)
   if err != nil { return err }
   defer file.Close()
   _, err = file.ReadFrom(buf)
   return err
}
