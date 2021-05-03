package main

import (
   "fmt"
   "github.com/robertkrimen/otto"
)

const js = `
var uy={VP:function(a){a.reverse()},
eG:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
li:function(a,b){a.splice(0,b)}};
vy=function(a){a=a.split("");uy.eG(a,50);uy.eG(a,48);uy.eG(a,23);uy.eG(a,31);return a.join("")};
`

const (
in = "POq0QJ8wRgIhAKcjYNqd8ijVquRn576AWqtlYRGdetJecLMpfraxyWqvAiEApZ6J5XQ-o_gOMKsvMMZS9I-N2baHgzMSWN8tYoYUYHI="
out = "AOq0QJ8wRgIhAKcjYNqd8ijfquRn576VWqtlYRGdetJecLMparPxyWqvAiEApZ6J5XQ-o_gOMKsvMMZS9I-N2baHgzMSWN8tYoYUYHI="
)

func main() {
   vm := otto.New()
   vm.Run(js)
   value, err := vm.Call("vy", nil, in)
   if err != nil {
      panic(err)
   }
   fmt.Println(value.String() == out)
}
