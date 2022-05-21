# Paramount

how to get MPD?

<https://paramountplus.com/shows/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU>

~~~
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU
~~~

## How to get `pid`

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU
?format=preview
~~~

Gives this:

~~~
QqofwoCXmQIE
~~~

Does that work?

## Web client

This is `pid`:

~~~
SDhY2iU6ETkC
~~~

its in the source of URL above, can we get it another way?

## Android client

This is it:

~~~
GET https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/TUL/streams/bb1b3cdf-c007-4c74-9a66-e5cd1a7d9339/manifest.mpd HTTP/2.0
user-agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
tracestate: @nr=0-2-2936348-766594812-bd68c403b3504a2a----1653170050800
traceparent: 00-72353910d56443e6b713f9c30e1cc947-bd68c403b3504a2a-00
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjI5MzYzNDgiLCJkLmFwIjoiNzY2NTk0ODEyIiwiZC50ciI6IjcyMzUzOTEwZDU2NDQzZTZiNzEzZjljMzBlMWNjOTQ3IiwiZC5pZCI6ImJkNjhjNDAzYjM1MDRhMmEiLCJkLnRpIjoxNjUzMTcwMDUwODAwfX0=
accept-encoding: identity
x-newrelic-id: Vg8EV1VXABAHUldXDgUPV1Y=
content-length: 0
~~~

From:

~~~
POST https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/streams HTTP/2.0
content-length: 1556
origin: https://imasdk.googleapis.com
user-agent: Mozilla/5.0 (Linux; Android 7.0; Android SDK built for x86 Build/NYC; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 Mobile Safari/537.36
content-type: application/x-www-form-urlencoded;charset=UTF-8
accept: */*
referer: https://imasdk.googleapis.com/native/sdkloader/native_sdk_v3.html?sdk_version=a.3.24.0&hl=en&omv=1.3.3-google_20200416&app=com.cbs.app
accept-encoding: identity
accept-language: en-US
x-requested-with: com.cbs.app

imafw__fw_is_lat=0&imafw__fw_app_bundle=com.cbs.app&imafw_csid=streaming_paramountplus_mobile_androidphone_vod&imafw_cpPre=0&imafw_playername_version=avia_3.5.33&imafw__fw_did=google_advertising_id%3A0a95e7a2-26f1-428d-a278-763568f49e85&imafw_sb=0&imafw_session=b&imafw_cpSession=0&imafw_sz=640x480&imafw_fms_ifa=0a95e7a2-26f1-428d-a278-763568f49e85&imafw_ima_sdkv=3.24.0&imafw__fw_continuous_play=0&imafw__fw_coppa=0&tfcd=0&imafw__fw_vcid2=0a95e7a2-26f1-428d-a278-763568f49e85&imafw_fms_vcid2type=ifa&imafw_description_url=https%3A%2F%2Fwww.paramountplus.com&imafw_subses=3&imafw_vguid=99c7c3ab-d828-438c-9a98-67fa82293ab4&imafw_section=free-content-hub&imafw__fw_h_referer=com.cbs.app&imafw_tfcd=0&api-key=7j3n273bfj7jcnvf22d713a51f&ms=CoACnCJdEORo57-oD7XOn3G6l5Ro3ULWTpeCWffghJ05MFhifrU3Dvzb7MhkeVQABHqKoTu51ipa1hw3u2vatK3iSD86tkCL910MALkIGbOTpZnLOCctM4pbtHb1Py2U0OoEpPrUWbQtObJIe85bw9LGLKrWqHj6-OscdV6pMxlrUJKj-gdE1jRWEJ-NlK7knpQrrx_fnuLZ1Uv-0qbnJSGUOSwRwZO7f33NiGlIOE6PcBLImTda5--s9fpdFEvFMAGheNnPJ9YbJIn5cRR0i8zq9IA7Q8doksBLFVuZf2ocq4kBlgI9XtZ4IOFs3KHC4F2CTtY39OzjkOVmd1uyWWdwPxIQAQfrG_qdRfIIyJabrt8raw&js=ima-android.3.24.0&correlator=3634640901539583&mpt=AVIA_Player&mpv=avia_3.5.33&ptt=20&osd=2&sas=1&sdki=405&sdkv=h.3.310.0%2Fn.android.3.24.0%2Fcom.cbs.app&uach=null&an=com.cbs.app&msid=com.cbs.app&eid=44758266%2C44760950%2C44761692%2C44762903&frm=-1&omid_p=Google1%2Fandroid.3.24.0&sdk_apis=7%2C8&sid=8736A1F5-FD70-49AF-84F8-F8230F6B9444&url=com.cbs.app.adsenseformobileapps.com%2F&is_lat=0&idtype=adid&rdid=0a95e7a2-26f1-428d-a278-763568f49e85
~~~
