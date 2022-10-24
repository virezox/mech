package paramount

import (
   "testing"
   "time"
)

var app_secrets = map[int]string{
   0: "439ba2e3622c344a", // 12.0.28
   1: "79b7e56e442e65ed", // 12.0.27
   2: "f012987182d6f16c", // 12.0.26
   3: "d0795c0dffebea73", // 8.1.28
   4: "a75bd3a39bfcbc77", // 8.1.26
   5: "c0966212aa651e8b", // 8.1.23
   6: "ddca2f16bfa3d937", // 8.1.22
   7: "817774cbafb2b797", // 8.1.20
   8: "1705732089ff4d60", // 8.1.18
   9: "add603b54be2fc3c", // 8.1.16
}

func Test_Session(t *testing.T) {
   for _, secret := range app_secrets {
      err := New_Session(secret)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
