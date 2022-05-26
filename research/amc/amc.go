package main

import (
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

var body = strings.NewReader(`
{
   "adtags": {
      "lat": 0,
      "mode": "on-demand",
      "ppid": 1,
      "playerWidth": 1920,
      "playerHeight": 1080,
      "url": "https://www.amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152"
   }
}
`)

func main() {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/playback-id/api/v1/playback/1011152", body,
   )
   if err != nil {
      panic(err)
   }
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["X-Amcn-Device-Ad-Id"] = []string{"3c96846a-add8-4af5-a65d-5ab42d5a4612"}
   req.Header["X-Amcn-Language"] = []string{"en"}
   req.Header["X-Amcn-Network"] = []string{"amcplus"}
   req.Header["X-Amcn-Platform"] = []string{"web"}
   req.Header["X-Amcn-Service-Id"] = []string{"amcplus"}
   req.Header["X-Amcn-Tenant"] = []string{"amcn"}
   req.Header["X-Ccpa-Do-Not-Sell"] = []string{"doNotPassData"}
   req.Header["Authorization"] = []string{"Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJhbWNuLWF1dGgsb2Itc3ViLWFtY3BsdXMiLCJkZWZhdWx0X3Byb2ZpbGVzIjpbeyJwcm9maWxlSWQiOjI3Nzg1NjksInByb2ZpbGVOYW1lIjoiRGVmYXVsdCIsInNlcnZpY2VJZCI6ImFtY3BsdXMifV0sImF1dGhfdHlwZSI6ImJlYXJlciIsImFtY24tYWNjb3VudC1jb3VudHJ5IjoidXMiLCJyb2xlcyI6WyJhbWNuLWF1dGgiLCJvYi1zdWItYW1jcGx1cyJdLCJpc3MiOiJpcC0xMC0yLTExMy01Ni5lYzIuaW50ZXJuYWwiLCJ0b2tlbl90eXBlIjoiYXV0aCIsImF1ZCI6InJlc291cmNlX3NlcnZlciIsImFtY24tYWNjb3VudC1pZCI6ImU1OWU2Mjk2LWJkYzQtNGIzMi1hOGQzLWIxMjIyMjkwNTQzNCIsImV4cCI6MTY1MzU5Mjg2MCwiaWF0IjoxNjUzNTkyMjYwLCJhbWNuLXNlcnZpY2UtZ3JvdXAtaWQiOiIxMCIsImp0aSI6IjQyMDY1NWY3LTYwNmMtNDExNS05MjBhLWQ0OGIyZTNkZWRjYiIsImFtY24tdXNlLWFjY291bnQtY291bnRyeSI6dHJ1ZX0.IA2L_9UQNJxWcBmNNtEy_1ywRNctbqpX2JYE_4jf94wWqS6lbGBbdGyRgqt9-3W1VZ65_DRRmbztglYU59dvA-WCtGMourzMOabLWCydxPx9e2z411tyiEnLCUo5pXKcK3VI51FQF9abq9gckxKxUSpBvSgmNHZ7uFUIJZVIEjTkXilRjwAujFQHomjcTGfCLMnrvl0BbKtzj0siW9Q9T8PgYT_TVztLDSzmac7kq1qkLVnUm060migJ-W-1rq45VrwmnR0lhAvpQT4qKUb2W9GFywaq_5058vIbsPQonzvJVCekJXrP1xOVcD6vuSa0vM-iPTSNfX_dU51Yg_uDjQ"}
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
