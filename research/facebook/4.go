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
   req.Header["Content-Length"] = []string{"1904"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{"datr=TELWYW7HP4RtCGikVySUp6x-", "fr=0eXRP555t4sF3Qm1C.AWXaLXrUjzhPadDjq-WuQdN7vsI.Bh2Mwn.wl.AAA.0.0.BiWJGF.AWXlEfWdOcM", "sb=y2kMYljEFad1tU7CPu3dtnFa", "dpr=1.25", "wd=1186x615", "locale=en_US", "m_pixel_ratio=1"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"www.facebook.com"}
   req.Header["Origin"] = []string{"https://www.facebook.com"}
   req.Header["Referer"] = []string{"https://www.facebook.com/CBSMornings/videos/shaquille-o-neal-says-he-didnt-spend-his-first-nba-check-until-he-started-having/514785953701098/"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   req.Header["X-Fb-Friendly-Name"] = []string{"CometUFICommentsProviderQuery"}
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

var body = strings.NewReader(`av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q322aew9G2S0zU20xi3y4o0B-q1ew65xOfw9q0yE465o-cwfG12wOKdwGwQw9m8wsU9k2C1FwIw9i1uwZwlo5qfK6E7e58jwGzE2swwwJK2W2K0zK5o4q3q1OBx_y8179obodEGdwda3e0Lo4qifxe&__csr=gZ4PnMxeAXERZRESZCG8ZaiFlF9yCFfx6t3eiXGFu2qhypVlyEoCgWAeACAUihpoGUgxamh1emt5AGcUaUyqm78O8yufxmdKu48DwwAKVefBwODDz8GE0n0Iw0UObw1lK14z80kf802Dy022-1cw08MO2C580-S0iS9Yw9EozV6nAzo0X20_E1Mk0kivIo0cCK0CUgw869zU3Ozo2lU0Fy09wgjwaa17w3z40ny9w3Do046903oQ0OU0v5wd26o1r40WpUqylwzw1m60aTg23wdG&__req=n&__hs=19097.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005359223&__s=ofp7hk%3Aob6ftr%3Arciyi6&__hsi=7086645409667728629-0&__comet_req=1&lsd=AVrNdoI-fb4&jazoest=2924&__spin_r=1005359223&__spin_b=trunk&__spin_t=1649988212&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometUFICommentsProviderQuery&variables=%7B%22UFI2CommentsProvider_commentsKey%22%3A%22CometVideoHomeNewPermalinkHeroUnitQuery%22%2C%22__false%22%3Afalse%2C%22__true%22%3Atrue%2C%22after%22%3A%22AQHRBlF4EPW5e3JaUBUGHPeEu9777PavHLB2MGd3H6LF8caqsGXo9ppggvdFNzDKNv2IalitA0XMtmpYBnOo5Dx6NQ%22%2C%22before%22%3Anull%2C%22displayCommentsContextEnableComment%22%3Anull%2C%22displayCommentsContextIsAdPreview%22%3Anull%2C%22displayCommentsContextIsAggregatedShare%22%3Anull%2C%22displayCommentsContextIsStorySet%22%3Anull%2C%22displayCommentsFeedbackContext%22%3Anull%2C%22feedLocation%22%3A%22TAHOE%22%2C%22feedbackSource%22%3A41%2C%22first%22%3A50%2C%22focusCommentID%22%3Anull%2C%22includeHighlightedComments%22%3Afalse%2C%22includeNestedComments%22%3Atrue%2C%22initialViewOption%22%3A%22RANKED_THREADED%22%2C%22isInitialFetch%22%3Afalse%2C%22isPaginating%22%3Atrue%2C%22last%22%3Anull%2C%22scale%22%3A1.5%2C%22topLevelViewOption%22%3Anull%2C%22useDefaultActor%22%3Afalse%2C%22viewOption%22%3Anull%2C%22id%22%3A%22ZmVlZGJhY2s6NTA4Nzg1NDIxNzkzNDg3Ng%3D%3D%22%7D&server_timestamps=true&doc_id=7280701748638794`)
