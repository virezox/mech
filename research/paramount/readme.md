# Paramount+

Need proxy to use with Get Widevine Keys:

~~~
Error 404: {"code":130301,"message":"Content is not available in this
location."}
~~~

https://github.com/ytdl-org/youtube-dl/issues/29038

Using this video:

<https://paramountplus.com/shows/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU>

Download the MPD:

~~~
yt-dlp -o enc.mp4 --allow-unplayable-formats -f 17 `
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2021/08/31/1940767811923/993595_cenc_dash/stream.mpd
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
   <cenc:pssh>
   AAAAWHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADgIARIQO+i+k3yYSDGEspQXP5FSryIgZXlUX1JZa3FOdUhfNlpZcmVwTHR4a2lQTzFIQTdkSVU4AQ==
   </cenc:pssh>
</ContentProtection>
~~~

1. <https://dashif.org/identifiers/content_protection>

Now go back to the video page, and you should see a request like this:

~~~
POST https://cbsi.live.ott.irdeto.com/widevine/getlicense?CrmId=cbsi&AccountId=cbsi&SubContentType=Default&ContentId=eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
content-length: 2
origin: https://www.paramountplus.com
dnt: 1
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjQ2NjA2OSwiZXhwIjoxNjUyNDczMjY5LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiIzNTMwYjViOC0wODIyLTQ5N2YtYTdkMC0yMDgwMWU1ODZmN2UifQ.onQORR8tOB0IE0QhKOqipvU7F6im-MZuxTQJCzBozGg
referer: https://www.paramountplus.com/
te: trailers
~~~

https://github.com/weapon121/WKS-KEY/releases/tag/WKS-KEY

Basically you just slap the headers in headers.py and run python l3.py and give
it the PSSH and `Lic_URL`
