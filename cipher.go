package youtube

import (
   "errors"
   "fmt"
   "regexp"
   "sort"
   "strconv"
)

const jsvarStr = `[a-zA-Z_\$][a-zA-Z_0-9]*`

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

func (v Video) baseJs() ([]byte, error) {
   body, err := readAll("https://www.youtube.com/embed/" + v.VideoDetails.VideoId)
   if err != nil { return nil, err }
   player := regexp.MustCompile("/player/([^/]+)/player_").FindSubmatch(body)
   if len(player) < 2 {
      return nil, errors.New("unable to find basejs URL in playerConfig")
   }
   return readAll(fmt.Sprintf(
      "https://www.youtube.com/s/player/%s/player_ias.vflset/en_US/base.js",
      player[1],
   ))
}

func findFuncBody(body []byte) []byte {
   f := `function(?: %s)?\(a\)\{a=a\.split\(""\);\s*((?:(?:a=)?%s\.%s\(a,\d+\);)+)return a\.join\(""\)\}`
   return regexp.MustCompile(fmt.Sprintf(f, jsvarStr, jsvarStr, jsvarStr)).FindSubmatch(body)[1]
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
   ci.regex = fmt.Sprintf(
      `(?:a=)?%s\.(%s|%s|%s)\(a,(\d+)\)`, obj, ci.reverse, ci.splice, ci.swap,
   )
   return ci
}

type cipher struct {
   regex string
   reverse string
   splice string
   swap string
}

func (v Video) parseDecipherOps() ([]decipherOperation, error) {
   body, err := v.baseJs()
   if err != nil { return nil, err }
   funcBody := findFuncBody(body)
   ci := newCipher(body)
   var ops []decipherOperation
   for _, s := range regexp.MustCompile(ci.regex).FindAllSubmatch(funcBody, -1) {
      switch string(s[1]) {
      case ci.swap:
         arg, _ := strconv.Atoi(string(s[2]))
         ops = append(ops, newSwapFunc(arg))
      case ci.splice:
         arg, _ := strconv.Atoi(string(s[2]))
         ops = append(ops, newSpliceFunc(arg))
      case ci.reverse:
         ops = append(ops, newReverseFunc())
      }
   }
   return ops, nil
}
