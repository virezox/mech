package main
 
import (
   "fmt"
   "image/jpeg"
   "os"
   "image"
)

func loadJpeg(filename string) (image.Image, error) {
   f, err := os.Open(filename)
   if err != nil {
      return nil, err
   }
   defer f.Close()
   img, err := jpeg.Decode(f)
   if err != nil {
      return nil, err
   }
   return img, nil
}

func diff(pathA, pathB string) (float32, error) {
   i100, err := loadJpeg(pathA)
   if err != nil {
      return 0, err
   }
   i50, err := loadJpeg(pathB)
   if err != nil {
      return 0, err
   }
   if i50.ColorModel() != i100.ColorModel() {
      return 0, fmt.Errorf("different color models")
   }
   b := i50.Bounds()
   if !b.Eq(i100.Bounds()) {
      return 0, fmt.Errorf("different image sizes")
   }
   var sum uint32
   for y := b.Min.Y; y < b.Max.Y; y++ {
      for x := b.Min.X; x < b.Max.X; x++ {
         r1, g1, b1, _ := i50.At(x, y).RGBA()
         r2, g2, b2, _ := i100.At(x, y).RGBA()
         sum += absDiff(r1, r2)
         sum += absDiff(g1, g2)
         sum += absDiff(b1, b2)
      }
   }
   nPixels := (b.Max.X - b.Min.X) * (b.Max.Y - b.Min.Y)
   return float32(sum) / (float32(nPixels) * 0xFFFF *3), nil
}

func absDiff(x, y uint32) uint32 {
   if x < y {
      return y - x
   }
   return x - y
}

func main() {
   f, err := diff("Lenna100.jpg", "Lenna50.jpg")
   if err != nil {
      panic(err)
   }
   fmt.Println(f)
}
