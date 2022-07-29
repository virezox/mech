package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

func (p Progressive) String() string {
   b := []byte("Width:")
   b = strconv.AppendInt(b, p.Width, 10)
   b = append(b, " Height:"...)
   b = strconv.AppendInt(b, p.Height, 10)
   b = append(b, " FPS:"...)
   b = strconv.AppendInt(b, p.FPS, 10)
   return string(b)
}

type Check struct {
   Request struct {
      Files struct {
         Progressive []Progressive
      }
   }
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

type Progressive struct {
   Width int64
   Height int64
   FPS int64
   URL string
}

func (p Progressive) Height_Distance(v int64) int64 {
   if p.Height > v {
      return p.Height - v
   }
   return v - p.Height
}
