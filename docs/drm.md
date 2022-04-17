# DRM

- https://reddit.com/r/television
- https://trakt.tv/shows/orphan-black/seasons/1/episodes/2

## Amazon Prime Video

Android:

https://play.google.com/store/apps/details?id=com.amazon.avod.thirdpartyclient

Web:

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

https://github.com/ytdl-org/youtube-dl/issues/1753

## Apple TV

https://github.com/ytdl-org/youtube-dl/issues/30808
