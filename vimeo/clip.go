package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Clip struct {
   ID int64
   Unlisted_Hash string
}

func New_Clip(reference string) (*Clip, error) {
   ref, err := url.Parse(reference)
   if err != nil {
      return nil, err
   }
   fields := strings.FieldsFunc(ref.Path, func(r rune) bool {
      return r == '/'
   })
   var clip Clip
   for _, field := range fields {
      if clip.ID >= 1 {
         clip.Unlisted_Hash = field
      } else if field != "video" {
         clip.ID, err = strconv.ParseInt(field, 10, 64)
         if err != nil {
            return nil, err
         }
      }
   }
   for _, key := range []string{"h", "unlisted_hash"} {
      hash := ref.Query().Get(key)
      if hash != "" {
         clip.Unlisted_Hash = hash
      }
   }
   return &clip, nil
}

func (c Clip) Check(password string) (*Check, error) {
   // URL
   ref := []byte("https://player.vimeo.com/video/")
   ref = strconv.AppendInt(ref, c.ID, 10)
   ref = append(ref, "/check-password"...)
   // body
   body := new(bytes.Buffer)
   body.WriteString("password=")
   io.WriteString(base64.NewEncoder(base64.StdEncoding, body), password)
   req, err := http.NewRequest("POST", string(ref), body)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   check := new(Check)
   if err := json.NewDecoder(res.Body).Decode(check); err != nil {
      return nil, err
   }
   return check, nil
}
