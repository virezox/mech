package youtube

import (
   "fmt"
   "regexp"
   "sort"
   "strconv"
)

const jsvarStr = `[a-zA-Z_\$][a-zA-Z_0-9]*`

type cipher struct {
   matches [][]string
   reverse string
   splice string
   swap string
}

func newCipher(body []byte) cipher {
   reverseStr := `:function\(a\)\{(?:return )?a\.reverse\(\)\}`
   spliceStr := `:function\(a,b\)\{a\.splice\(0,b\)\}`
   swapStr := `:function\(a,b\)\{var c=a\[0\];a\[0\]=a\[b(?:%a\.length)?\];a\[b(?:%a\.length)?\]=c(?:;return a)?\}`
   objResult := regexp.MustCompile(fmt.Sprintf(
      `var (%s)=\{((?:(?:%s%s|%s%s|%s%s),?\n?)+)\};`,
      jsvarStr, jsvarStr, swapStr, jsvarStr, spliceStr, jsvarStr, reverseStr,
   )).FindSubmatch(body)
   obj := objResult[1]
   objBody := objResult[2]
   var ci cipher
   result := regexp.MustCompile(fmt.Sprintf(
      "(?m)(?:^|,)(%s)%s", jsvarStr, reverseStr,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.reverse = string(result[1])
   }
   result = regexp.MustCompile(
      fmt.Sprintf("(?m)(?:^|,)(%s)%s", jsvarStr, spliceStr,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.splice = string(result[1])
   }
   result = regexp.MustCompile(fmt.Sprintf(
      "(?m)(?:^|,)(%s)%s", jsvarStr, swapStr,
   )).FindSubmatch(objBody)
   if len(result) > 1 {
      ci.swap = string(result[1])
   }
   funcBody := regexp.MustCompile(fmt.Sprintf(
      `function(?: %s)?\(a\)\{a=a\.split\(""\);\s*((?:(?:a=)?%s\.%s\(a,\d+\);)+)return a\.join\(""\)\}`,
      jsvarStr, jsvarStr, jsvarStr,
   )).FindSubmatch(body)[1]
   ci.matches = regexp.MustCompile(fmt.Sprintf(
      `(?:a=)?%s\.(%s|%s|%s)\(a,(\d+)\)`, obj, ci.reverse, ci.splice, ci.swap,
   )).FindAllStringSubmatch(string(funcBody), -1)
   return ci
}

type decipherOperation func([]byte) []byte

func newReverseFunc() decipherOperation {
   return func(bs []byte) []byte {
      sort.Slice(bs, func(d, e int) bool { return true })
      return bs
   }
}

func newSpliceFunc(pos int) decipherOperation {
   return func(bs []byte) []byte {
      return bs[pos:]
   }
}

func newSwapFunc(arg int) decipherOperation {
   return func(bs []byte) []byte {
      pos := arg % len(bs)
      bs[0], bs[pos] = bs[pos], bs[0]
      return bs
   }
}

func parseDecipherOps(sig, body []byte) []byte {
   ci := newCipher(body)
   for _, s := range ci.matches {
      switch s[1] {
      case ci.swap:
         arg, _ := strconv.Atoi(s[2])
         sig = newSwapFunc(arg)(sig)
      case ci.splice:
         arg, _ := strconv.Atoi(s[2])
         sig = newSpliceFunc(arg)(sig)
      case ci.reverse:
         sig = newReverseFunc()(sig)
      }
   }
   return sig
}
