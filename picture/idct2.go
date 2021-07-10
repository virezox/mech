package picture

func clamp(x int) int {
   if x < 0 {
      return 0
   }
   if x > 255 {
      return 255
   }
   return x
}

func f2f(x float64) int {
   return int(x * 4096 + 0.5)
}

func fsh(x int) int {
   return x * 4096
}

type transform struct {
   x0 int
   x1 int
   x2 int
   x3 int
}

func idct_1d(s0, s1, s2, s3, s4, s5, s6, s7, []int) transform {
   t0 := make([]int, 1)
   t1 := make([]int, 1)
   t2 := make([]int, 1)
   t3 := make([]int, 1)
   p2 := s2
   p3 := s6
   p1 := (p2[0] + p3[0]) * f2f(0.5411961)
   t2[0] = p1 + p3[0] * f2f(-1.847759065)
   t3[0] = p1 + p2[0] * f2f(0.765366865)
   p2 = s0
   p3 = s4
   t0[0] = fsh(p2[0] + p3[0])
   t1[0] = fsh(p2[0] - p3[0])
   dct := transform{
      x0: t0[0] + t3[0],
      x1: t1[0] + t2[0],
      x2: t1[0] - t2[0],
      x3: t0[0] - t3[0],
   }
   t0 = s7
   t1 = s5
   t2 = s3
   t3 = s1
   p3[0] = t0[0] + t2[0]
   p4 := t1[0] + t3[0]
   p1 = t0[0] + t3[0]
   p2[0] = t1[0] + t2[0]
   p5 := (p3[0] + p4) * f2f(1.175875602)
   t0[0] *= f2f(0.298631336)
   t1[0] *= f2f(2.053119869)
   t2[0] *= f2f(3.072711026)
   t3[0] *= f2f(1.501321110)
   p1 = p5 + p1 * f2f(-0.899976223)
   p2[0] = p5 + p2[0] * f2f(-2.562915447)
   p3[0] *= f2f(-1.961570560)
   p4 *= f2f(-0.390180644)
   t3[0] += p1 + p4
   t2[0] += p2[0] + p3[0]
   t1[0] += p2[0] + p4
   t0[0] += p1 + p3[0]
   return dct
}
