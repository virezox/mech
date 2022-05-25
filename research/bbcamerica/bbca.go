package main

import (
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

var body = strings.NewReader(`
{
  "adobeShortMediaToken": "",
  "hba": false,
  "adtags": {
    "lat": 0,
    "url": "https://www.bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529",
    "playerWidth": 1920,
    "playerHeight": 1080,
    "ppid": 1,
    "mode": "on-demand"
  },
  "useLowResVideo": false
}
`)

func main() {
   req, err := http.NewRequest(
      "POST",
      "https://gw.cds.amcn.com/playback-id/api/v1/playback/1052529",
      body,
   )
   if err != nil {
      panic(err)
   }
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Authorization"] = []string{"Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xMTEtNjcuZWMyLmludGVybmFsIiwidG9rZW5fdHlwZSI6ImF1dGgiLCJleHAiOjE2NTM1MDczMDcsImRldmljZS1pZCI6IjAzMWE3ZWQ0LTQ0NTUtNDZkNS05YjJhLTBmMTVkODRkNzVlYyIsImlhdCI6MTY1MzUwNjcwNywianRpIjoiNzU0NDU1MmEtNzcwYy00OGM1LWI1ZTgtMzY0MzkwMjgyOGExIn0.otEDejVgDHnkKuo-Ya5hm5b46ZENk1BC0S7964JV7fG9d-NB1Pnu_k6eQyLxmZ5BCErlcPIABbG6couXZ1C4cxRjn0R9N5XBRCs585SNo2C7XrjkN3ScxnTmv_5axocapKkSfm3QkDKv9BRHhUBuLeE7HTC61WuN4DZWFwVYJ_ro2b_o1cKtceXNo7PaP_krgBjq61c0InqB5Vxr4fnIQ_L3-yOLgLbkXlI7ficsmTrrAaKHEFSSK6HmiVEoF3qpM2ciZ76i4PkSBCg5n73TjbahybAPNstbRMMnVk8lEUlTeR3t92KbIk5iWWArDJ8YODOn6hiPxIFy8cd3Rm1REw"}
   req.Header["Content-Length"] = []string{"241"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Origin"] = []string{"https://www.bbcamerica.com"}
   req.Header["Referer"] = []string{"https://www.bbcamerica.com/"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   req.Header["X-Amcn-Adobe-Id"] = []string{""}
   req.Header["X-Amcn-Audience-Id"] = []string{""}
   req.Header["X-Amcn-Cache-Hash"] = []string{"8679a33532b1f9b3310c9af5f95e855ed49159900277aeb54aa8da1e5a5c445e"}
   req.Header["X-Amcn-Device-Ad-Id"] = []string{"f4b9e72e-2ec3-47a1-a8a6-5f2115ee6da3"}
   req.Header["X-Amcn-Device-Id"] = []string{"f4b9e72e-2ec3-47a1-a8a6-5f2115ee6da3"}
   req.Header["X-Amcn-Language"] = []string{"en"}
   req.Header["X-Amcn-Mvpd"] = []string{""}
   req.Header["X-Amcn-Network"] = []string{"bbca"}
   req.Header["X-Amcn-Platform"] = []string{"web"}
   req.Header["X-Amcn-Service-Group-Id"] = []string{"6"}
   req.Header["X-Amcn-Service-Id"] = []string{"bbca"}
   req.Header["X-Amcn-Tenant"] = []string{"amcn"}
   req.Header["X-Ccpa-Do-Not-Sell"] = []string{"passData"}
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
