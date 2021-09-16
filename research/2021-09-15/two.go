package main
import "fmt"

const body =
   "client_id=134&" +
   "client_secret=1myK12VeCL3dWl9o%2FncV2VyUUbOJuNPVJK6bZZJxHvk%3D&" +
   "grant_type=password&password=PASSWORD&username=4095486538&" +
   "username_is_user_id=1"

func main() {
   s := dm(body, 9)
   fmt.Println(s)
}

func dm(sInput string, nInput int) string {
   cInput := []byte(sInput)
   if nInput % 2 != 0 {
      n11 := 3
      // FIXME
      for n11 * n11 <= nInput {
         if nInput % n11 != 0 {
            n11 += 2
         }
      }
   }
   b13 := make([]bool, nInput+1)
   b13[0] = true
   b13[1] = true
   n41 := -1
   n50 := -1
   for n6 := 2; n6 <= nInput; n6++ {
      if b13[n6] == false {
         n49 := n6 + n6
         for n49 <= nInput {
            b13[n49] = true
            n49 += n6
         }
         n41 = n50
         n50 = n6
      }
   }
   n91 := 10 - (n41 % 10)
   n14 := 26 - (n50 % 26)
   for n2 := 0; n2 < len(cInput); n2++ {
      c30 := cInput[n2]
      if c30 < 97 || c30 > 122 {
         if c30 < 65 || c30 > 90 {
            if c30 >= 48 && c30 <= 57 {
               c30 += byte(n91)
               if c30 > 57 {
                  c30 -= 10
               }
            }
         } else {
            c30 += byte(n14)
            if c30 > 90 {
               c30 -= 26
            }
         }
      } else {
         c30 += byte(n14)
      }
      cInput[n2] = c30
   }
   return string(cInput)
}
