package tiktok

import (
   "encoding/hex"
   "net/http"
   "net/http/httputil"
   "os"
   "strconv"
   "testing"
   "time"
)

func TestRegister(t *testing.T) {
   res, err := register()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(append(buf, '\n'))
}

func TestTikTok(t *testing.T) {
   req, err := http.NewRequest(
      "GET", "https://api-h2.tiktokv.com/aweme/v1/aweme/detail/", nil,
   )
   if err != nil {
      t.Fatal(err)
   }
   deviceParams.Set("aweme_id", "7038818332270808325")
   req.URL.RawQuery = deviceParams.Encode()
   gorgon, err := genXGorgon(req.URL.RawQuery)
   if err != nil {
      t.Fatal(err)
   }
   req.Header = http.Header{
      "user-agent": {"okhttp/3.10.0.1"},
      "x-gorgon": {hex.EncodeToString(gorgon)},
      "x-khronos": {strconv.FormatInt(time.Now().Unix(), 10)},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
