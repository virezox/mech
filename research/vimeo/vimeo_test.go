package vimeo

import (
   "fmt"
   "testing"
)

var tests = []string{
   "https://embed.vhx.tv/subscriptions/28599?vimeo=1",
   "https://embed.vhx.tv/videos/1264265?vimeo=1",
}

func Test_Embed(t *testing.T) {
   fmt.Println(tests)
}
