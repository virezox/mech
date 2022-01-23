# NBC

Why am I not getting HD anymore? This request gets HD:

~~~
GET /m/HNK2IC/gXFdWKpkhfJ_,NTOfvaDxclgP,oXbMtTEQ0_qw,yB2QIcy3g5Wk,7q_fWxi_Yq8p,C9WlXvcAQdMR.m3u8?sid=4329b7f7-a6f7-4c76-bf47-12420a78123b&policy=188057562&date=1642947721595&ip=72.181.23.38&schema=1.1&manifest=M3U&tracking=true&switch=HLSServiceSecure&am_sdkv=android_4.10.71.2&_fw_did=google_advertising_id%3A76004574-fc3f-4683-9430-df6598dda1fb&f1=google_advertising_id%3A76004574-fc3f-4683-9430-df6598dda1fb&am_extmp=default&uuid=76004574-fc3f-4683-9430-df6598dda1fb&am_appv=nbc_7.28.1&mode=on-demand&uoo=0&w1=1920&sfid=9244572&s3=android_4.10.71.2&am_buildv=7.28.1&am_abvrtd=null&s4=poc&_fw_ae=&mParticleId=-7855216639727670072&csid=oneapp_phone_android_app_ondemand&metr=1023&comscore_device=Android_SDK_built_for_x86&afid=200265138&c3=-7855216639727670072&nielsen_platform=MBL&_fw_nielsen_app_id=PAD3C6E72-ED61-417F-A865-3AB63FDB6197&am_crmid=-7855216639727670072&am_playerv=exoplayer_2.11.8&am_abtestid=null&p1=&p2=exoplayer_2.11.8&nw=169843&h1=1080&comscore_did_x=76004574-fc3f-4683-9430-df6598dda1fb&nielsen_device_group=PHN&comscore_platform=android&debug=false&m1=-7855216639727670072&am_cpsv=4.0.0-2&bundleId=com.nbcuni.nbc&userAgent=Mozilla%2F5.0+%28Linux%3B+Android+7.0%3B+Android+SDK+built+for+x86+Build%2FNYC%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F69.0.3497.100+Mobile+Safari%2F537.36&e1=default&prof=nbcu_android_cts_bl&a2=4.0.0-2&a3=nbc_7.28.1&a5=7.28.1&a6=0&rdid=e3a71011a55b81b9&android_id=76004574-fc3f-4683-9430-df6598dda1fb&did=76004574-fc3f-4683-9430-df6598dda1fb&am_stitcherv=poc&sig=527bb661672c20bd61184eeb60616d2846aa9634c931cf14a008402ffffb2158 HTTP/1.1
User-Agent: CloudpathPlayer/7.28.1 (Linux;Android 7.0) ExoPlayerLib/2.11.8
Accept-Encoding: gzip
Host: east.manifest.na.theplatform.com
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
POST /access/vod/nbcuniversal/9000199368 HTTP/1.1
user-agent: Mozilla/5.0 (Linux; Android 7.0; Android SDK built for x86 Build/NYC; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 Mobile Safari/537.36
accept: application/access-v1+json
cache-control: no-cache
app-session-id: dc5dac8c-ca7f-4ef6-88f8-0dc529bacc46
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=3669f8e808b30e255271c4af18ef8504a82ec2e7a899e4791484b166ba290502,time=1642947721954
content-type: application/json
Host: access-cloudpath.media.nbcuni.com
Connection: Keep-Alive
Accept-Encoding: gzip
Content-Length: 815

{"device":"android","auth":false,"adobeMvpdId":"","deviceId":"e3a71011a55b81b9","device_type":"google_advertising_id","nw":"169843","mParticleId":"-7855216639727670072","externalAdvertiserId":"NBC_VOD_9000199368","did":"e3a71011a55b81b9","uuid":"76004574-fc3f-4683-9430-df6598dda1fb","appv":"NBC_7.28.1","buildv":"7.28.1","am_appv":"7.28.1","am_buildv":"2000002882","player_height":"1080","player_width":"1920","sdkv":"android_4.10.71.2","playerv":"exoplayer_2.11.8","bundleId":"com.nbcuni.nbc","us_privacy_string":"","mpx":{"accountId":"2304985974"},"tracking":{"deviceGroup":"PHN","platform":"MBL","appId":"PAD3C6E72-ED61-417F-A865-3AB63FDB6197","androidId":"76004574-fc3f-4683-9430-df6598dda1fb","comscore_device":"Android_SDK_built_for_x86","googleAdId":"76004574-fc3f-4683-9430-df6598dda1fb"},"prefetch":false}
~~~
