# Widevine

https://github.com/ytdl-org/youtube-dl/issues/30873

Using this video:

https://www.cda.pl/video/391634853

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f 0 --allow-unplayable-formats `
http://vwaw760.cda.pl/3916348/3916348.mpd
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
   <cenc:pssh>
   AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEGM2NDI1NGNkYjdlZGUzZTBI49yVmwY=
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST /license-proxy-widevine/cenc/?specConform=true HTTP/1.1
Host: lic.drmtoday.com
x-dt-custom-data: eyJ1c2VySWQiOiJhbm9uaW0iLCJzZXNzaW9uSWQiOiJrRGJZZUlYeVM4VjNt...
~~~

Now go to Get Widevine Keys [2], and enter the information from above:

~~~
PSSH:
AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEGM2NDI1NGNkYjdlZGUzZTBI49yVmwY=

License:
https://lic.drmtoday.com/license-proxy-widevine/cenc/?specConform=true

Headers:
x-dt-custom-data: eyJ1c2VySWQiOiJhbm9uaW0iLCJzZXNzaW9uSWQiOiJrRGJZZUlYeVM4VjNt...
~~~

You should get a result like this:

~~~
63363432353463646237656465336530:66363435386631333730396636386565
~~~

Finally, you can decrypt [3] the media:

~~~
mp4decrypt --key 63363432353463646237656465336530:663634353866... enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://getwvkeys.cc
3. https://bento4.com/downloads
