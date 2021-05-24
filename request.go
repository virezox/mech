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

func decrypt(sig string, js []byte) (string, error) {
   re := `\n[^.]+\.split\(""\);[^\n]+`
   child := regexp.MustCompile(re).Find(js)
   if child == nil {
      return "", fmt.Errorf("find %v", re)
   }
   re = `\w+`
   childName := regexp.MustCompile(re).Find(child)
   if childName == nil {
      return "", fmt.Errorf("find %v", re)
   }
   re = `;\w+`
   parentName := regexp.MustCompile(re).Find(child)
   if parentName == nil {
      return "", fmt.Errorf("find %v", re)
   }
   re = fmt.Sprintf(`var %s=[^\n]+\n[^\n]+\n[^}]+}};`, parentName[1:])
   parent := regexp.MustCompile(re).Find(js)
   if parent == nil {
      return "", fmt.Errorf("find %v", re)
   }
   vm := otto.New()
   vm.Run(string(parent) + string(child))
   value, err := vm.Call(string(childName), nil, sig)
   if err != nil { return "", err }
   return value.String(), nil
   /*
May 19 2021:
var ry={Ui:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
Yg:function(a){a.reverse()},
K6:function(a,b){a.splice(0,b)}};
sy=function(a){a=a.split("");ry.K6(a,2);ry.Yg(a,59);ry.K6(a,3);return a.join("")};
   */
}

func (f Format) Write(w io.Writer) error {
   var req *http.Request
   if f.URL != "" {
      var err error
      req, err = http.NewRequest("GET", f.URL, nil)
      if err != nil { return err }
   } else {
      val, err := url.ParseQuery(f.SignatureCipher)
      if err != nil { return err }
      baseJS, err := NewBaseJS()
      if err != nil { return err }
      create, err := os.ReadFile(baseJS.Create)
      if err != nil { return err }
      sig, err := decrypt(val.Get("s"), create)
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
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("StatusCode %v", res.StatusCode)
      }
      io.Copy(w, res.Body)
      pos += chunk
   }
   return nil
}
