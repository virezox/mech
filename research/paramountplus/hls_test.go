package hls

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "testing"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36"

func TestDecrypt(t *testing.T) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "cbsios-vh.akamaihd.net"
   req.URL.Path = "/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_,2228,4628,3128,1628,848,503,000.mp4.csmil/index_1_av.m3u8"
   val := make(url.Values)
   val["hdntl"] = []string{"exp=1643411403~acl=/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_*~data=hdntl~hmac=7261be21205f1ead3f6936d7b91c014ed22fbba849dbafc068f303ebfc8864cc"}
   val["id"] = []string{"AgBItRcmFy81Sksm82G0S6Vus5DhGvuvBZwDsGQTvpPJN dt XkZKPiuTw6mxQdAIFPdZjWHxM4qug=="}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Header["User-Agent"] = []string{userAgent}
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   index, err := io.ReadAll(res.Body)
   if err != nil {
      t.Fatal(err)
   }
   var block *Block
   for _, form := range Unmarshal(index) {
      req, err := http.NewRequest("GET", form["URI"], nil)
      if err != nil {
         t.Fatal(err)
      }
      req.Header.Set("User-Agent", userAgent)
      if form["METHOD"] != "" {
         block, err = newCipher(req)
         if err != nil {
            t.Fatal(err)
         }
      } else if block != nil {
         dst, err := block.decrypt(req)
         if err != nil {
            t.Fatal(err)
         }
         if err := os.WriteFile("segment.ts", dst, os.ModePerm); err != nil {
            t.Fatal(err)
         }
         break
      }
   }
}
