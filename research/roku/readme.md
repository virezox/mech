# Roku

https://github.com/ytdl-org/youtube-dl/issues/23645

Using this video:

https://therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f

Download the MPD:

~~~
yt-dlp -o enc.mp4 --allow-unplayable-formats -f 9 `
https://vod.delivery.roku.com/41e834bbaecb4d27890094e3d00e8cfb/aaf72928242741a6ab8d0dfefbd662ca/87fe48887c78431d823a845b377a0c0f/index.mpd
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKZgAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIBW-ZfZBFLrJdKgAFVJXA35OSjy4wtym39JdDx2a5QSebndwcLe7ji0mb8cxO4B0cWin3BPPiq_Xb1X1siMd9EnP4FhzcZu4yaWkM7q0kmgnRY5IcY1oZmiYYDWaNE7wKnDQWhrZKK_wmTDca9xwL19y3M4WASKwsnYr5WEj-dEeYihJ9RhCRmHZS-YKusGmLTEWghg HTTP/2.0
~~~

Now go to Get Widevine Keys, and after "License" enter the URL from above:

~~~
https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKcoAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIDnFzZs94Ig6XMgvBctSEO07eLZilJFfEdyMSl2GO2dt1QbpyMfjY0T1fY34jcGNH2OvTvOa2GDjkrj0sGVhPfBGwhPy5JzpfyIKDGZ0uwEb3710A_j4V87rQpHdufzhZeJoCeNjS6duPqmABFy91sH9CBRnsBBCtCdRsrBBp-jD8BwL6SDEsEXkYQjvDxSGQ1rAZ6J
~~~

You should get a result like this:

~~~
28339ad78f734520da24e6e0573d392e:13d7c7cf295444944b627ef0ad2c1b3c
~~~

Finally, you can decrypt [1] the media:

~~~
mp4decrypt `
--key 28339ad78f734520da24e6e0573d392e:13d7c7cf295444944b627ef0ad2c1b3c `
enc.mp4 dec.mp4
~~~

1. https://bento4.com/downloads

## CSRF

https://www.roku.com
