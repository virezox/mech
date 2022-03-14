package vimeo

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

type Check struct {
   Request struct {
      Files struct {
         Progressive []Progressive
      }
   }
}

func (c Clip) Check(password string) (*Check, error) {
   body := new(bytes.Buffer)
   body.WriteString("password=")
   io.WriteString(base64.NewEncoder(base64.StdEncoding, body), password)
   req, err := http.NewRequest(
      "POST", 
      fmt.Sprintf("https://player.vimeo.com/video/%v/check-password", c.ID),
      body,
   )
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
   Width int
   Height int
   FPS int
   URL string
}

func (p Progressive) Format(f fmt.State, verb rune) {
   fmt.Fprint(f, "Width:", p.Width)
   fmt.Fprint(f, " Height:", p.Height)
   fmt.Fprint(f, " FPS:", p.FPS)
   if verb == 'a' {
      fmt.Fprint(f, " URL:", p.URL)
   }
}
