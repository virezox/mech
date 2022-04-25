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

I tried removing these, but still DRM:

~~~
deviceDrmOverride
supportedDRMKeyScheme
~~~

Using this video:

https://www.amazon.com/gp/video/detail/B00TJZ3RH0

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f video=100000-1 --allow-unplayable-formats `
'https://a306avoddashs3ww-a.akamaihd.net/d/2$47ezd62yJ7jlYq1ThUxmuUIzCbE~/1@f1b4830392ef7c2bb28625d4657360f1/ondemand/iad_2/f1d5/69e5/db5d/4b7c-92a2-f2f680d5cb99/83fb35a0-d889-483f-a1da-49bb5eb09a79_corrected.mpd?custom=true&amznDtid=AOAGZA014O5RE&encoding=segmentBase'
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
POST /cdp/catalog/GetPlaybackResources?deviceID=92e9a8b30b18815dbbf5d10e5d62cfcf41444be0a263d9b8f3c73b43&deviceTypeID=AOAGZA014O5RE&gascEnabled=false&marketplaceID=ATVPDKIKX0DER&uxLocale=en_US&firmware=1&clientId=f22dbddb-ef2c-48c5-8876-bed0d47594fd&deviceApplicationName=Firefox64&playerType=xp&operatingSystemName=Windows&operatingSystemVersion=10.0&asin=B00TJYYKD6&consumptionType=Streaming&desiredResources=Widevine2License&resourceUsage=ImmediateConsumption&videoMaterialType=Feature&userWatchSessionId=eacbfd5a-51db-4aed-bc5f-892c2a29dd7e&deviceProtocolOverride=Https&deviceStreamingTechnologyOverride=DASH&deviceDrmOverride=CENC&deviceBitrateAdaptationsOverride=CVBR%2CCBR&deviceAdInsertionTypeOverride=SSAI&deviceHdrFormatsOverride=None&deviceVideoCodecOverride=H264&deviceVideoQualityOverride=HD&playerAttributes=%7B%22middlewareName%22%3A%22Firefox64%22%2C%22middlewareVersion%22%3A%2288.0%22%2C%22nativeApplicationName%22%3A%22Firefox64%22%2C%22nativeApplicationVersion%22%3A%2288.0%22%2C%22supportedAudioCodecs%22%3A%22AAC%22%2C%22frameRate%22%3A%22HFR%22%2C%22H264.codecLevel%22%3A%224.2%22%2C%22H265.codecLevel%22%3A%220.0%22%7D HTTP/1.1
Host: atv-ps.amazon.com
~~~

Now go to Get Widevine Keys [2], and enter the information from above:

~~~
PSSH:
CAESEKd0rwuuV0+rvRVR2ZDUqqMaBmFtYXpvbiI1Y2lkOms0SXl6bk9VUittWDZsVkNoMTcwd0E9PSxwM1N2QzY1WFQ2dTlGVkhaa05TcW93PT0qAlNEMgA=

License:
https://atv-ps.amazon.com/cdp/catalog/GetPlaybackResources?deviceID=92e9a8b30b18815dbbf5d10e5d62cfcf41444be0a263d9b8f3c73b43&deviceTypeID=AOAGZA014O5RE&gascEnabled=false&marketplaceID=ATVPDKIKX0DER&uxLocale=en_US&firmware=1&clientId=f22dbddb-ef2c-48c5-8876-bed0d47594fd&deviceApplicationName=Firefox64&playerType=xp&operatingSystemName=Windows&operatingSystemVersion=10.0&asin=B00TJYYKD6&consumptionType=Streaming&desiredResources=Widevine2License&resourceUsage=ImmediateConsumption&videoMaterialType=Feature&userWatchSessionId=eacbfd5a-51db-4aed-bc5f-892c2a29dd7e&deviceProtocolOverride=Https&deviceStreamingTechnologyOverride=DASH&deviceDrmOverride=CENC&deviceBitrateAdaptationsOverride=CVBR%2CCBR&deviceAdInsertionTypeOverride=SSAI&deviceHdrFormatsOverride=None&deviceVideoCodecOverride=H264&deviceVideoQualityOverride=HD&playerAttributes=%7B%22middlewareName%22%3A%22Firefox64%22%2C%22middlewareVersion%22%3A%2288.0%22%2C%22nativeApplicationName%22%3A%22Firefox64%22%2C%22nativeApplicationVersion%22%3A%2288.0%22%2C%22supportedAudioCodecs%22%3A%22AAC%22%2C%22frameRate%22%3A%22HFR%22%2C%22H264.codecLevel%22%3A%224.2%22%2C%22H265.codecLevel%22%3A%220.0%22%7D
~~~

You should get a result like this:

~~~
Error 404: {} 
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://getwvkeys.cc
