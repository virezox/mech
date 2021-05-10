package youtube

import (
   "fmt"
   "github.com/robertkrimen/otto"
   "os"
   "path/filepath"
   "regexp"
)

func decrypt(sig string, js []byte) (string, error) {
   child, err := find(`\n[^.]+\.split\(""\);[^\n]+`, js)
   if err != nil { return "", err }
   // child name
   childName, err := find(`\w+`, child)
   if err != nil { return "", err }
   // parent name
   parentName, err := find(`;\w+`, child)
   if err != nil { return "", err }
   parentName = parentName[1:]
   // parent
   parent, err := find(
      fmt.Sprintf(`var %s=[^\n]+\n[^\n]+\n[^}]+}};`, parentName), js,
   )
   if err != nil { return "", err }
   // run
   vm := otto.New()
   vm.Run(string(parent) + string(child))
   value, err := vm.Call(string(childName), nil, sig)
   if err != nil {
      return "", fmt.Errorf("parent %q %v", parent, err)
   }
   return value.String(), nil
   /*
May 7 2021:
var uy={wd:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
kI:function(a){a.reverse()},
NY:function(a,b){a.splice(0,b)}};wy.prototype.set=function(a,b){this.i[a]!==b&&(this.i[a]=b,this.url="")};
vy=function(a){a=a.split("");uy.wd(a,41);uy.NY(a,3);uy.kI(a,41);uy.NY(a,2);uy.kI(a,5);uy.wd(a,62);uy.NY(a,3);uy.wd(a,69);uy.NY(a,2);return a.join("")};

May 5 2021:
var uy={bH:function(a,b){a.splice(0,b)},
Fg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
S6:function(a){a.reverse()}};
vy=function(a){a=a.split("");uy.bH(a,3);uy.Fg(a,7);uy.Fg(a,50);uy.S6(a,71);uy.bH(a,2);uy.S6(a,80);uy.Fg(a,38);return a.join("")};

May 4 2021:
var uy={an:function(a){a.reverse()},
gN:function(a,b){a.splice(0,b)},
J4:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c}};
vy=function(a){a=a.split("");uy.gN(a,2);uy.J4(a,47);uy.gN(a,1);uy.an(a,49);uy.gN(a,2);uy.J4(a,4);uy.an(a,71);uy.J4(a,15);uy.J4(a,40);return a.join("")};

May 3 2021:
var uy={VP:function(a){a.reverse()},
eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
li:function(a,b){a.splice(0,b)}};
vy=function(a){a=a.split("");uy.eG(a,50);uy.eG(a,48);uy.eG(a,23);uy.eG(a,31);return a.join("")};
   */
}

func find(pat string, sub []byte) ([]byte, error) {
   re, err := regexp.Compile(pat)
   if err != nil { return nil, err }
   match := re.Find(sub)
   if match == nil {
      return nil, fmt.Errorf("find %v", pat)
   }
   return match, nil
}


func getBaseJs(update bool) ([]byte, error) {
   cache, err := os.UserCacheDir()
   if err != nil { return nil, err }
   cache = filepath.Join(cache, "youtube")
   play := filepath.Join(cache, "base.js")
   if update {
      buf, err := httpGet(Origin + "/iframe_api")
      if err != nil { return nil, err }
      id, err := find(`/player\\/\w+`, buf.Bytes())
      if err != nil { return nil, err }
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
