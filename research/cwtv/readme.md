# CWTV

- <http://images.cwtv.com/feed/mobileapp/video-meta/apiversion_8/guid_deec61a8-e0a1-4c01-8906-4e0b363350d5>
- https://github.com/ytdl-org/youtube-dl/issues/30662
- https://play.google.com/store/apps/details?id=com.cw.fullepisodes.android
- https://www.cwtv.com/shows/4400/?play=deec61a8-e0a1-4c01-8906-4e0b363350d5

~~~
GET /feed/mobileapp/video-meta/apiversion_8/guid_deec61a8-e0a1-4c01-8906-4e0b363350d5 HTTP/1.1
Host: images.cwtv.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.17 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

then:

~~~
GET /s/cwtv/media/guid/2703454149/deec61a8-e0a1-4c01-8906-4e0b363350d5?format=SMIL&formats=M3U&tracking=true&mbr=false HTTP/1.1
Host: link.theplatform.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.17 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
