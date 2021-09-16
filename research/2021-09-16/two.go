package bandcamp

func dm(sInput string, nInput int) string {
   bInput := []byte(sInput)
   if nInput % 2 != 0 {
      v11 := 3
      for v11 * v11 <= nInput {
         if nInput % v11 != 0 {
            v11 += 2
         } else {
            int v10 = 0
         }
         if v10 != 0 {
            nInput++
         }
      }
      v10 = 1
   }
   int v13 = new boolean[nInput + 1]
   v13[0] = 1
   v13[1] = 1
   int v41 = -1
   int v50 = -1
   int v6 = 2
   for v6 <= nInput {
      if v13[v6] == 0 {
         int v49 = v6 + v6
         for v49 <= nInput {
            v13[v49] = 1
            v49 += v6
         }
         v41 = v50
         v50 = v6
      }
      v6++
   }
   String v91 = 10 - (v41 % 10)
   int v14 = 26 - (v50 % 26)
   for v2 := 0; v2 < len(bInput); v2++ {
      int v30 = bInput[v2]
      if v30 < 97 || v30 > 122 {
         if v30 < 65 || v30 > 90 {
            if v30 >= 48 && v30 <= 57 {
               v30 += v91
               if v30 > 57 {
                  v30 -= 10
               }
            }
         } else {
            v30 += v14
            if v30 > 90 {
               v30 -= 26
            }
         }
      } else {
         v30 += v14
      }
      bInput[v2] = (char) v30
   }
   return string(bInput)
}
