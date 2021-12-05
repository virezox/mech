# December 4 2021

With some videos:

https://nbc.com/the-blacklist/video/the-skinner/9000210182

The Android request returns 1920x1080:

~~~
POST /access/vod/nbcuniversal/9000210182 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=3ad75ff5acb36579df917dd2bcc92b8e11ada78715257cc8339836572b6504fc,time=1638662573514
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2304985974}}
~~~

With other videos:

https://nbc.com/la-brea/video/pilot/9000194212

The Android request returns 960x540:

~~~
POST /access/vod/nbcuniversal/9000194212 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=422af527fadc9f2cde373a267b65c4f09e895cf593186820e43b33e25b600b65,time=1638662778578
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2304985974}}
~~~

While the web request for same video returns 1920x1080:

~~~
GET /s/NnzsPC/media/guid/2410887629/9000194212?manifest=m3u HTTP/1.1
Host: link.theplatform.com
~~~

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://www.nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358
