package youtube

import (
   "fmt"
   "regexp"
   "strconv"
)

type cipher struct {
   matches [][]string
   reverse string
   splice string
   swap string
}

func (ci cipher) decrypt(sig []byte) error {
   for _, match := range ci.matches {
      switch match[1] {
      case ci.swap:
         arg, err := strconv.Atoi(match[2])
         if err != nil { return err }
         pos := arg % len(sig)
         sig[0], sig[pos] = sig[pos], sig[0]
      case ci.splice:
         arg, err := strconv.Atoi(match[2])
         if err != nil { return err }
         sig = sig[arg:]
      case ci.reverse:
         for n := len(sig) - 2; n >= 0; n-- {
            sig = append(sig[:n], append(sig[n + 1:], sig[n])...)
         }
      }
   }
   return nil
}

const (
   jsReverse = `:function\(a\)\{(?:return )?a\.reverse\(\)\}`
   jsSplice = `:function\(a,b\)\{a\.splice\(0,b\)\}`
   jsSwap = `:function\(a,b\)\{var c=a\[0\];a\[0\]=a\[b(?:%a\.length)?\];a\[b(?:%a\.length)?\]=c(?:;return a)?\}`
   jsVar = `[a-zA-Z_\$][a-zA-Z_0-9]*`
)

func newCipher(body []byte) cipher {
   /*
   var uy={VP:function(a){a.reverse()},
   eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
   li:function(a,b){a.splice(0,b)}};
   */
   objResult := regexp.MustCompile(fmt.Sprintf(
      `var (%v)=\{((?:(?:%v%v|%v%v|%v%v),?\n?)+)\};`,
      jsVar, jsVar, jsSwap, jsVar, jsSplice, jsVar, jsReverse,
   )).FindSubmatch(body)
   obj := objResult[1]
   objBody := objResult[2]
   var ci cipher
   result := regexp.MustCompile(fmt.Sprintf(
      "(?m)(?:^|,)(%v)%v", jsVar, jsReverse,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.reverse = string(result[1])
   }
   result = regexp.MustCompile(
      fmt.Sprintf("(?m)(?:^|,)(%v)%v", jsVar, jsSplice,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.splice = string(result[1])
   }
   result = regexp.MustCompile(fmt.Sprintf(
      "(?m)(?:^|,)(%v)%v", jsVar, jsSwap,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.swap = string(result[1])
   }
   funcBody := regexp.MustCompile(fmt.Sprintf(
      `function(?: %v)?\(a\)\{a=a\.split\(""\);\s*((?:(?:a=)?%v\.%v\(a,\d+\);)+)return a\.join\(""\)\}`,
      jsVar, jsVar, jsVar,
   )).FindSubmatch(body)[1]
   ci.matches = regexp.MustCompile(fmt.Sprintf(
      `(?:a=)?%s\.(%v|%v|%v)\(a,(\d+)\)`, obj, ci.reverse, ci.splice, ci.swap,
   )).FindAllStringSubmatch(string(funcBody), -1)
   return ci
}
