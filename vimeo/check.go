package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

func (c Clip) Check(password string) (*Check, error) {
   addr := []byte("https://player.vimeo.com/video/")
   addr = strconv.AppendInt(addr, c.ID, 10)
   addr = append(addr, "/check-password"...)
   body := new(bytes.Buffer)
   body.WriteString("password=")
   io.WriteString(base64.NewEncoder(base64.StdEncoding, body), password)
   req, err := http.NewRequest("POST", string(addr), body)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
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

func (p Progressive) WithURL(s string) Progressive {
   p.URL = s
   return p
}

func (p Progressive) String() string {
   var buf []byte
   buf = append(buf, "Width:"...)
   buf = strconv.AppendInt(buf, p.Width, 10)
   buf = append(buf, " Height:"...)
   buf = strconv.AppendInt(buf, p.Height, 10)
   buf = append(buf, " FPS:"...)
   buf = strconv.AppendInt(buf, p.FPS, 10)
   if p.URL != "" {
      buf = append(buf, " URL:"...)
      buf = append(buf, p.URL...)
   }
   return string(buf)
}

type Check struct {
   Request struct {
      Files struct {
         Progressive []Progressive
      }
   }
}
