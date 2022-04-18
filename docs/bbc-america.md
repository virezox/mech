# BBC America

## Android client

returns 400:

~~~
POST https://gw.cds.amcn.com/playback-id/api/v1/playback/1053337?debug=true&debug
traceparent: 00-70b016229f4f40eb8f41dc9170183527-3764fc7e0c3e471c-00
tracestate: @nr=0-2-2330446-389368516-3764fc7e0c3e471c--00--1650255039754
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjIzMzA0NDYiLCJkLmF
x-amcn-device-id: 537a3bf0-3a3b-43bd-b063-e8f47209d2f0
x-amcn-language: en
x-amcn-network: bbca
x-amcn-platform: android
x-amcn-tenant: amcn
x-amcn-audience-id: amcn
x-amcn-device-ad-id: 4e2a2bf7-2a38-4a8e-8e86-cf4ed9205435
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-ccpa-do-not-sell: default
x-amcn-mvpd: Spectrum
x-amcn-adobe-id: F16F30B0-3162-44D5-F70D-FE200DB43CFA
authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.ey
x-amcn-cache-hash: 1829b289f9b81b541d22a80a995526b6a6d927505545cde9f48430f1c5c6a2
content-type: application/json; charset=UTF-8
content-length: 185
accept-encoding: gzip
user-agent: okhttp/4.9.0
x-newrelic-id: VgUEUVJXDhADXFhRAQkCV1I=

{"adobeShortMediaToken":"","adtags":{"amznParams":"","lat":0,"mode":"on-demand","
~~~

https://play.google.com/store/apps/details?id=com.bbca.bbcafullepisodes

## Web client

returns 404:

~~~
GET /p/M_UwQC/64jR_L5FWBcE?location=https://www2.bbcamerica.com&autoPlay=true HTTP/1.1
Host: player.theplatform.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: identity
DNT: 1
Connection: keep-alive
Referer: https://www2.bbcamerica.com/
Upgrade-Insecure-Requests: 1
Pragma: no-cache
Cache-Control: no-cache
content-length: 0
~~~
