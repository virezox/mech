# BBC America

## Unlocked

https://bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529

We can get these:

~~~
POST /playback-id/api/v1/playback/1052529 HTTP/1.1
Host: gw.cds.amcn.com
Authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xMDktMTkxLmVjMi5pbnRlcm5hbCIsInRva2VuX3R5cGUiOiJhdXRoIiwiZXhwIjoxNjUzNjAwODgzLCJkZXZpY2UtaWQiOiIhIiwiaWF0IjoxNjUzNjAwMjgzLCJqdGkiOiJiOTBiNWNjZS1mMGYzLTRhYjUtOTkzYi1lZjFjMDM0NjBhYTUifQ.O-nyBK3MPksSglvQAmPIZlxb-OFXlGEYyjLv9HDyBbjwurxTlYba1os08-Wriee4TKzknByW1ePiJ7Au_imWEYIAlvVYKSTjOCtfiTH5caG_n3CYhK1Vglwf4O0ci4ZqwjWTz7SVSCRQXwrKWXSiBk3sc6MqLSliE7Gb8Wle_oCC6EhSzN8Hd8PB1IrZ9_IRBlsfVLYVWBj4L9DfTQZFeTWoPnaUqnvCl0PGOo8N1rdwE6xFFGW5zjaSlfaRgitQeZwmeiHqr_H2eSMb6vqtPzHhAGH3rXPv0VdvudgsV2DM9HW5fDJNKgV14i2TbaU73VMRI6BjcrSjmthfT9zjnw
Content-Type: application/json
X-Amcn-Device-Ad-Id: !
X-Amcn-Language: en
X-Amcn-Network: bbca
X-Amcn-Platform: web
X-Amcn-Service-Id: bbca
X-Amcn-Tenant: amcn
X-Ccpa-Do-Not-Sell: passData
Content-Length: 92

{"adtags":{"lat":0,"mode":"on-demand","ppid":0,"playerHeight":0,"playerWidth":0,
"url":"!"}}

HTTP/2.0 200 OK
~~~

## Locked

https://bbcamerica.com/shows/orphan-black/episodes/season-1-instinct--1011152

~~~
POST /playback-id/api/v1/playback/1011152 HTTP/1.1
Host: gw.cds.amcn.com
Authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xMDktMTkxLmVjMi5pbnRlcm5hbCIsInRva2VuX3R5cGUiOiJhdXRoIiwiZXhwIjoxNjUzNjAwODgzLCJkZXZpY2UtaWQiOiIhIiwiaWF0IjoxNjUzNjAwMjgzLCJqdGkiOiJiOTBiNWNjZS1mMGYzLTRhYjUtOTkzYi1lZjFjMDM0NjBhYTUifQ.O-nyBK3MPksSglvQAmPIZlxb-OFXlGEYyjLv9HDyBbjwurxTlYba1os08-Wriee4TKzknByW1ePiJ7Au_imWEYIAlvVYKSTjOCtfiTH5caG_n3CYhK1Vglwf4O0ci4ZqwjWTz7SVSCRQXwrKWXSiBk3sc6MqLSliE7Gb8Wle_oCC6EhSzN8Hd8PB1IrZ9_IRBlsfVLYVWBj4L9DfTQZFeTWoPnaUqnvCl0PGOo8N1rdwE6xFFGW5zjaSlfaRgitQeZwmeiHqr_H2eSMb6vqtPzHhAGH3rXPv0VdvudgsV2DM9HW5fDJNKgV14i2TbaU73VMRI6BjcrSjmthfT9zjnw
Content-Type: application/json
X-Amcn-Device-Ad-Id: !
X-Amcn-Language: en
X-Amcn-Network: bbca
X-Amcn-Platform: web
X-Amcn-Service-Id: bbca
X-Amcn-Tenant: amcn
X-Ccpa-Do-Not-Sell: passData
Content-Length: 92

{"adtags":{"lat":0,"mode":"on-demand","ppid":0,"playerHeight":0,"playerWidth":0,
"url":"!"}}

HTTP/2.0 500 Internal Server Error
~~~

If I log in, it still fails:

~~~
POST https://gw.cds.amcn.com/playback-id/api/v1/playback/1011152 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Fir
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
referer: https://www.bbcamerica.com/
authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.ey
content-type: application/json
x-amcn-device-id: c4201f0e-5cab-415b-b30b-5af467694c0c
x-amcn-device-ad-id: c4201f0e-5cab-415b-b30b-5af467694c0c
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-amcn-tenant: amcn
x-amcn-network: bbca
x-amcn-platform: web
x-amcn-mvpd: Spectrum
x-amcn-adobe-id: AD541738-C817-38FF-E4B1-5464EC1A7C07
x-amcn-audience-id: 
x-ccpa-do-not-sell: passData
x-amcn-cache-hash: f701a9de5ffe5924e361b81f0f77c6dc6949859d08880ac03f56222c7f3e75
x-amcn-language: en
origin: https://www.bbcamerica.com
content-length: 257
dnt: 1
te: trailers

{"adobeShortMediaToken":"W29iamVjdCBPYmplY3Rd","hba":true,"adtags":{"lat":0,
"url":"https://www.bbcamerica.com/shows/orphan-black/episodes/season-1-instinct--1011152",
"playerWidth":1920,"playerHeight":1080,"ppid":1,"mode":"on-demand"},
"useLowResVideo":false}

HTTP/2.0 400 Bad Request
~~~
