package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"1053"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{"datr=TELWYW7HP4RtCGikVySUp6x-", "dpr=1.25", "fr=0eXRP555t4sF3Qm1C.AWXaLXrUjzhPadDjq-WuQdN7vsI.Bh2Mwn.wl.AAA.0.0.BiWJGF.AWXlEfWdOcM", "locale=en_US", "m_pixel_ratio=1", "sb=y2kMYljEFad1tU7CPu3dtnFa", "wd=1186x615"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"www.facebook.com"}
   req.Header["Origin"] = []string{"https://www.facebook.com"}
   req.Header["Referer"] = []string{"https://www.facebook.com/TasteLifeOfficial/videos/how-a-professional-makes-beef-bacon-from-start-to-finish/2883317948625723/"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   req.Header["X-Fb-Friendly-Name"] = []string{"CometHovercardQueryRendererQuery"}
   req.Header["X-Fb-Lsd"] = []string{"AVpJcKogJYw"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.facebook.com"
   req.URL.Path = "/api/graphql/"
   req.URL.RawPath = ""
   val := make(url.Values)
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
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

var body = strings.NewReader(`av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cwfG12wOKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q0GpovUy0hOm2S3qazo3iwPwbS16AzUjw&__csr=gRb2a5Oh9kNt2cyl49F5ybydpWBBVXJ6KeCJbgGFk7eegGcQumibBCUK9x69yGwwKV8hVHgOcy5yRx-byGzUgU9HAyo98nF6BADwFmumu68uByUx00X5wiQ04MU989E-01aVwiU02btwFwk80ZW0h1iG3K4bpuUW04H81H8jw52dCgCyw378K0C40AoC2a0VpE2kG04-e4o3Kw3zoCm05ao047fw2To8U09xU1-o4K2q264UCm6Xw1mu0cXwbe&__req=n&__hs=19096.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005355623&__s=rp1owa%3A1i38ia%3Aleoeya&__hsi=7086586611304571549-0&__comet_req=1&lsd=AVpJcKogJYw&jazoest=21007&__spin_r=1005355623&__spin_b=trunk&__spin_t=1649974522&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometHovercardQueryRendererQuery&variables=%7B%22actionBarRenderLocation%22%3A%22WWW_COMET_HOVERCARD%22%2C%22context%22%3A%22DEFAULT%22%2C%22entityID%22%3A%22252891408454022%22%2C%22includeTdaInfo%22%3Afalse%2C%22scale%22%3A1.5%7D&server_timestamps=true&doc_id=5724338484260388`)
