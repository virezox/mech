# Roku

https://github.com/ytdl-org/youtube-dl/issues/23645

Using this video:

https://therokuchannel.roku.com/watch/2b9a65b56b425c62b5e43775eaefb830

Download the MPD:

~~~
yt-dlp -o enc.mp4 --allow-unplayable-formats -f 9 `
https://vod.delivery.roku.com/009395caab254571bb4d13906bcaf350/b222fadf974b441d9fe0c73210dce69a/e67a2e2bff1f4ab985b60b2e66f938d0/index.mpd
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
   <cenc:pssh>
   AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQkuvrLzP1OOGuxjrBuIhQTBoKaW50ZXJ0cnVzdCIBKg==
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKfIAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIAG4rzB7tkB8ashwAfaCBU2RAtTTrWLOSrtEQIo-hktBUWNflEwiT4FRRMUyPAnNDRPRCYzj4W8PSgFCthMBFClDaSwvL1Np4CkeyUewZIDmIai4tE0Kc2LyWyg16TqyFtiuQG-roDVDM6yyOOzXNbSWchdSg7MolQEd393UqH3VSb_A5IVJiy3w3h933Wy6myjE4PE HTTP/2.0
~~~

Now go to Get Widevine Keys, and enter the information from above:

~~~
PSSH:
AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQkuvrLzP1OOGuxjrBuIhQTBoKaW50ZXJ0cnVzdCIBKg==

License:
https://wv.service.expressplay.com/hms/wv/rights/?ExpressPlayToken=BQA1P5QRKfIAJDIzN2U4NTE4LTQwN2QtNDI3Zi05NTkyLWFmMTJiMzRkMmU0NwAAAIAG4rzB7tkB8ashwAfaCBU2RAtTTrWLOSrtEQIo-hktBUWNflEwiT4FRRMUyPAnNDRPRCYzj4W8PSgFCthMBFClDaSwvL1Np4CkeyUewZIDmIai4tE0Kc2LyWyg16TqyFtiuQG-roDVDM6yyOOzXNbSWchdSg7MolQEd393UqH3VSb_A5IVJiy3w3h933Wy6myjE4PE
~~~

You should get a result like this:

~~~
92ebeb2f33f538e1aec63ac1b888504c:0179c3e147f31fd2a1aa666477ca9344
~~~

Finally, you can decrypt [2] the media:

~~~
mp4decrypt --key 1:0179c3e147f31fd2a1aa666477ca9344 enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://bento4.com/downloads
