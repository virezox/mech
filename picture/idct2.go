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

type inverseDCT struct {
   t0, t1, t2, t3 []int
   x0, x1, x2, x3 int
}

func newInverseDCT() inverseDCT {
   var n inverseDCT
   n.t0 = make([]int, 1)
   n.t1 = make([]int, 1)
   n.t2 = make([]int, 1)
   n.t3 = make([]int, 1)
   return n
}

func (n *inverseDCT) transform(s0, s1, s2, s3, s4, s5, s6, s7 []int) {
   p2 := s2
   p3 := s6
   p1 := (p2[0] + p3[0]) * f2f(0.5411961)
   n.t2[0] = p1 + p3[0] * f2f(-1.847759065)
   n.t3[0] = p1 + p2[0] * f2f(0.765366865)
   p2 = s0
   p3 = s4
   n.t0[0] = fsh(p2[0] + p3[0])
   n.t1[0] = fsh(p2[0] - p3[0])
   n.x0 = n.t0[0] + n.t3[0]
   n.x3 = n.t0[0] - n.t3[0]
   n.x1 = n.t1[0] + n.t2[0]
   n.x2 = n.t1[0] - n.t2[0]
   n.t0 = s7
   n.t1 = s5
   n.t2 = s3
   n.t3 = s1
   p3[0] = n.t0[0] + n.t2[0]
   p4 := n.t1[0] + n.t3[0]
   p1 = n.t0[0] + n.t3[0]
   p2[0] = n.t1[0] + n.t2[0]
   p5 := (p3[0] + p4) * f2f(1.175875602)
   n.t0[0] *= f2f(0.298631336)
   n.t1[0] *= f2f(2.053119869)
   n.t2[0] *= f2f(3.072711026)
   n.t3[0] *= f2f(1.501321110)
   p1 = p5 + p1 * f2f(-0.899976223)
   p2[0] = p5 + p2[0] * f2f(-2.562915447)
   p3[0] *= f2f(-1.961570560)
   p4 *= f2f(-0.390180644)
   n.t3[0] += p1 + p4
   n.t2[0] += p2[0] + p3[0]
   n.t1[0] += p2[0] + p4
   n.t0[0] += p1 + p3[0]
}
