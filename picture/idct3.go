package picture

func idct_block(out []int, stride int, d []int) {
   v := make([]int, 64)
   for i := 0; i < 8; i++ {
      if d[8] | d[16] | d[24] | d[32] | d[40] | d[48] | d[56] == 0 {
         v[0] = d[0] * 4
         v[8] = d[0] * 4
         v[16] = d[0] * 4
         v[24] = d[0] * 4
         v[32] = d[0] * 4
         v[40] = d[0] * 4
         v[48] = d[0] * 4
         v[56] = d[0] * 4
      } else {
         n := newInverseDCT()
         n.transform(d, d[8:], d[16:], d[24:], d[32:], d[40:], d[48:], d[56:])
         n.x0 += 512
         n.x1 += 512
         n.x2 += 512
         n.x3 += 512
         v[0] = (n.x0 + n.t3) >> 10
         v[56] = (n.x0 - n.t3) >> 10
         v[8] = (n.x1 + n.t2) >> 10
         v[48] = (n.x1 - n.t2) >> 10
         v[16] = (n.x2 + n.t1) >> 10
         v[40] = (n.x2 - n.t1) >> 10
         v[24] = (n.x3 + n.t0) >> 10
         v[32] = (n.x3 - n.t0) >> 10
      }
      d = d[1:]
      v = v[1:]
   }
   // FIXME
   for i := 0, o := out; i < 8; ++i, v += 8, o += stride {
      idct_1d(v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])
      x0 += 65536 + (128 << 17)
      x1 += 65536 + (128 << 17)
      x2 += 65536 + (128 << 17)
      x3 += 65536 + (128 << 17)
      o[0] = clamp((x0 + t3) >> 17)
      o[7] = clamp((x0 - t3) >> 17)
      o[1] = clamp((x1 + t2) >> 17)
      o[6] = clamp((x1 - t2) >> 17)
      o[2] = clamp((x2 + t1) >> 17)
      o[5] = clamp((x2 - t1) >> 17)
      o[3] = clamp((x3 + t0) >> 17)
      o[4] = clamp((x3 - t0) >> 17)
   }
}
