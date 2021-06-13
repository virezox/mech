package main

import (
   "encoding/json"
   "flag"
   "fmt"
   "github.com/89z/mech/deezer"
   "io"
   "net/http"
   "os"
   "path/filepath"
   "strings"
)

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func getArl(har string) (string, error) {
   file, err := os.Open(har)
   if err != nil { return "", err }
   defer file.Close()
   var archive struct {
      Log struct {
         Entries []struct {
            Request struct {
               Cookies []struct { Name, Value string }
            }
         }
      }
   }
   if err := json.NewDecoder(file).Decode(&archive); err != nil {
      return "", err
   }
   for _, entry := range archive.Log.Entries {
      for _, cookie := range entry.Request.Cookies {
         if cookie.Name == "arl" { return cookie.Value, nil }
      }
   }
   return "", fmt.Errorf("Arl cookie not found")
}

func writeFile(sngId, arl, format string) error {
   track, err := deezer.NewTrack(sngId, arl)
   if err != nil { return err }
   var audio rune
   if format == "flac" {
      audio = deezer.FLAC
   } else {
      audio = deezer.MP3_320
   }
   source, err := track.Source(sngId, audio)
   if err != nil { return err }
   fmt.Println(invert, "GET", reset, source)
   res, err := http.Get(source)
   if err != nil { return err }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil { return err }
   if err := deezer.Decrypt(sngId, body); err != nil {
      return err
   }
   name := fmt.Sprintf("%v - %v.%v", track.ART_NAME, track.SNG_TITLE, format)
   return os.WriteFile(strings.Map(clean, name), body, os.ModePerm)
}

func clean(r rune) rune {
   if strings.ContainsRune(`"*/:<>?\|`, r) {
      return -1
   }
   return r
}

func main() {
   var format string
   flag.StringVar(&format, "f", "mp3", "format")
   har, err := os.UserConfigDir()
   if err != nil {
      panic(err)
   }
   har = filepath.Join(har, "deezer.har")
   flag.StringVar(&har, "h", har, "HTTP archive")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("deezer [flags] SNG_ID")
      flag.PrintDefaults()
      return
   }
   arl, err := getArl(har)
   if err != nil {
      panic(err)
   }
   if err := writeFile(flag.Arg(0), arl, format); err != nil {
      panic(err)
   }
}
