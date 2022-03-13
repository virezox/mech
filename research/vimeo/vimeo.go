package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

type Progressive struct {
   Width int64
   Height int64
   URL string
}

func (p Progressive) String() string {
   var buf []byte
   buf = append(buf, "Width:"...)
   buf = strconv.AppendInt(buf, p.Width, 10)
   buf = append(buf, " Height:"...)
   buf = strconv.AppendInt(buf, p.Height, 10)
   if p.URL != "" {
      buf = append(buf, " URL:"...)
      buf = append(buf, p.URL...)
   }
   return string(buf)
}

type Video struct {
   Request struct {
      Files struct {
         Progressive Progressive
      }
   }
}

func NewVideo(id int64, password string) (*Video, error) {
   addr := []byte("https://player.vimeo.com/video/")
   addr = strconv.AppendInt(addr, id, 10)
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
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}
