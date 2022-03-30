# BBC America

- https://bbcamerica.com/shows/doctor-who/episodes/season-13-eve-of-the-daleks--1051589
- https://github.com/ytdl-org/youtube-dl/issues/30182
- https://play.google.com/store/apps/details?id=com.bbca.bbcafullepisodes

Install system certificate. This fails:

~~~
GET https://gw.cds.amcn.com/content-compiler-cr/api/v1/content/amcn/bbca//url/shows/doctor-who/episodes/season-13-eve-of-the-daleks--1051589?device=mobile HTTP/2.0
tracestate: @nr=0-2-2330446-389368516-7c3892ee12774a66--00--1644778911883
traceparent: 00-37d36f48ebda4577a3cedac0930b42ed-7c3892ee12774a66-00
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjIzMzA0NDYiLCJkLmFwIjoiMzg5MzY4NTE2IiwiZC50ciI6IjM3ZDM2ZjQ4ZWJkYTQ1NzdhM2NlZGFjMDkzMGI0MmVkIiwiZC5pZCI6IjNiOTAxMzY5MDI2NzRiMGQiLCJkLnRpIjoxNjQ0Nzc4OTExODgzfX0=
x-amcn-device-id: 893968a6-8fee-4201-9964-f1e2919af107
x-amcn-language: en
x-amcn-network: bbca
x-amcn-platform: android
x-amcn-tenant: amcn
x-amcn-audience-id: amcn
x-amcn-device-ad-id: f187f55d-3f78-4267-8830-719a99cb7413
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-ccpa-do-not-sell: default
authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xMjEtMTg0LmVjMi5pbnRlcm5hbCIsInRva2VuX3R5cGUiOiJhdXRoIiwiZXhwIjoxNjQ0Nzc5NTA5LCJkZXZpY2UtaWQiOiI4OTM5NjhhNi04ZmVlLTQyMDEtOTk2NC1mMWUyOTE5YWYxMDciLCJpYXQiOjE2NDQ3Nzg5MDksImp0aSI6ImY0YTcxMTI3LWI4ZWUtNDAwZi05NDUyLTY1NjNjYzNlMWQ0MyJ9.iZoKjh_FkcPZPKCTZWNpAr4lGop4aME8tlIqTdcy1A1RR2j21DrOvAkatb9n58Xpx74wi5xdQE6_r2_khKkoWbnmEHDxOAH4NNRnj_9uoJzjj1U6dKiInmSjmNXopm-t9XMjhXTSy1e6vnLqs4Tk920XRrd5Hfr8wn5nY_vN74oRasVFlsNegl4K0BMFVww9C893l01-zjoqNbXb49m-paj0O9l8b1PeZ_iDlhU9h9x635RXEYaBuBtmPm7t6-VbeEsaNjUSyDOCz9Kc4eufzv2gAFgKQ2bTH_79LFB4UXKBJBlFpdMTSDA2rjnlQ8j-7_OOqSpB65OxjHGWCXbGeA
x-amcn-cache-hash: f2d50469102e2f5aa19bceefa107147385fc26f35913874d27a75526d7c629bb
accept-encoding: gzip
user-agent: okhttp/4.9.0
x-newrelic-id: VgUEUVJXDhADXFhRAQkCV1I=
content-length: 0
~~~

This works:

~~~
GET https://gw.cds.amcn.com/content-compiler-cr/api/v1/content/amcn/bbca/type/season-episodes/id/1010621?device=mobile HTTP/2.0
tracestate: @nr=0-2-2330446-389368516-08d3ff55e923435b--00--1644779038118
traceparent: 00-a767d69200054e829735108164d28b92-08d3ff55e923435b-00
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjIzMzA0NDYiLCJkLmFwIjoiMzg5MzY4NTE2IiwiZC50ciI6ImE3NjdkNjkyMDAwNTRlODI5NzM1MTA4MTY0ZDI4YjkyIiwiZC5pZCI6IjdlOGY4ZGZlMWE2YjQyOWYiLCJkLnRpIjoxNjQ0Nzc5MDM4MTE4fX0=
x-amcn-device-id: 893968a6-8fee-4201-9964-f1e2919af107
x-amcn-language: en
x-amcn-network: bbca
x-amcn-platform: android
x-amcn-tenant: amcn
x-amcn-audience-id: amcn
x-amcn-device-ad-id: f187f55d-3f78-4267-8830-719a99cb7413
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-ccpa-do-not-sell: default
authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi01OC0xNzAuZWMyLmludGVybmFsIiwidG9rZW5fdHlwZSI6ImF1dGgiLCJleHAiOjE2NDQ3Nzk2MzUsImRldmljZS1pZCI6Ijg5Mzk2OGE2LThmZWUtNDIwMS05OTY0LWYxZTI5MTlhZjEwNyIsImlhdCI6MTY0NDc3OTAzNSwianRpIjoiY2Q4NzhkM2UtNjVkMC00ZmM0LTk2MzItY2I4NmIzYTc2OWY5In0.GmFDC0jxbh7S627L8CssGDyZwy7XLf9zBFMy1muCO2TB50jbeV3YlYLciJtX8jiPpKcxUFhOWo_JAbbvaatY219Pb5X4bQVTox01_OnD31nrSVSrP3TB_QqFnh9HYVqEZP0pd014LLtl5V7PmiC46l09FqYvkezP7RzUiBD2fOUQ24_rXBxLcgY2qw7WDz_sEHK3t8gA8tTxg00Mt0xWoRhTzN1NYyTnRW-b9EMWEcvk-9_MMUcZXcXF2oxjq2G3GUUIP8QQG8HjaeMOLfcyF0gSYXZW-HYPOMsXdTyKKJTzGhcJ2KPDo8ZdBcnPzMScqjT6q4qKBqVf632aTvZ_DA
x-amcn-cache-hash: f2d50469102e2f5aa19bceefa107147385fc26f35913874d27a75526d7c629bb
accept-encoding: gzip
user-agent: okhttp/4.9.0
x-newrelic-id: VgUEUVJXDhADXFhRAQkCV1I=
content-length: 0
~~~

Also this:

~~~
POST https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth HTTP/2.0
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjIzMzA0NDYiLCJkLmFwIjoiMzg5MzY4NTE2IiwiZC50ciI6ImRjMDY3MDEwNWMxMzQyNmM5ODBmNTI5ODUzMmUzMzJkIiwiZC5pZCI6IjUyNzViNTZlZTI2YTQ1OGUiLCJkLnRpIjoxNjQ0Nzc4OTExNTM2fX0=
tracestate: @nr=0-2-2330446-389368516-ac8bd5b050b64bbe--00--1644778911536
traceparent: 00-dc0670105c13426c980f5298532e332d-ac8bd5b050b64bbe-00
x-amcn-device-id: 893968a6-8fee-4201-9964-f1e2919af107
x-amcn-language: en
x-amcn-network: bbca
x-amcn-platform: android
x-amcn-tenant: amcn
x-amcn-audience-id: amcn
x-amcn-device-ad-id: f187f55d-3f78-4267-8830-719a99cb7413
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-ccpa-do-not-sell: default
content-length: 0
accept-encoding: gzip
user-agent: okhttp/4.9.0
x-newrelic-id: VgUEUVJXDhADXFhRAQkCV1I=
~~~
