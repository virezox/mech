Perceptual hash distance between valid and invalid images

I found this article [1] very helpful, but I wanted to do some of my own testing. I wanted to take a "source" image, and compare it against something pretty close, and against something pretty different. To that end, here are the three images:

source:
https://ia800309.us.archive.org/9/items/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg

good:
http://i.ytimg.com/vi/11Bvzknjo2Q/hqdefault.jpg

bad:
http://i.ytimg.com/vi/jCMi9_6vnxk/hqdefault.jpg

The last two images have padding, so I removed that before doing any hashing. My program is below. My thinking is that, the distance value should be low for the "good" image, and high for the "bad" image. I tried with Average Hash, Difference Hash and Perceptual Hash. Results:

hash       | good distance | bad distance | bad - good
-----------|---------------|--------------|-----------
Average    | 1             | 26           | 25
Difference | 7             | 35           | 28
Perceptual | 10            | 22           | 12

So what concerns me, is that with Perceptual, you certainly have a difference between "good" and "bad", but its way lower than the other options. I wanted to present this finding, as I saw from the article you ended up choosing the Perceptual option.

~~~go
package main

import (
   "github.com/corona10/goimagehash"
   "image"
   "image/jpeg"
   "net/http"
)

const (
   urlA =
      "https://ia800309.us.archive.org/9/items" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1" +
      "/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg"
   urlB = "http://i.ytimg.com/vi/11Bvzknjo2Q/hqdefault.jpg"
   urlC = "http://i.ytimg.com/vi/jCMi9_6vnxk/hqdefault.jpg"
)

func hash(addr string, crop bool) (*goimagehash.ImageHash, error) {
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return nil, err
   }
   if crop {
      height := 270
      x0 := (480 - height) / 2
      y0 := (360 - height) / 2
      r := image.Rect(x0, y0, x0 + height, y0 + height)
      i = i.(*image.YCbCr).SubImage(r)
   }
   return goimagehash.PerceptionHash(i)
}

func main() {
   a, err := hash(urlA, false)
   if err != nil {
      panic(err)
   }
   b, err := hash(urlB, true)
   if err != nil {
      panic(err)
   }
   c, err := hash(urlC, true)
   if err != nil {
      panic(err)
   }
   ab, err := a.Distance(b)
   if err != nil {
      panic(err)
   }
   ac, err := a.Distance(c)
   if err != nil {
      panic(err)
   }
   println(ab, ac)
}
~~~

1. https://content-blockchain.org/research/testing-different-image-hash-functions
