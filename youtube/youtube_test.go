package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func Test_ID(t *testing.T) {
   for _, test := range id_tests {
      id, err := Video_ID(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}

const image_test = "UpNXI3_ctAc"

func Test_Image(t *testing.T) {
   for _, img := range Images {
      addr := img.Format(image_test)
      fmt.Println("HEAD", addr)
      res, err := http.Head(addr)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
