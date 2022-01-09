package main
import (
"net/http"
"net/http/httputil"
"net/url"
"os"
"io"
"strings"
)

var body = strings.NewReader(`
{
   "deviceUuid": 1880,
   "includeItem": true,
   "sourceId": "TR:1168891"
}
`)

var req = &http.Request{
ContentLength: int64(body.Len()),
Body:io.NopCloser(body),
Header:http.Header{
   "Content-Type":[]string{"application/json"},
   "X-Authtoken":[]string{"BIQzNzpF+y/+qr67/VyiRuY05aHotwNVIjGdDIGEtQog1UNfywnxI4cA=="},
},
Method:"POST",
URL:&url.URL{
Host:"pandora.com",
Path:"/api/v1/playback/source",
Scheme:"https"}}
func main() {
res, err := new(http.Transport).RoundTrip(req)
if err != nil {
panic(err)}
defer res.Body.Close()
buf, err := httputil.DumpResponse(res, true)
if err != nil {
panic(err)}
os.Stdout.Write(buf)}
