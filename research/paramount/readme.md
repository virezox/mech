# Paramount

> Theatricality and deception, powerful agents to the uninitiated.
>
> But we are initiated, aren’t we, Bruce?
>
> The Dark Knight Rises (2012)

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
authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY1MjQ4NzM3NSwiZXhwIjoxNjUyNDk0NTc1LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwiaXNlIjp0cnVlLCJqdGkiOiI1MjA2NGJhYS03MDAwLTRjYjQtYjRjNS1iNDUyYzE5NzQ3OTMifQ.g3g52ntnnRKrcCYX_2bJMCzljWnUrQujD1YGvQbeSzQ
~~~

<https://github.com/Jnzzi/4464_L3-CDM>

## CMAC

https://github.com/dchest/cmac
