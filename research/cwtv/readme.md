# CWTV

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

then:

~~~
GET /nosec/The_CW/386/980/132474949570/4400-112-GroupEfforts-P112-CW.m3u8 HTTP/1.1
Host: 3aa37dc0e8bb47e08042e0ebb25acb34.dlvr1.net
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
GET /nosec/The_CW/386/980/132474949570/4400-112-GroupEfforts-P112-CW_132473413843_m3u8_video_1920x1080_8000000_primary_audio_eng_x3b089552e15a45689836dc3c5b75b903_8.m3u8 HTTP/1.1
Host: stream-hls.cwtv.com
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
GET /nosec/The_CW/386/980/132474949570/4400-112-GroupEfforts-P112-CW_132473413843_m3u8_video_1920x1080_8000000_primary_audio_eng_8_x3b089552e15a45689836dc3c5b75b903_00001.ts HTTP/1.1
Accept-Encoding: identity
Host: stream-hls.cwtv.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.17 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
