# BBC America

Using this video:

https://bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529

## Web client

This is it:

~~~
GET https://ssaimanifest.prod.boltdns.net/us-east-1/playback/once/v1/dash/live-timeline/bccenc/6240731308001/8ebcf878-2abe-4ca8-8edc-9e46cdb0e6b8/01af0a57-214d-4fdd-86fd-f792135ce46f/19ec62cf-9c55-45c6-a8b5-6e7e639bae85/content.mpd?bc_token=NjI4ZTkwYThfMTViZDkwMzM2ZGM1OTMxNmNmMjE4Mzk3ZGU1ODU0YjE1MzFiMmMwNjQyMGEzZWYyNDIyMjk0YzVkYWEwMGMzMw%3D%3D&rule=discos-enabled HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
origin: https://www.bbcamerica.com
dnt: 1
referer: https://www.bbcamerica.com/
te: trailers
content-length: 0
~~~

From:

~~~
GET https://ssaimanifest.prod.boltdns.net/playback/once/v1/vmap/dash/live-timeline/bccenc/6240731308001/8ebcf878-2abe-4ca8-8edc-9e46cdb0e6b8/01af0a57-214d-4fdd-86fd-f792135ce46f/content.vmap?bc_token=NjI4ZWViNDNfMWJmM2QzYzM1YWEzMWY2ZDcwMjUyYTZkMzg5NzhkNjA2N2ZjZDI2NWM4YjYzY2RhNTZmOWEyZjgwMGQyOTA2OA%3D%3D&behavior_id=17308414-5636-4b1c-8c21-a168289f0440&csid=bbca_web&fw_did=f4b9e72e-2ec3-47a1-a8a6-5f2115ee6da3&idtype=gauid&prof=bbca_bc_web&mode=on-demand&caid=AMCNVR0000040191&islat=0&width=1920&height=1080&usprivacy=1YN-&ae=-1&ipaddress=72.181.23.38&rule=discos-enabled HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
origin: https://www.bbcamerica.com
dnt: 1
referer: https://www.bbcamerica.com/
te: trailers
content-length: 0
~~~

From:

~~~
POST https://gw.cds.amcn.com/playback-id/api/v1/playback/1052529 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: gzip, deflate, br
referer: https://www.bbcamerica.com/
authorization: Bearer eyJraWQiOiJwcm9kLTEiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJlbnRpdGxlbWVudHMiOiJ1bmF1dGgiLCJhdWQiOiJyZXNvdXJjZV9zZXJ2ZXIiLCJhdXRoX3R5cGUiOiJiZWFyZXIiLCJyb2xlcyI6WyJ1bmF1dGgiXSwiaXNzIjoiaXAtMTAtMi0xMTEtNjcuZWMyLmludGVybmFsIiwidG9rZW5fdHlwZSI6ImF1dGgiLCJleHAiOjE2NTM1MDczMDcsImRldmljZS1pZCI6IjAzMWE3ZWQ0LTQ0NTUtNDZkNS05YjJhLTBmMTVkODRkNzVlYyIsImlhdCI6MTY1MzUwNjcwNywianRpIjoiNzU0NDU1MmEtNzcwYy00OGM1LWI1ZTgtMzY0MzkwMjgyOGExIn0.otEDejVgDHnkKuo-Ya5hm5b46ZENk1BC0S7964JV7fG9d-NB1Pnu_k6eQyLxmZ5BCErlcPIABbG6couXZ1C4cxRjn0R9N5XBRCs585SNo2C7XrjkN3ScxnTmv_5axocapKkSfm3QkDKv9BRHhUBuLeE7HTC61WuN4DZWFwVYJ_ro2b_o1cKtceXNo7PaP_krgBjq61c0InqB5Vxr4fnIQ_L3-yOLgLbkXlI7ficsmTrrAaKHEFSSK6HmiVEoF3qpM2ciZ76i4PkSBCg5n73TjbahybAPNstbRMMnVk8lEUlTeR3t92KbIk5iWWArDJ8YODOn6hiPxIFy8cd3Rm1REw
content-type: application/json
x-amcn-device-id: f4b9e72e-2ec3-47a1-a8a6-5f2115ee6da3
x-amcn-device-ad-id: f4b9e72e-2ec3-47a1-a8a6-5f2115ee6da3
x-amcn-service-id: bbca
x-amcn-service-group-id: 6
x-amcn-tenant: amcn
x-amcn-network: bbca
x-amcn-platform: web
x-amcn-mvpd: 
x-amcn-adobe-id: 
x-amcn-audience-id: 
x-ccpa-do-not-sell: passData
x-amcn-cache-hash: 8679a33532b1f9b3310c9af5f95e855ed49159900277aeb54aa8da1e5a5c445e
x-amcn-language: en
origin: https://www.bbcamerica.com
content-length: 241
dnt: 1
te: trailers

{"adobeShortMediaToken":"","hba":false,"adtags":{"lat":0,"url":"https://www.bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529","playerWidth":1920,"playerHeight":1080,"ppid":1,"mode":"on-demand"},"useLowResVideo":false}
~~~
