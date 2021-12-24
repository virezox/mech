# NBC

Get stream information:

~~~
PS C:\> nbc -hi 9000199366
POST http://access-cloudpath.media.nbcuni.com/access/vod/nbcuniversal/9000199366
GET https://east.manifest.na.theplatform.com/m/NnzsPC/T1Q_Z0YvTepY,thXzeZMxUtNl,E5dGFwKVZB7N,hB6R4we_olVJ,csINwUhvOCSf,RsLMUFAzCWSW.m3u8?sid=84203611-9cb8-4749-ad17-9ee202f2b4ce&policy=189081367&date=1640374038471&ip=72.181.23.38&schema=1.1&manifest=M3U&tracking=true&switch=HLSServiceSecure&am_sdkv=null&p2=null&_fw_did=&nw=169843&f1=&am_extmp=default&uuid=optout-160589338476416&am_appv=null&mode=on-demand&uoo=1&sfid=9244572&s3=null&am_buildv=null&am_abvrtd=null&s4=poc&_fw_ae=&debug=false&csid=oneapp_phone_android_app_ondemand&metr=1023&am_cpsv=4.0.0-2&bundleId=&userAgent=Go-http-client%2F1.1&e1=default&afid=200265138&prof=nbcu_android_cts_bl&c3=null&a2=4.0.0-2&a3=null&a5=null&a6=0&am_crmid=null&rdid=android&am_playerv=null&did=optout&am_stitcherv=poc&am_abtestid=null&sig=1d040d5e0d82c6f55af1549371cda44aec7f8ad3e4324c0d9ed4a2faf59a82de
ID:0 BANDWIDTH:2190000 CODECS:avc1.4d001f,mp4a.40.2 RESOLUTION:960x540
ID:1 BANDWIDTH:8750000 CODECS:avc1.640028,mp4a.40.2 RESOLUTION:1920x1080
ID:2 BANDWIDTH:5532000 CODECS:avc1.64001f,mp4a.40.2 RESOLUTION:1280x720
ID:3 BANDWIDTH:3478000 CODECS:avc1.4d001f,mp4a.40.2 RESOLUTION:960x540
ID:4 BANDWIDTH:1082000 CODECS:avc1.4d001e,mp4a.40.2 RESOLUTION:768x432
~~~

Download stream:

~~~
nbc -h 2 9000199366
~~~

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358
