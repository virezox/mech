package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
   "time"
)

var body = strings.NewReader(`{"supported_drm_types":["Dash","Hls","Mpeg"],"consumption_type":"Streaming","use_adaptive_bit_rate":true,"response_groups":"content_reference,chapter_info,pdf_url,last_position_heard,ad_insertion"}`)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "api.audible.com"
   req.URL.Path = "/1.0/content/B00551W570/licenserequest"
   req.URL.Scheme = "https"
   req.Header["X-Adp-Alg"] = []string{"SHA256WithRSA:1.0"}
   req.Header["X-Adp-Signature"] = []string{"MN9WDIkd3qb8E+funPEvsMzvTvZe5QUCC1uXX1VbPI24DY5QWnNzY8/fdTXowOz77MIt1URFCIFV9wNhEb5SbXpjff6h1morL35WhQa98Scu8pZTa7zDZifI/STXsgtH1BDL4KPxvlEY5QtKIeEr76GCjpdEQixYtZ1S14uXd4m3ZD+i2EF1ctN5tx950H0utIXEYYaQh8YdBgtb0YLcjjlfkhQMadGiPDAkeOzt//zBW++rC9Uey1p8TbaU+ndYrS6Hn2w+mqgzniMgffdqH/lQM6kX61AuDFmsH3gEvMwV9RqoMDRZ9zvhFJWi42TmdZkqjdev7GSkcX+M0ECdMA==:2022-02-16T22:48:21Z"}
   req.Header["X-Adp-Token"] = []string{"{enc:T1q+czVAQ/YHNSIzeOQ5fuTcZG0whnPrL1LOGou6lSh7EreQXVc+MaFuS0x/FHKYpQxOg7iPc0rZZusHLpxssOC+qrjbGd2LbnPGJ54e5TZFS2IzBvP8/DUH7nFDfCIfnmcSfnpIUgHHglW8vGWOvaZNakzvcgY7NPdHEn2l7KBlFmw/zB7zgm8flBRL+1qufER06I21F6ACNqNnHUVjjCLMEUs2RL/XRhkEH++58hZ1uJAqcqO4iuHmQej2T5DjpOE3ZAbLscDcLuB0l1r9XOq4dlJSIqnqDQVuU9mcY2YjFcjIpgNmj88T4VrhhNid7O5Q148eIM2Hzi9kQWwQjPrsue3i2vEPIDGNNzA9k+hhQc8xvzaBI0cse8KsNhXperFSJn7oyj24B+Zye9u5akPKnEo8GnqgW9hwNcXQKh1t/mMJdE9D4lgkgIA54xm8MgfIvUYwu8WjHZkb/555KSvTCCuL1tVj9I4UJF+ApYZZK0ulF6swGBiksAoFnL0W7d7vUXsWundbKDsN66yrmKEBb/y2dUDENvF4jJjg/CfS5FWWbA9Vnb8rakOLtMGH4oWW+11xupNlnUFxtFqUm2wSPQFqRgBBr8bEz7SOvj1PNo0vjoYv+bLLf+dCaAFqdAiqT45bNGnxISwwFyqI8E2CVJ1lysViC5ko5vq8EYHS4oobp0Vx3uc4aciJV9sbv+JaFK29z8+cunrea55wRqKub0FIqf6rxXyHYWTw5lfql9qZm40tifjvDhcTRlcBZgHi/ubyLJdxNMJFcbmj3/JQw5wGx6ibTwPGRDYauWH8wixOcaBEVEyNTVFqfn8flvQMfn7dlHKfwJJ9SCrflNS9zgZ70mACahvZwkMqciQh4Jno7SNq7sxbxJZ1G3DO2zytA4KH7BF9au1NaXcii9lGdJiGRgKURKzVfhyK5nmRjQpT8UkjZjuCu6CCp0IRGfrvknmd+mldNge/Kg/ZFOha9W7MqOejjl4EHyT6eMBiWrLNCvLtssbEJPiSL+BIiSnca/FRwSXXu/A+1GDaacxnO4L0ubBkAytTvejfz4s=}{key:JPUuMlehlhRVleS6cNidcxP5tMlGKSE1+N1FezIk7tyxoCvINR3c+ormG4UzUFetBeeaX33Lh7QISwfBNT4cmF6p2lh6AzYBP58EvCit2rxSn/XEf6WenDOcs9J1xFQWJKx2JsFRnm5r14WllhX473ArrCGnilE/iN4zSG/xKk/YFa8io1CvR1aLqJQaz/wJwdIKzM4g9Motdty98kOXtugIel5h4wIabGDtjU6GwGpDsB8CoDJEJmw0+aZEBjSVxTlgKfV6Q/1olzQgSnAde23cdT7OkPUwh9JNGBIYCk0XnAkArz3QsBex165AJLs9MLu6YeSHX1sT1rjhYzm18A==}{iv:gcCAyMzZ1Cn0RlPMPWy07Q==}{name:QURQVG9rZW5FbmNyeXB0aW9uS2V5}{serial:Mg==}"}
   time.Sleep(time.Second)
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
