package widevine

import (
   "os"
   "testing"
)

var peacock = request{
  PSSH: "AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAWylyu1Ra6Obew32huzP9I49yVmwY=",
  License: "https://ovp.peacocktv.com/drm/widevine/acquirelicense?bt=43-0Bzfu3_lu2EG4AlSU4eQZPavwvjM-lP5h7cwPYI5YHEHWYzH8lqxPDqB6SJMtSwbi0_HDKaAsOjzDBwMLIvsYBKmwLhrhdpcQ_MJZXpSIR5e52sohHDwx9VyX3rtJlv8X3vOB6Fkn77yMYTzF5R2YnvXanei895p9hf9nb8hrBKY7DWMNSx03Qy7NqAKrgwZzRc00_RoolxslOVKZ2yWhvPUhCOECwnxEwHa07zNGfOBT6znd6v_gjyYs2s3YdTjK8URKKMHl8P7esyxB5Bwl6ln0svU55jTYs4V81FxUbMfjRjn49isEWBJaCkwnd4sorzkazTiXAN2g4HxexXuwISsZ3CWbOJM5MzFHGdZ8jh7ox8ZOCTvS0VPPQ==",
  Headers: "x-sky-signature: SkyOTT client=\"NBCU-WEB-v6\",signature=\"7xZxtIkRNy5a0ZpVQUVcEfQtHPQ=\",timestamp=\"1650813143\",version=\"1.0\"",
  Cache: false,
}

func TestWidevine(t *testing.T) {
   res, err := peacock.post()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
