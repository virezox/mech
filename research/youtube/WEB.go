package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
   "strings"
)

func WEB() {
   res, err := http.Get("https://www.youtube.com")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}

type token struct {
   Access_Token string
}

func WEB_CREATOR() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   tok, err := format.Open[token](cache, "mech/youtube.json")
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest("GET", "https://studio.youtube.com", nil)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = "approve_browser_access=true"
   req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}

func WEB_EMBEDDED_PLAYER() {
   res, err := http.Get("https://www.youtube.com/embed/MIchMEqVwvg")
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}

func WEB_KIDS() {
   req, err := http.NewRequest("GET", "https://www.youtubekids.com", nil)
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "Firefox/44")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}

func WEB_REMIX() {
   req, err := http.NewRequest("GET", "https://music.youtube.com", nil)
   if err != nil {
      panic(err)
   }
   req.Header.Set("User-Agent", "Firefox/44")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   dst := new(strings.Builder)
   io.Copy(dst, res.Body)
   low := strings.Index(dst.String(), `"clientName"`)
   fmt.Println(dst.String()[low:low+99])
}
