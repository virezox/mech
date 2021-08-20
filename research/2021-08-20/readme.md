# August 20 2021

~~~go
// state 2: break if !WhitespaceToken
for {
   if tt, data := l.Next(); tt != js.WhitespaceToken {
      if tt == js.DivToken {
         tt, data = l.RegExp()
      }
      vals[k] = data
      break
   }
}
// state 3: break if SemicolonToken
for {
   tt, data := l.Next()
   if tt == js.SemicolonToken {
      break
   }
   vals[k] = append(vals[k], data...)
}
~~~
