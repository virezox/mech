# March 25 2022

~~~
PS D:\Desktop> youtube-views.exe^C
PS D:\Desktop> yt-dlp.exe HtVdAasjOgU
[youtube] HtVdAasjOgU: Downloading webpage
[youtube] HtVdAasjOgU: Downloading android player API JSON
[youtube] HtVdAasjOgU: Downloading android embedded player API JSON
[youtube] HtVdAasjOgU: Downloading web embedded config
[youtube] HtVdAasjOgU: Downloading player c6736352
[youtube] HtVdAasjOgU: Downloading web embedded player API JSON
[info] HtVdAasjOgU: Downloading 1 format(s): 248+251
[download] Resuming download at byte 4193280
[download] Destination: The Witcher 3 - Wild Hunt - The Sword Of Destiny Trailer [HtVdAasjOgU].f248.webm
[download]  16.1% of 37.29MiB at  5.50MiB/s ETA 00:05
~~~

versus:

~~~
PS D:\Desktop> youtube.exe -b HtVdAasjOgU
POST https://www.youtube.com/youtubei/v1/player
Status: LOGIN_REQUIRED
Reason: This video may be inappropriate for some users.
PS D:\Desktop> youtube.exe -b HtVdAasjOgU -e
POST https://www.youtube.com/youtubei/v1/player
Status: LOGIN_REQUIRED
Reason: This video may be inappropriate for some users.
~~~

They doing three for some reason, two work. First (10.08k):

~~~
POST /youtubei/v1/player?key=AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw&prettyPrint=false HTTP/1.1
Content-Length: 355
Host: www.youtube.com
Cookie: PREF=hl=en&tz=UTC; CONSENT=YES+cb.20210328-17-p0.en+FX+854; GPS=1; YSC=K2BdHNo_Yys; VISITOR_INFO1_LIVE=kK46uI5xLYc
X-Youtube-Client-Name: 55
X-Youtube-Client-Version: 16.49
Origin: https://www.youtube.com
Content-Type: application/json
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.115 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Accept-Encoding: gzip, deflate, br
Connection: close

{"context": {"client": {"clientName": "ANDROID_EMBEDDED_PLAYER", "clientVersion": "16.49", "hl": "en", "timeZone": "UTC", "utcOffsetMinutes": 0}, "thirdParty": {"embedUrl": "https://google.com"}}, "videoId": "HtVdAasjOgU", "playbackContext": {"contentPlaybackContext": {"html5Preference": "HTML5_PREF_WANTS"}}, "contentCheckOk": true, "racyCheckOk": true}
~~~

second (22.66k):

~~~
POST /youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false HTTP/1.1
Content-Length: 1085
Host: www.youtube.com
Cookie: PREF=hl=en&tz=UTC; CONSENT=YES+cb.20210328-17-p0.en+FX+854; GPS=1; YSC=K2BdHNo_Yys; VISITOR_INFO1_LIVE=kK46uI5xLYc
X-Youtube-Client-Name: 56
X-Youtube-Client-Version: 1.20220323.01.00
Origin: https://www.youtube.com
Content-Type: application/json
X-Goog-Visitor-Id: CgtrSzQ2dUk1eExZYyi5jPiRBg%3D%3D
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.115 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Accept-Encoding: gzip, deflate, br
Connection: close

{"context": {"client": {"hl": "en", "gl": "US", "remoteHost": "72.181.23.38", "deviceMake": "", "deviceModel": "", "visitorData": "CgtrSzQ2dUk1eExZYyi5jPiRBg%3D%3D", "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.115 Safari/537.36,gzip(gfe)", "clientName": "WEB_EMBEDDED_PLAYER", "clientVersion": "1.20220323.01.00", "osName": "Windows", "osVersion": "10.0", "originalUrl": "https://www.youtube.com/embed/HtVdAasjOgU?html5=1", "platform": "DESKTOP", "clientFormFactor": "UNKNOWN_FORM_FACTOR", "configInfo": {"appInstallData": "CLmM-JEGELfLrQUQ8IKuBRDUg64FEMOHrgUQ__etBRCS660FENi-rQU%3D"}, "timeZone": "UTC", "browserName": "Chrome", "browserVersion": "92.0.4515.115", "utcOffsetMinutes": 0}, "user": {"lockedSafetyMode": false}, "request": {"useSsl": true}, "clickTracking": {"clickTrackingParams": "IhMIoIvS0+7h9gIVBreCCh1NGQwA"}}, "videoId": "HtVdAasjOgU", "playbackContext": {"contentPlaybackContext": {"html5Preference": "HTML5_PREF_WANTS", "signatureTimestamp": 19075}}, "contentCheckOk": true, "racyCheckOk": true}
~~~
