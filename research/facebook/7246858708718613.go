package main

import (
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

// 273086
func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"1622"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{"datr=TELWYW7HP4RtCGikVySUp6x-", "dpr=1.25", "fr=0eXRP555t4sF3Qm1C.AWXaLXrUjzhPadDjq-WuQdN7vsI.Bh2Mwn.wl.AAA.0.0.BiWJGF.AWXlEfWdOcM", "locale=en_US", "m_pixel_ratio=1", "sb=y2kMYljEFad1tU7CPu3dtnFa", "wd=1186x615"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"www.facebook.com"}
   req.Header["Origin"] = []string{"https://www.facebook.com"}
   req.Header["Referer"] = []string{"https://www.facebook.com/watch"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"}
   req.Header["X-Fb-Friendly-Name"] = []string{"CometVideoHomeNewPermalinkHeroUnitQuery"}
   req.Header["X-Fb-Lsd"] = []string{"AVpJcKogJYw"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.facebook.com"
   req.URL.Path = "/api/graphql/"
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
   println(len(buf))
   return
   os.Stdout.Write(buf)
}

var body = strings.NewReader(`av=0&__user=0&__a=1&__dyn=7xeUmwlEnwn8K2WnFw9-2i5U4e0yoW3q32360CEbo19oe8hw2nVE4W0om782Cw8G11xmfz81sbzo5-0Boy1PwBgao6C0Mo5W3S1lwlE-UqwsUkxe2Gew9O222SUbEaU2eU5O0GpovU19pobodEGdw46wbS16Awzw&__csr=gLkIrqkh9PCIhF2fJkjmAm8K8Rxyut2EWaiQay47efxvhVp8Km8yUryoS2i78Kcz8xoJovyUGE-48coC2i5XCBz8alDzooxW4oC03N604N8989E-01aVw0a9hiG3K4boKew1aO06g82Ag2hyo0qwU0iuw08yC0ayyo&__req=a&__hs=19096.HYP%3Acomet_loggedout_pkg.2.0.0.0.&dpr=1.5&__ccg=EXCELLENT&__rev=1005355623&__s=rp1owa%3A1i38ia%3Aleoeya&__hsi=7086586611304571549-0&__comet_req=1&lsd=AVpJcKogJYw&jazoest=21007&__spin_r=1005355623&__spin_b=trunk&__spin_t=1649974522&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=CometVideoHomeNewPermalinkHeroUnitQuery&variables=%7B%22displayCommentsContextEnableComment%22%3Anull%2C%22displayCommentsContextIsAdPreview%22%3Anull%2C%22displayCommentsContextIsAggregatedShare%22%3Anull%2C%22displayCommentsContextIsStorySet%22%3Anull%2C%22displayCommentsFeedbackContext%22%3Anull%2C%22focusCommentID%22%3Anull%2C%22privacySelectorRenderLocation%22%3A%22COMET_STREAM%22%2C%22renderLocation%22%3A%22video_home%22%2C%22scale%22%3A1.5%2C%22useDefaultActor%22%3Afalse%2C%22videoID%22%3A%222883317948625723%22%2C%22UFI2CommentsProvider_commentsKey%22%3A%22CometVideoHomeNewPermalinkHeroUnitQuery%22%2C%22caller%22%3A%22TAHOE%22%2C%22channelEntryPoint%22%3A%22TAHOE%22%2C%22channelID%22%3A%22%22%2C%22feedbackSource%22%3A41%2C%22feedLocation%22%3A%22TAHOE%22%2C%22isLoggedOut%22%3Atrue%2C%22skipFetchingChaining%22%3Atrue%2C%22streamChainingSection%22%3Afalse%2C%22videoChainingContext%22%3Anull%7D&server_timestamps=true&doc_id=7246858708718613`)
var LogLevel format.LogLevel

func video(v int64) (*http.Response, error) {
   vars, err := json.Marshal(map[string]int64{"videoID": v})
   if err != nil {
      return nil, err
   }
   body := url.Values{
      "doc_id": {"7246858708718613"},
      "variables": {string(vars)},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://www.facebook.com/api/graphql/", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
