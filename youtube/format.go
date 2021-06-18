package youtube

import (
   "fmt"
   "github.com/robertkrimen/otto"
   "io"
   "net/http"
   "net/url"
   "os"
   "regexp"
)

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   SignatureCipher string
   URL string
}

func (f Format) Write(w io.Writer) error {
   req, err := f.request()
   if err != nil { return err }
   var pos int64
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil { return err }
      defer res.Body.Close()
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("status %v", res.Status)
      }
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

func (f Format) request() (*http.Request, error) {
   if f.URL != "" {
      return http.NewRequest("GET", f.URL, nil)
   }
   baseJS, err := NewBaseJS()
   if err != nil { return nil, err }
   js, err := os.ReadFile(baseJS.Create)
   if err != nil { return nil, err }
   re := `\n[^.]+\.split\(""\);.+`
   child := regexp.MustCompile(re).Find(js)
   if child == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = `\w+`
   childName := regexp.MustCompile(re).Find(child)
   if childName == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = `;(\w+)`
   parentName := regexp.MustCompile(re).FindSubmatch(child)
   if parentName == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   re = fmt.Sprintf(`var %s=.+\n.+\n[^}]+}};`, parentName[1])
   parent := regexp.MustCompile(re).Find(js)
   if parent == nil {
      return nil, fmt.Errorf("find %v", re)
   }
   val, err := url.ParseQuery(f.SignatureCipher)
   if err != nil { return nil, err }
   vm := otto.New()
   if _, err := vm.Run(string(parent) + string(child)); err != nil {
      return nil, err
   }
   sig, err := vm.Call(string(childName), nil, val.Get("s"))
   if err != nil { return nil, err }
   req, err := http.NewRequest("GET", val.Get("url"), nil)
   if err != nil { return nil, err }
   val = req.URL.Query()
   val.Set("sig", sig.String())
   req.URL.RawQuery = val.Encode()
   return req, nil
}
