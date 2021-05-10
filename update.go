package youtube

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "regexp"
)

func getBaseJs(update bool) ([]byte, error) {
   cache, err := os.UserCacheDir()
   if err != nil { return nil, err }
   cache = filepath.Join(cache, "youtube")
   play := filepath.Join(cache, "base.js")
   if update {
      buf, err := httpGet(Origin + "/iframe_api")
      if err != nil { return nil, err }
      re := regexp.MustCompile(`/player\\/\w+`)
      id := re.Find(buf.Bytes())
      if id == nil {
         return nil, fmt.Errorf("Find %v", re)
      }
      base := fmt.Sprintf("/s/player/%s/player_ias.vflset/en_US/base.js", id[9:])
      buf, err = httpGet(Origin + base)
      if err != nil { return nil, err }
      os.Mkdir(cache, os.ModeDir)
      file, err := os.Create(play)
      if err != nil { return nil, err }
      defer file.Close()
      file.ReadFrom(buf)
   }
   return os.ReadFile(play)
}

func (f Format) Write(w io.Writer, update bool) error {
   var req *http.Request
   if f.URL != "" {
      var err error
      req, err = http.NewRequest("GET", f.URL, nil)
      if err != nil { return err }
   } else {
      val, err := url.ParseQuery(f.SignatureCipher)
      if err != nil { return err }
      baseJs, err := getBaseJs(update)
      if err != nil { return err }
      sig, err := decrypt(val.Get("s"), baseJs)
      if err != nil { return err }
      req, err = http.NewRequest("GET", val.Get("url"), nil)
      if err != nil { return err }
      val = req.URL.Query()
      val.Set("sig", sig)
      req.URL.RawQuery = val.Encode()
   }
   var pos int64
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Client).Do(req)
      if err != nil { return err }
      defer res.Body.Close()
      io.Copy(w, res.Body)
      pos += chunk
   }
   return nil
}
