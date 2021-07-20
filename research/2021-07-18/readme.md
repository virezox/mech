# July 18 2021

Can we do `/get_video_info` with `Authorization` header? Yes:

~~~
curl -v -H 'Authorization: Bearer ya29.a0ARrdaM9gbcgLGqyyJDKV3agV60hXypXRdj...' `
'https://www.youtube.com/get_video_info?c=TVHTML5&cver=7.20210428.10.00&' +
'el=detailpage&html5=1&video_id=Cr381pDsSsA'
~~~

Can we do `/get_video_info` with `Cookie` header? Yes:

~~~
curl -v `
-H 'Cookie: __Secure-3PSID=_wdNhho4fTGOQgwrBshGc4QT3onztCB6xvfzwFh7OrWNTMjn...' `
('https://www.youtube.com/get_video_info?c=TVHTML5&cver=7.20210428.10.00&' +
'el=detailpage&html5=1&video_id=Cr381pDsSsA')
~~~

Can we do `/get_video_info` with query parameter? Yes:

~~~
curl -v ('https://www.youtube.com/get_video_info?c=TVHTML5&' +
'cver=7.20210428.10.00&el=detailpage&html5=1&video_id=Cr381pDsSsA&' +
'access_token=ya29.a0ARrdaM9gbcgLGqyyJDKV3agV60hXypXRdjspznK-4N2JF9nzW6z7BS9...')
~~~

Can we do `/youtubei/v1/player` with `Authorization` header? Yes:

~~~
curl -v `
-H 'Authorization: Bearer ya29.a0ARrdaM-R9c0crGrwdmS-FRugGQbDdRtkVEM5YFOm6d...' `
-H 'content-type: application/json' `
-d '@youtube.json' `
https://www.youtube.com/youtubei/v1/player
~~~

Can we do `/youtubei/v1/player` with `Cookie` header? Yes:

~~~
curl -v `
-H 'Authorization: SAPISIDHASH 1626641976_8b25e923dce0638caf2ef1d7dbac1e253...' `
-H 'Content-Type: application/json' `
-H 'Cookie: __Secure-3PSID=_wdNhnFIjWLd...; __Secure-3PAPISID=Y35R-aq0Bud-...' `
-H 'X-Origin: https://www.youtube.com' `
-d '@youtube.json' `
'https://www.youtube.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8'
~~~

Can we do `/youtubei/v1/player` with query parameter? Yes:

~~~
curl -v -H 'Content-Type: application/json' -d '@youtube.json' `
'https://www.youtube.com/youtubei/v1/player?access_token=ya29.a0ARrdaM8XMbTHQ...'
~~~

- https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps
- https://github.com/pytube/pytube/issues/1057
