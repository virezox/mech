// OCR
package ocr

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "mime/multipart"
   "net/http"
   "os"
   "strings"
)

const (
   API = "http://api.ocr.space/parse/image"
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func createForm(form map[string]string) (string, io.Reader, error) {
   body := new(bytes.Buffer)
   mp := multipart.NewWriter(body)
   defer mp.Close()
   for key, val := range form {
      if strings.HasPrefix(val, "@") {
         val = val[1:]
         file, err := os.Open(val)
         if err != nil {
            return "", nil, err
         }
         defer file.Close()
         part, err := mp.CreateFormFile(key, val)
         if err != nil {
            return "", nil, err
         }
         io.Copy(part, file)
      } else {
         mp.WriteField(key, val)
      }
   }
   return mp.FormDataContentType(), body, nil
}

type Image struct {
   ParsedResults []struct {
      ParsedText string
   }
}

func NewImage(name string) (Image, error) {
   form := map[string]string{"OCREngine": "2", "file": "@" + name}
   ct, body, err := createForm(form)
   if err != nil {
      return Image{}, err
   }
   req, err := http.NewRequest("POST", API, body)
   if err != nil {
      return Image{}, err
   }
   req.Header.Set("Content-Type", ct)
   req.Header.Set("apikey", "helloworld")
   fmt.Println(invert, "POST", reset, form)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return Image{}, err
   }
   defer res.Body.Close()
   var img Image
   if err := json.NewDecoder(res.Body).Decode(&img); err != nil {
      return Image{}, err
   }
   return img, nil
}
