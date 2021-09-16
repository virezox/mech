package bandcamp

func dm(String p8, int p9) string {
   char[] v8_1 = p8.toCharArray()
   int v0 = v8_1.length
   int v2 = 0
   if ((p9 % 2) != 0) {
      int v1_1 = 3
      while ((v1_1 * v1_1) <= p9) {
         if ((p9 % v1_1) != 0) {
            v1_1 += 2
         } else {
            int v1_0 = 0
         }
         if (v1_0 != 0) {
            p9++
         }
      }
      v1_0 = 1
   }
   int v1_3 = new boolean[(p9 + 1)]
   v1_3[0] = 1
   v1_3[1] = 1
   int v4_1 = -1
   int v5_0 = -1
   int v6 = 2
   while (v6 <= p9) {
      if (v1_3[v6] == 0) {
         int v4_9 = (v6 + v6)
         while (v4_9 <= p9) {
            v1_3[v4_9] = 1
            v4_9 += v6
         }
         v4_1 = v5_0
         v5_0 = v6
      }
      v6++
   }
   String v9_1 = (10 - (v4_1 % 10))
   int v1_4 = (26 - (v5_0 % 26))
   while (v2 < v0) {
      int v3_0
      v3_0 = v8_1[v2]
      if ((v3_0 < 97) || (v3_0 > 122)) {
         if ((v3_0 < 65) || (v3_0 > 90)) {
            if ((v3_0 >= 48) && (v3_0 <= 57)) {
               v3_0 += v9_1
               if (v3_0 > 57) {
                  v3_0 -= 10
               }
            }
         } else {
            v3_0 += v1_4
            if (v3_0 > 90) {
               v3_0 -= 26
            }
         }
      } else {
         v3_0 += v1_4
      }
      v8_1[v2] = ((char) v3_0)
      v2++
   }
   return new String(v8_1)
}
