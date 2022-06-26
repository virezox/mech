# Amazon Prime Video

https://github.com/ytdl-org/youtube-dl/issues/1753

## Android

https://play.google.com/store/apps/details?id=com.amazon.avod.thirdpartyclient

Same request as web client.

## Web

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

I tried removing these, but still protected:

~~~
deviceDrmOverride
supportedDRMKeyScheme
~~~

Using this video:

https://www.amazon.com/gp/video/detail/B00TJZ3RH0

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f video=100000-1 --allow-unplayable-formats `
https://a306avoddashs3ww-a.akamaihd.net/d/2$47ezd62yJ7jlYq1ThUxmuUIzCbE~/1@f1...
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:EDEF8BA9-79D6-4ACE-A3C8-27DCD51D21ED">
   <cenc:pssh>
   CAESEKd0rwuuV0+rvRVR2ZDUqqMaBmFtYXpvbiI1Y2lkOms0SXl6bk9VUittWDZsVkNoMTcwd0E9PSxwM1N2QzY1WFQ2dTlGVkhaa05TcW93PT0qAlNEMgA=
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST /cdp/catalog/GetPlaybackResources?deviceID=92e9a8b30b18815dbbf5... HTTP/1.1
Host: atv-ps.amazon.com
~~~

Now go to Get Widevine Keys, and enter the information from above:

~~~
PSSH:
CAESEKd0rwuuV0+rvRVR2ZDUqqMaBmFtYXpvbiI1Y2lkOms0SXl6bk9VUittWDZsVkNoMTcwd0E9PSxwM1N2QzY1WFQ2dTlGVkhaa05TcW93PT0qAlNEMgA=

License:
https://atv-ps.amazon.com/cdp/catalog/GetPlaybackResources?deviceID=92e9a8b30...
~~~

You should get a result like this:

~~~
Error 404: {} 
~~~

1. <https://dashif.org/identifiers/content_protection>
