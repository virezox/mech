package youtube
import "strconv"

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
