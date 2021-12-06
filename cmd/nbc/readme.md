# NBC

One:

~~~
POST /access/vod/nbcuniversal/9000199358 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
Authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=f7abce51aad6c7f7cee46d580a492d039d1d57bad5e3355deebbc307cd5c4396,time=1638763299504
Content-Type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC","mpx":{"accountId":2410887629}}
~~~

Two:

~~~
GET /m/NnzsPC/P3SxIJ4UKaJ7,x_U_3FRM6_mL,E8eDNLskrRwg,74sfimEHVQvk,ThnunlAOwYup,EfiqbgG8AVZ8.m3u8?sid=88ac4e8e-c1d0-4b42-9876-46184e1ed125&policy=189081367&date=1638763299537&ip=72.181.23.38&schema=1.1&manifest=M3U&tracking=true&switch=HLSServiceSecure&am_sdkv=null&p2=null&_fw_did=&nw=169843&f1=&am_extmp=default&uuid=optout-224138223010581&am_appv=null&mode=on-demand&uoo=1&sfid=9244572&s3=null&am_buildv=null&am_abvrtd=null&s4=poc&_fw_ae=&debug=false&csid=oneapp_phone_android_app_ondemand&metr=1023&am_cpsv=4.0.0-2&bundleId=&userAgent=Go-http-client%2F1.1&e1=default&afid=200265138&prof=nbcu_android_cts_bl&c3=null&a2=4.0.0-2&a3=null&a5=null&a6=0&am_crmid=null&rdid=android&am_playerv=null&did=optout&am_stitcherv=poc&am_abtestid=null&sig=4fd3deb1effe538c8131cc73d81e9dae8cbdce2d63a239d23aff36491c731d38 HTTP/1.1
Host: east.manifest.na.theplatform.com
~~~

Three:

~~~
POST /v2/graphql HTTP/1.1
Host: friendship.nbc.co
Content-Type: application/json

{"extensions":{"persistedQuery":{"sha256Hash":"73014253e5761c29fc76b950e7d4d181c942fa401b3378af4bac366f6611601e"}},
"variables":{"app":"nbc","name":"9000199358","platform":"android","type":"VIDEO","userId":""}}
~~~
