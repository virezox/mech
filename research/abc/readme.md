# ABC

- https://abc.com/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you
- https://github.com/ytdl-org/youtube-dl/issues/29544

## Android

~~~
com.disney.datg.videoplatforms.android.abc
~~~

Install system certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://abc.com/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you
~~~

this is it:

~~~
GET /ausw/slices/037/d874124ecca24c88a3c9575e78686acf/0376293b438b4e9d9472606fa38d98bb/D00000013.m4f?pbs=882b17f400ef4c3a9aa0e41b3817d8a9&drm=1&cloud=aws&si=1&d=4.096&cdn=eci HTTP/1.1
X-NewRelic-ID: VQIEVFdTGwcDXVFQDggG
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Accept-Encoding: identity
Host: x-disney-datg-stgec.uplynk.com
Connection: Keep-Alive
content-length: 0
~~~

from:

~~~
GET /api/v3/preplay2/0376293b438b4e9d9472606fa38d98bb/e37fe20bfe144625ee37601cd2cf4537/3R3zwdMznFEC9beJbldE896RXCTOU2dm02LuT9ZxRdmU.mpd?pbs=882b17f400ef4c3a9aa0e41b3817d8a9&cdn=ec HTTP/1.1
X-NewRelic-ID: VQIEVFdTGwcDXVFQDggG
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Accept-Encoding: gzip
Host: content-ausw1-ur-dtci2.uplynk.com
Connection: Keep-Alive
content-length: 0
~~~

from:

~~~
POST https://prod.gatekeeper.us-abc.symphony.edgedatg.com/vp2/ws-secure/entitlement/2020/g/playmanifest_secure.json HTTP/2.0
user-agent: com.disney.datg.videoplatforms.android.abc/10.23.1 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
datg-usertz: -0500
accept: application/json
appversion: 10.23.1
authorization: JWT eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI1ZThlNzQwZi0xZTlkLTQ4MmItYTc4NS1jNzM3N2Q2YWM4NzAiLCJpYXQiOjE2NDg3Njc5OTEsInN1YiI6InA5MGI0MGI0Yi1jYTAwLTQwMzItYjhiOC04ZDg1Yzc2MzEzMDkiLCJpc3MiOiJhYmMifQ.Vx6lLXTAF3my45j7115SO5rtAlaTTseCoxJvYGP9dKo
content-type: application/x-www-form-urlencoded; charset=utf-8
content-length: 1747
accept-encoding: gzip
cookie: SWID=B0C7F874-E673-4BFC-A2AF-59D77554D839
x-newrelic-id: VQIEVFdTGwcDXVFQDggG

