# Paramount+

https://github.com/ytdl-org/youtube-dl/issues/30491

## How to get endpoint?

~~~
https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
~~~

https://play.google.com/store/apps/details?id=com.cbs.app

Install user certificate, then use Frida:

https://github.com/httptoolkit/frida-android-unpinning

Example two (`/media/guid/` not found):

- <https://paramountplus.com/shows/star-trek-prodigy/video/_9JdOMR84RwqcRtCup9UynfldbPH4LdH/star-trek-prodigy-janeway-upgraded-s1-e10-paramount->
- https://link.theplatform.com/s/dJ5BDC/Jl7FSgX2t2JT?format=SMIL&manifest=m3u&Tracking=true&mbr=true

~~~
https://can-services.cbs.com/canServices/playerService/video/search.xml?
partner=cbs&
contentId=_9JdOMR84RwqcRtCup9UynfldbPH4LdH
~~~

Example one (`/media/guid/` not found):

- <https://paramountplus.com/shows/the-harper-house/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/the-harper-house-the-harper-house>
- https://link.theplatform.com/s/dJ5BDC/SDhY2iU6ETkC?format=SMIL&manifest=m3u&Tracking=true&mbr=true

Android uses this:

~~~
GET /s/dJ5BDC/fNsRH_fjko5T?format=SMIL&Tracking=true&sig=0062224f03ccac6b6a1501b010c706455919c82f06fe441ab0706f63 HTTP/1.1
X-NewRelic-ID: VQ4FVlJUARABVVRXAwEOVFc=
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Host: link.theplatform.com
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

which comes from:

~~~
GET https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json?locale=en-us&at=ABAJi4xSDPXIEUKTlJ6BFQpMdL3hrvn5xbm%2BXly%2B9QZJFycgSL%2F4%2FYiDMKY4XWomRkI%3D HTTP/2.0
cache-control: no-cache
tracestate: @nr=0-2-1827479-115540923-----1646415500980
traceparent: 00-276e7b3dc2004b57bec2509a17c96d57--00
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjE4Mjc0NzkiLCJkLmFwIjoiMTE1NTQwOTIzIiwiZC50ciI6IjI3NmU3YjNkYzIwMDRiNTdiZWMyNTA5YTE3Yzk2ZDU3IiwiZC5pZCI6ImRmNmQwNmRiODNiMjRiZWYiLCJkLnRpIjoxNjQ2NDE1NTAwOTgwfX0=
accept-encoding: gzip
user-agent: okhttp/4.9.0
x-newrelic-id: VQ4FVlJUARABVVRXAwEOVFc=
content-length: 0
~~~
