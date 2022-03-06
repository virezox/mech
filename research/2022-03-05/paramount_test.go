package paramount

import (
   "fmt"
   "testing"
)

var addrs = []string{
   "paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   "paramountplus.com/shows/bull/video/TUT_4UVB87huHEOfPCjMkxOW_Xe1hNWw/bull-gone",
}

func TestBufio(t *testing.T) {
   for _, addr := range addrs {
      dst := doBufio(addr)
      fmt.Println(dst)
   }
}
