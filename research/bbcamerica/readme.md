# BBC America

https://github.com/ytdl-org/youtube-dl/issues/30182

## Android client

https://play.google.com/store/apps/details?id=com.bbca.bbcafullepisodes

## Web client

Using this video:

https://bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f cdbdaf53-d750-457b-93fc-0090aabbbe3e-0 `
--allow-unplayable-formats `
'https://ssaimanifest.prod.boltdns.net/us-east-1/playback/once/v1/dash/live-timeline/bccenc/6240731308001/8ebcf878-2abe-4ca8-8edc-9e46cdb0e6b8/01af0a57-214d-4fdd-86fd-f792135ce46f/c84cd254-2a7d-42ab-ae7a-665ca19239e3/content.mpd?bc_token=NjI2NDBhZGFfNGY2MDc0ZTA2NmIyYTJmY2Q5MDM3NTVlNDBlNGJhMWU2ODE5ODM2ZmExYzdjOWU2YmIyNmE2ZTI4MzI1ODk1Yg%3D%3D&rule=discos-enabled'
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
bc:licenseAcquisitionUrl="https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6240731308001/01af0a57-214d-4fdd-86fd-f792135ce46f/883780c4-a981-494c-b994-9e93792ff8a7?fastly_token=NjI2NDZiMzVfNGUxMDY1MjI4ZWJkYmFlMzc5YjVlZjVkZTM0MjlmZDE1YTEyNjc3NWJkNmIwOWNhNGEwZjg3MmM1ZmEzZTEyOQ%3D%3D">
   <cenc:pssh>
   AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQiDeAxKmBSUy5lJ6TeS/4pxoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6240731308001/01af0a57-214d-4fdd-86fd-f792135ce46f/883780c4-a981-494c-b994-9e93792ff8a7?fastly_token=NjI2NDY2ZDJfYmUxYTFkYWFlMjgwNWNkNjVkODdjNDZkOGIxZTQyNjIwZGRlNWQ5ZDIyMGJmMDcwYTc5NTRjOGM3M2IzZjNlYg%3D%3D HTTP/2.0
bcov-auth: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhY2NpZCI6IjYyNDA3MzEzMDgwMD...
~~~

Now go to Get Widevine Keys, and enter the information from above:

~~~
PSSH:
AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQiDeAxKmBSUy5lJ6TeS/4pxoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=

License:
https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6240731308001/01af0a57-214d-4fdd-86fd-f792135ce46f/883780c4-a981-494c-b994-9e93792ff8a7?fastly_token=NjI2NDY2ZDJfYmUxYTFkYWFlMjgwNWNkNjVkODdjNDZkOGIxZTQyNjIwZGRlNWQ5ZDIyMGJmMDcwYTc5NTRjOGM3M2IzZjNlYg%3D%3D

Headers:
bcov-auth: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhY2NpZCI6IjYyNDA3MzEzMDgwMD...
~~~

You should get a result like this:

~~~
883780c4a981494cb9949e93792ff8a7:680a46ebd6cf2b9a6a0b05a24dcf944a
~~~

Finally, you can decrypt [2] the media:

~~~
mp4decrypt --key 883780c4a981494cb9949e93792ff8a7:680a46ebd6cf... enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://bento4.com/downloads
