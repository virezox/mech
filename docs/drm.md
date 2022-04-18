# DRM

- https://reddit.com/r/television
- https://trakt.tv/shows/orphan-black/seasons/1/episodes/2

## Amazon Prime Video

https://github.com/ytdl-org/youtube-dl/issues/1753

Web:

~~~
POST /cdp/catalog/GetPlaybackResources?deviceID=6d0a26c59307054d7db77c897b6b43a8e
Host: atv-ps.amazon.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Fir
Accept: */*
Accept-Language: en-US,en;q=0.5
Origin: https://www.amazon.com
DNT: 1
Connection: keep-alive
Referer: https://www.amazon.com/
Cookie: session-id=134-0209150-2708545; session-id-time=2082758401l; ubid-main=13
Content-Length: 0
~~~

I tried removing these, but still DRM:

~~~
deviceDrmOverride
supportedDRMKeyScheme
~~~

Android:

https://play.google.com/store/apps/details?id=com.amazon.avod.thirdpartyclient

Same request as web client.

## BBC America

Android client returns 400:

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

Web client gets a 404:

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

## Roku

Web client:

~~~
POST https://therokuchannel.roku.com/api/v3/playback HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Fir
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
content-type: application/json
content-length: 237
csrf-token: LyGryUqa-Scv8uHPkGI7tjKzTDD8SA8zOV6w
x-roku-reserved-session-id: 05fc8849-e41c-4426-bf1b-2d2a5abb4535
x-roku-reserved-experiment-state: W10=
x-roku-reserved-time-zone-offset: -05:00
x-roku-reserved-rida: 0a7b449a-fb09-5972-ac18-3e6b6c881904
x-roku-reserved-lat: 1
origin: https://therokuchannel.roku.com
dnt: 1
referer: https://therokuchannel.roku.com/watch/32c95b576307502b98f7fe32c4aa0a22
cookie: _csrf=SFcRr4kVNcLvpAKbIyg5Zpkg
cookie: watch.locale=j%3A%7B%22language%22%3A%22en%22%2C%22country%22%3A%22US%22%
cookie: ks.locale=j%3A%7B%22language%22%3A%22en%22%2C%22country%22%3A%22US%22%7D
cookie: _usn=05fc8849-e41c-4426-bf1b-2d2a5abb4535
cookie: amoeba=TE1mMDFuSzJQI0NvbnRyb2wscjc5UVFPU0VnI0NvbnRyb2wsWlZCSW95Q25xI0Nvbn
cookie: my.state=j%3A%7B%22signin_post_redirect%22%3A%22%2Fsignup%2Fpayment%22%2C
cookie: _uc=bc636024-f020-46f0-b1d7-4a9e08e1c7f1%3Aed4fa6b7d14f8d9b71663a8c359133
cookie: ks.session=BzqQRzEkKGSqeyotwyXmj9yoM29lMbvK5VMdPUyYmXq2MU3w43bOduraRE5KaT
cookie: AWSELB=2B01C54F0CCDD680B517F82B818B0F0AC248947D8EFEF5D8E5861938196639ED9B
cookie: AWSELBCORS=2B01C54F0CCDD680B517F82B818B0F0AC248947D8EFEF5D8E5861938196639
te: trailers

{"rokuId":"32c95b576307502b98f7fe32c4aa0a22","playId":"s-amcplus_tv_svod.en.WVcxa
~~~

## Spectrum

Web client:

~~~
POST /ipvs/api/smarttv/stream/vod/v2/bbcamerica.com::BBCH1782402270017003?csid=st
Host: api.spectrum.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Fir
Accept: application/json, text/plain, */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: identity
Content-Type: application/json
Authorization: OAuth oauth_account_type="RESIDENTIAL", oauth_consumer_key="l7xx66
device_id: 0f3c9082-27a6-4430-bfa3-8fb6c05f6a95
Origin: https://watch.spectrum.net
DNT: 1
Connection: keep-alive
Referer: https://watch.spectrum.net/

{"drmEncodings":[{"drm":"cenc","encoding":"dash"}]}
~~~

Android client makes same request as web client.

https://play.google.com/store/apps/details?id=com.TWCableTV
