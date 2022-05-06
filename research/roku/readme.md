# Roku

https://github.com/ytdl-org/youtube-dl/issues/23645

Using this video:

https://therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f

Download the MPD:

~~~
yt-dlp -o enc.mp4 --allow-unplayable-formats -f 9 `
https://vod.delivery.roku.com/41e834bbaecb4d27890094e3d00e8cfb/aaf72928242741a6ab8d0dfefbd662ca/87fe48887c78431d823a845b377a0c0f/index.mpd
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
   <cenc:pssh>
   AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQKDOa149zRSDaJObgVz05LhoKaW50ZXJ0cnVzdCIBKg==
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKZgAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIBW-ZfZBFLrJdKgAFVJXA35OSjy4wtym39JdDx2a5QSebndwcLe7ji0mb8cxO4B0cWin3BPPiq_Xb1X1siMd9EnP4FhzcZu4yaWkM7q0kmgnRY5IcY1oZmiYYDWaNE7wKnDQWhrZKK_wmTDca9xwL19y3M4WASKwsnYr5WEj-dEeYihJ9RhCRmHZS-YKusGmLTEWghg HTTP/2.0
~~~

Now go to Get Widevine Keys, and enter the information from above:

~~~
PSSH:
AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQKDOa149zRSDaJObgVz05LhoKaW50ZXJ0cnVzdCIBKg==

License:
https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKZgAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIBW-ZfZBFLrJdKgAFVJXA35OSjy4wtym39JdDx2a5QSebndwcLe7ji0mb8cxO4B0cWin3BPPiq_Xb1X1siMd9EnP4FhzcZu4yaWkM7q0kmgnRY5IcY1oZmiYYDWaNE7wKnDQWhrZKK_wmTDca9xwL19y3M4WASKwsnYr5WEj-dEeYihJ9RhCRmHZS-YKusGmLTEWghg
~~~

You should get a result like this:

~~~
28339ad78f734520da24e6e0573d392e:13d7c7cf295444944b627ef0ad2c1b3c
~~~

Finally, you can decrypt [2] the media:

~~~
mp4decrypt `
--key 28339ad78f734520da24e6e0573d392e:13d7c7cf295444944b627ef0ad2c1b3c `
enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://bento4.com/downloads
