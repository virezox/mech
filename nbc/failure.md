# Failure

## Android

With some videos:

https://nbc.com/the-blacklist/video/the-skinner/9000210182

The Android request returns 1920x1080:

~~~
POST /access/vod/nbcuniversal/9000210182 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
User-Agent: Mozilla/5
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=e7b31c766165b06e422e0e3ca9723d2ab5eb9861f5aea48dca78393eb33aa9fd,time=1638665774851
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2410887629}}
~~~

With other videos:

https://nbc.com/la-brea/video/pilot/9000194212

The Android request returns 960x540:

~~~
POST /access/vod/nbcuniversal/9000194212 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
User-Agent: Mozilla/5
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=43d1ebdb5dfe3a21b9d76f38d370cd83d8316076987d581be74d43562840aca1,time=1638665836328
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2410887629}}
~~~

While the web request for same video returns 1920x1080:

~~~
GET /s/NnzsPC/media/guid/2410887629/9000194212?manifest=m3u HTTP/1.1
Host: link.theplatform.com
~~~

## Progressive

If you make a request like this:

http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL

You get results like this:

~~~html
<video
src="http://nbcmpx-vh.akamaihd.net/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_7830k.mp4?hdnea=st=1638675793~exp=1638688423~acl=/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_*~id=f46974e2-ba14-44ef-8e3e-5764799dc427~hmac=e1e983df3726aad47d33f5d328216f96d068c8fc67386a86efdb806a774af0bb"
system-bitrate="7932473" height="1080" width="1920"/>
~~~

However it appears none of the links actually work:

~~~
Access Denied
You don't have permission to access
"http://nbcmpx-vh.akamaihd.net/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_7830k.mp4?"
on this server.
Reference #18.100ec617.1638675890.1159b69a 
~~~

Also, these requests:

- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL&switch=http
- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL&switch=progressive

Return empty links:

~~~html
<video src="" system-bitrate="7932473" height="1080" width="1920"/>
~~~

https://github.com/ytdl-org/youtube-dl/issues/7806