ad.paln=AQzzBGQECSeN5jY-braA3gRSmJ_cVT7zxeUrdlGizMjkyATMlbEqWUZvR8haRQcZ1mItxSMoqiBsSr9guBSvPBYaojugxsOMLiTl2-PgDWyd4i3XJG_DD4yAn_djFZvBtWiwdBfbrr2COxiG46LvfKppDYmDEmRfZqVhW1_OnI16QlirCH7r_ZDRlfT1PrTRRb4iJRQlCtwXhXlLmRbHIOY9pV9_M-Bgz4ZPEkyNYIlkKh_jIbApx5qSzUD6rG7u4KFCACD-U8SEWd8opKI0qho_GNZBaZtaOUkJf8q2tqKJ7vFUI9ssku4EtRElN-7_7hYyatH3Hcgt9KeqhMLsdpJRqLI6xpopefCnp7071Vb9BEFsv4U5VtUUoSxCLxcvYNst3-Z68YgIcf3sIiUkm2YJWV46Gc42Q9-0HhtPmi1y37hK9w6h4IfHJFBtaQEErtPp2QD15NtmbR-iiSfgVXdUkguhUmWJAhD1xVLJlD1jmtjKCdQyNqTq4eOwdrKAS8mGqvkZRlSVIX-vp4Tao5oKPf7MK805bnbnBCPrRe9NVRp_znHgF9njyZlkd8FvlSCCkEUyttBoDJPsFFq3n-clxkLUKnb5DNYjSThhaxI76_-jDs2k5dnTHgqgr7eYuKFqXp_5KNi3566uZM8psHGPrw3vrOc4FqWzHKVLZwtJA2jHHx9bSHZsZgW6kixZI0cSzXvWcUg4gjNyYXYaW9V_xDr8NueIbNXP-UXmaT55S0n9mhMB4pwqT4fyKGQPYGrqhpjl27hmRhqCjdDTc7Dc_NhlyJYbJ4U4OYKzEzfnFH0QVuZ2IJeAK88pzo7X0yyX3oedgddoUCpaTsARqYowG5TyAcTUlDi4TOGOdMm6mtiL5IqhKQgpDsHT8s2MQrMdaikR3dJ6a0nDfH_b6mcHKVpx8yDCjGbqHMYkyC1tT_EYxzugDQRqsmYAQ1LqbJItoEfel5NnC6y1fci9GMJ1c9R2Nq_LtRO-XghRWX60tM1IpoCb-XCvLtweJI16Jy-ujyzriquLE7IEjzhPIWiD5ib723aPfRer_hsuBDn6B0nlieMKfNKftIfMIVlvnynZyQ%3D%3D&adid=0b70ba69-8caa-403e-af2b-846d96d7a4f4&affiliate=WFAA&airing_id=VDKA26847512&app_name=abcv2&brand=001&bundleId=com.disney.datg.videoplatforms.android.abc&device=031_04&deviceOS=7.0&deviceType=Google+Android+SDK+built+for+x86&device_id=391c1b66bbb2ce6f2ae2bfde591fdc7e&hdcp_level=0.0&hip=3571f507982d4eba645235d75c3b0d984ae542962795413bce1afee4b82c430e&hlsver=6&isAutoplay=false&latitude=0.0&longitude=0.0&maxBitrate=2500&minBitrate=0&nielsenAppId=PDFB2C928-9709-4DBE-B612-7C7480302256&player_id=9f0ced04-6783-4fe9-b5e0-4bd2fa0a744b&prefBitrate=10&token_type=ap&tracking=1&video_id=VDKA26847512&video_type=lf&vps=1920x1080&zipcode=75039
~~~

from:

~~~
GET https://prod.gatekeeper.us-abc.symphony.edgedatg.com/api/ws/pluto/v1/layout/route?authlevel=0&brand=001&country=usa&defaultlanguage=en-US&device=031_04&distributionchannel=2&url=%2Fshows%2Fgreys-anatomy%2Fepisode-guide%2Fseason-18%2F12-the-makings-of-you HTTP/2.0
user-agent: com.disney.datg.videoplatforms.android.abc/10.23.1 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
datg-usertz: -0500
accept: application/json
appversion: 10.23.1
accept-encoding: gzip
cookie: SWID=B0C7F874-E673-4BFC-A2AF-59D77554D839
x-newrelic-id: VQIEVFdTGwcDXVFQDggG
content-length: 0
~~~

## Web

this is it:

~~~
GET /ext/d874124ecca24c88a3c9575e78686acf/5e6dda216a5e43309543689d3fd85f91.m3u8?cqs=SzomztIPTJ7Udx3oYBccK29_qhQFcZQqZjxzdzxH87QGecQlFtA9qFoOkko7hDK2TsOXzzCtEwG1NI8tXzdccinbncF8JrcLToUFgxjGUtpsZCNgl4TApXaWZGFaYPEjGxMfWh62vza2iMPE7p4u7-uwFNMTQpA4w5HSYHee_scZOIiav0ClvcYQqu6WxBy9S4QUuk4PYT4x_go3VP56uQv8r_ExhBOQlShg173zu-nRCJAsef-WV7r_zOc0g_k3v9Mle8-tw_G3dT6jr7J7Bl6HhTv4Wm6_7ySPWvdY_wGooTLwgCAAaOT1v8LUZPZhJE9Ihin_Lvs550CC_44kO_FPD6jvMX-BePovNPgxNs4IQ9MsDDIsxUOOig4az_5MzTkYfIQnTHQTPMxnnLDWsyo_xIKyuNdnKqZWNhl33X0BXsWY_c08aKccx3B2nkIXN_lCiMGAXjuTVVcQeWFUIxrAkbb5UFsRpcMyayhez_c=&kid=4098a6a720374bfcbb4e362b652bcd51 HTTP/1.1
Host: content-dtci.uplynk.com
X-Forwarded-For: 6.174.126.101
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3574.0 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close
content-length: 0
~~~

from:

~~~
GET /vp2/ws/contents/3000/videos/001/001/-1/-1/-1/VDKA26847512/-1/-1.json HTTP/1.1
Host: api.contents.watchabc.go.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3574.0 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close
content-length: 0
~~~
