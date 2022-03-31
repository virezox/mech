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
