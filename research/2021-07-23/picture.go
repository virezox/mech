package picture

import (
   "image"
   "image/jpeg"
   "os"
)

func absDiff(a, b uint32) uint32 {
   if a < b {
      return b - a
   }
   return a - b
}

func (a slow) difference(b *slow) float32 {
   var (
      part uint32
      total uint64
   )
   r := a.Bounds()
   for x := r.Min.X; x < r.Max.X; x++ {
      for y := r.Min.Y; y < r.Max.Y; y++ {
         r1, g1, b1, _ := a.At(x, y).RGBA()
         r2, g2, b2, _ := b.At(x, y).RGBA()
         part += absDiff(r1, r2)
         part += absDiff(g1, g2)
         part += absDiff(b1, b2)
         total += 0xFFFF + 0xFFFF + 0xFFFF
      }
   }
   return float32(part) / float32(total)
}

type slow struct {
   image.Image
}

func newPicture(name string) (*slow, error) {
   f, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   i, err := jpeg.Decode(f)
   if err != nil {
      return nil, err
   }
   return &slow{i}, nil
}

