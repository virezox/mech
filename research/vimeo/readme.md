# Vimeo

- https://github.com/ytdl-org/youtube-dl/issues/29690
- https://github.com/ytdl-org/youtube-dl/issues/30622
- https://vimeo.com/477957994/2282452868

~~~
GET /video/477957994/config?api=0&autopause=0&autoplay=0&background=0&badge=0&byline=1&bypass_privacy=1&collections=0&color=00adef&context=Vimeo%5CController%5CApi%5CResources%5CVideoController.&controls=0&default_to_hd=0&force_embed=0&h=2282452868&info_on_pause=0&js_api=0&like=0&logo=0&loop=0&max_height=0&max_width=0&muted=0&outro=nothing&playbar=0&portrait=1&privacy_banner=0&quality=540p&requested_height=0&requested_width=0&responsive=1&responsive_width=1&share=0&speed=1&title=1&volume=0&watch_later=0&s=9a3452c8547f591fd26e4685174eacc889b6d72f_1644521888 HTTP/1.1
Host: player.vimeo.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.74 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

Comes from this:

~~~
GET /videos/477957994:2282452868?fields=config_url%2Ccreated_time%2Cdescription%2Clicense%2Cmetadata.connections.comments.total%2Cmetadata.connections.likes.total%2Crelease_time%2Cstats.plays HTTP/1.1
Host: api.vimeo.com
Authorization: jwt eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ0MjIwNDAsInVzZXJfaWQiOm51bGwsImFwcF9pZCI6NTg0NzksInNjb3BlcyI6InB1YmxpYyIsInRlYW1fdXNlcl9pZCI6bnVsbH0.A_fJ6syeghDdWI3tpoDlRiLevTX_0HgZ78KwRxFonDo
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.74 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

Comes from this:

~~~
GET /_rv/jwt HTTP/1.1
Host: vimeo.com
X-Requested-With: XMLHttpRequest
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.74 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
