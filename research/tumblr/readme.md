# Tumblr

- https://github.com/RipMeApp/ripme/issues/1950
- https://lyssafreyguy.tumblr.com/post/187741823636
- https://play.google.com/store/apps/details?id=com.tumblr

Install system cert. Here is video:

~~~
GET /tumblr_pxw935uiWw1r3c91u.mp4 HTTP/1.1
Icy-MetaData: 1
User-Agent: tumblr/23.2.0.00 (Linux;Android 7.0) ExoPlayerLib/2.15.1
Accept-Encoding: identity
Host: va.media.tumblr.com
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
GET https://api-http2.tumblr.com/v2/blog/lyssafreyguy/posts/187741823636/permalink?reblog_info=true&filter=clean HTTP/2.0
user-agent: Tumblr/Android/23.2.0.00
x-version: device/23.2.0.00/0/7.0/tumblr/
x-identifier: ef8f7b13-e09b-46f7-a84b-19205a9c7815
x-identifier-date: 1644630699
accept-language: en-US
pragma: no-cache
x-yuser-agent: YMobile/1.0 (com.tumblr/1; Android/7.0; NYC; generic_x86; Google; Android SDK built for x86; 4.99; 1080x1794;)
x-real-user-agent: YMobile/1.0 (com.tumblr/1; Android/7.0; NYC; generic_x86; Google; Android SDK built for x86; 4.99; 1080x1794;)
smart-user-agent: Mozilla/5.0 (Linux; Android 7.0; Android SDK built for x86) [Tumblr/1]
webview-user-agent: Mozilla/5.0 (Linux; Android 7.0; Android SDK built for x86 Build/NYC; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 Mobile Safari/537.36
di: DI/1.0 (310; 260; [Cellular])
x-background: false
x-fb-buyer-uid: eJx9VFGPmzgQ/isoT62UEAzGQPfJISRBSyACkt6qnCwDZouWhBUh23ar/veOYS+5vdMdDxbzzXhmPDPf/Jy4dLfzlsxfJpNPky9/TqaTJU0p28WR6yWJH65ZtEv9KJTq06VpwCBxY88L2Wd/mW4AxQhJcHnPDl6cgClgRLVV7Q2FCHThB376ICMYUzw1p9YUoSkiUwQ/tgxKVzvQ5o5jcKsyOMYFLzhxRK7nTkUMi2DNtKXHfeLFjK69MAX7bftaNw3P5qaqKR+C+nT5fqfQU9m1dalYqnYTIBElv9RNr1Rtp3y3ibIAqczm4YN7p3x7+ajQ5+dGfBb5fd2DP8NSDaJ8uN+k22CqNPWTUNaieGo/KgfRnev2lM0xxHS/du1RZHPiqJpqYMdSkaYp2zavG6EkvOJdfXX2ZbWgYTanl7IWp0KEov/Wdk+rtnvL8W61SKT+Jh2yuXwE3Ftk86I9qv3lmDedBEClG6oOUTVpsAAZ6YZmaJpGQD4kkJPsAfwHbjYXJ7ZPZJ1X+zD0AuZGsce8A1QxGTqooSlGmMChW3Bow58h//AUQ4PloctDtiBOV2y1YDA4zA+TlAaBtwQvFW/OAtT+kiXRPnY9wMKIrbcJgFt6L+V12z420uhAgx2N6VaG/5lN6jMTx0s2+ZRN+u4issk0m/DnJ3auX8WAGsR2HGI5g6avj+K1PQnWVtVZ9IPBTEdEk9/b3Wd27nnXi5J1gp+hX9IooPvQ3cAAraJ9uIQX+Doa7CF+KfLL4yPPmzHg8JpBNyjYC28uoybMMujSoJITVbL+x/N7zS850LvdYu8HsjDXxoyEAISPPb71I4jW63dFjBJJJgCsgUfgLaRbWcF0GIEbCzeev95ILhDbADT00s9RfM/Sh520xgAt4KmBFG4TBKhL49j3YoDpNRl/uaLykllouMTOrCSIzDC2qplTGs6MV0Xh5HmBhYX+e0/AbO3DNH647Yso+UcUWAnQBZYmbOj/THqj7uBnXBTMC+ni/VQlUhuFsKhkhpwbJtHIzOYOZFhU1YzbRTkrkVZZwsKmXZlj1d6qeOXK3zyl/lBQRDAmBgLeqsSRt/bhuKr+ihxELh3qN5AIgA1NmPdHtAvow1BAObFyxKOlF9xe+u+VIzNK4LJ0hTnSicULm9ga0VBBcI4wNgrbsi2kI2RwEwbauJP1euOYDCXZF0XpUBld+kvT2F/s02tl/qcv4CX1bl2hIYWpG+lXdUKw8zMvxjk2LYQc03D0YciLr7x7rE+Pg2qkS9/2vGFHcWy7HyNs2pZFTB3bI/1eeN1IKr2z0SzTNMBwdJvzvhdXFfRm5M3SC5Ox/rpKdHPy6zcgX9la
x-nimbus-session-id: 1e3bb6c2-089a-4e79-a451-70384e5ad202
x-s-id: NDVjMDRkNDktZDYxNi00NDdmLTlkMzktYWZjYzliYmM0ZTcx
x-s-id-enabled: true
x-performance-logging: true
authorization: OAuth oauth_consumer_key="BUHsuO5U9DF42uJtc8QTZlOmnUaJmBJGuU1efURxeklbdiLn9L", oauth_nonce="-6067022953793868178", oauth_signature="%2BrM9LNWUD4zlE00dcCI9EhCJ1kw%3D", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1644631110", oauth_token="Lb5eNiwqknunoZpJvtYmMxahPIyzMC1lgYlfjPm6nTpqZx977D", oauth_version="1.0"
accept-encoding: gzip
content-length: 0
~~~
