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
   req.Header["Content-Length"] = []string{"977"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{"datr=TELWYW7HP4RtCGikVySUp6x-", "fr=0eXRP555t4sF3Qm1C.AWXaLXrUjzhPadDjq-WuQdN7vsI.Bh2Mwn.wl.AAA.0.0.BiWJGF.AWXlEfWdOcM", "sb=y2kMYljEFad1tU7CPu3dtnFa", "dpr=1.25", "wd=1186x615", "locale=en_US", "m_pixel_ratio=1"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"www.facebook.com"}
   req.Header["Origin"] = []string{"https://www.facebook.com"}
   req.Header["Referer"] = []string{"https://www.facebook.com/CBSMornings/videos/shaquille-o-neal-says-he-didnt-spend-his-first-nba-check-until-he-started-having/514785953701098/"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   req.Header["X-Fb-Friendly-Name"] = []string{"CometTahoeUpNextEndCardQuery"}
   req.Header["X-Fb-Lsd"] = []string{"AVrNdoI-fb4"}
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

var body = strings.NewReader(`av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cwfG12wOKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q3q1OBx_y8179obodEGdwda3e0Lo4qifxe&__csr=gZ4PnONkWjKGtZRES_CG8Zai8mACbiA-6kcVbKGBU9F69DBmaxyp3GgWiqjx95ByHx24Fp44VpQmiEPwHy9Fosz8y9U-5oSVUgyu22iXAU-m3auucyE0n3Iw0UObw1lK14z80kf801aEwj802ccwFxi0fJw4Jyv82q68-hBV8S0eMwfW0s5054DX6039Hw9K4821yo-0YES0Bu0aow2o44U2ywhU0UN05Uyo0VS011yg0Sd0cK07No3gxC0mN0eCu6EBo8U0lxw3f83qw&__req=g&__hs=19097.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005359223&__s=ofp7hk%3Aob6ftr%3Arciyi6&__hsi=7086645409667728629-0&__comet_req=1&lsd=AVrNdoI-fb4&jazoest=2924&__spin_r=1005359223&__spin_b=trunk&__spin_t=1649988212&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometTahoeUpNextEndCardQuery&variables=%7B%22upNextVideoID%22%3A%22%22%2C%22scale%22%3A1.5%2C%22currentID%22%3A%22514785953701098%22%7D&server_timestamps=true&doc_id=5242228839129775`)
