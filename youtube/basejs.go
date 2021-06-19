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
   req, err := http.NewRequest("GET", Origin + "/iframe_api", nil)
   if err != nil {
      return err
   }
   fmt.Println(invert, "GET", reset, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   re := regexp.MustCompile(`/player\\/(\w+)`)
   id := re.FindSubmatch(body)
   if id == nil {
      return fmt.Errorf("FindSubmatch %v", re)
   }
   os.Mkdir(b.Cache, os.ModeDir)
   file, err := os.Create(b.Create)
   if err != nil {
      return err
   }
   defer file.Close()
   req.URL.Path = fmt.Sprintf(
      "/s/player/%s/player_ias.vflset/en_US/base.js", id[1],
   )
   fmt.Println(invert, "GET", reset, req.URL)
   if res, err := new(http.Transport).RoundTrip(req); err != nil {
      return err
   } else {
      defer res.Body.Close()
      if res.StatusCode != http.StatusOK {
         return fmt.Errorf("status %v", res.Status)
      }
      file.ReadFrom(res.Body)
   }
   return nil
}
