## Spectrum

Web client:

~~~
POST /ipvs/api/smarttv/stream/vod/v2/bbcamerica.com::BBCH1782402270017003?csid=st
Host: api.spectrum.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Fir
Accept: application/json, text/plain, */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: identity
Content-Type: application/json
Authorization: OAuth oauth_account_type="RESIDENTIAL", oauth_consumer_key="l7xx66
device_id: 0f3c9082-27a6-4430-bfa3-8fb6c05f6a95
Origin: https://watch.spectrum.net
DNT: 1
Connection: keep-alive
Referer: https://watch.spectrum.net/

{"drmEncodings":[{"drm":"cenc","encoding":"dash"}]}
~~~

Android client makes same request as web client.

https://play.google.com/store/apps/details?id=com.TWCableTV
