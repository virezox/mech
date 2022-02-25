# MTV

- https://github.com/ytdl-org/youtube-dl/issues/30678
- https://mtv.com/video-clips/s5iqyc/mtv-cribs-dj-khaled
- https://play.google.com/store/apps/details?id=com.mtvn.mtvPrimeAndroid

~~~
api.viacom.tech -language:html
~~~

https://github.com/OpenKD/repository.cobramod/blob/main/repo/plugin.video.viacom.mtv/resources/lib/provider.py

Install system certificate. This is it:

~~~
GET /h/a/dG9wYXoxYGM5M2FmYWIxLTYyYTctMTFlYy1hNGYxLTcwZGYyZjg2NmFjZWA5ZTY0NzY4OTI3ZjY3NDAzMGQ5YWY3YjhmMDM1YmZkMmY1MDVmNjM0YHFEelBYNC03cW5weENyc0dCaGQyMTNDQ3lPTkEzLVR3bWJjaFdmb3owMXM/segment-14-f5-v1.ts HTTP/1.1
Accept-Encoding: identity
Connection: Keep-Alive
Host: mtv.orchestrator.viacomcbs-tech.com
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
X-NewRelic-ID: VQMGWFZaDBABVFdTAgADVVA=
content-length: 0
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjE1MTkwOTQiLCJkLmFwIjoiMTA2MTUxNDM0IiwiZC50ciI6ImM0OWQ4YTY1MmZjOTQzMmViYzg0MjU1ZDMwZDIxYjNjIiwiZC5pZCI6ImE4ZTE1NTJmMGQ1ZDQ3ZjgiLCJkLnRpIjoxNjQ1NzU2NDczMTMxfX0=
traceparent: 00-c49d8a652fc9432ebc84255d30d21b3c-a8e1552f0d5d47f8-00
tracestate: @nr=0-2-1519094-106151434-a8e1552f0d5d47f8----1645756473131
~~~

which comes from:

~~~
GET /ondemand/hls/content/2565036/vid/mgid:arc:episode:android.playplex.mtv.com:c93afab1-62a7-11ec-a4f1-70df2f866ace/TUL/streams/7396001c-fbd4-4418-92ec-a7faa8e19405/media/15115289387e1e34ef7ee8f62528587a.m3u8?hdntl=exp=1645842837~acl=%2f*~id=db4dc46b-795b-4a6f-9c4d-f7a53c2688ad~data=hdntl~hmac=35df712533eda2832787dc01c245bf148e8ad0113e91564f116ed19fd9d1e4fa HTTP/1.1
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjE1MTkwOTQiLCJkLmFwIjoiMTA2MTUxNDM0IiwiZC50ciI6IjliMGJmNjk0MTUzYTRiYjI5YjIzZTYyNzM3MzFjZTEzIiwiZC5pZCI6IjdkZjczOTgxYjRkYzQyNWEiLCJkLnRpIjoxNjQ1NzU2NDM5MjE1fX0=
tracestate: @nr=0-2-1519094-106151434-7df73981b4dc425a----1645756439215
traceparent: 00-9b0bf694153a4bb29b23e6273731ce13-7df73981b4dc425a-00
Host: topaz.dai.viacomcbs.digital
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: hdntl=exp=1645842837~acl=%2f*~id=db4dc46b-795b-4a6f-9c4d-f7a53c2688ad~data=hdntl~hmac=35df712533eda2832787dc01c245bf148e8ad0113e91564f116ed19fd9d1e4fa
X-NewRelic-ID: VQMGWFZaDBABVFdTAgADVVA=
content-length: 0
~~~

which comes from:

~~~
GET /h/a/dG9wYXoxYGM5M2FmYWIxLTYyYTctMTFlYy1hNGYxLTcwZGYyZjg2NmFjZWA5ZTY0NzY4OTI3ZjY3NDAzMGQ5YWY3YjhmMDM1YmZkMmY1MDVmNjM0YHFEelBYNC03cW5weENyc0dCaGQyMTNDQ3lPTkEzLVR3bWJjaFdmb3owMXM/master.m3u8?hdnea=st=1645755836~exp=1645759436~acl=/h/a/dG9wYXoxYGM5M2FmYWIxLTYyYTctMTFlYy1hNGYxLTcwZGYyZjg2NmFjZWA5ZTY0NzY4OTI3ZjY3NDAzMGQ5YWY3YjhmMDM1YmZkMmY1MDVmNjM0YHFEelBYNC03cW5weENyc0dCaGQyMTNDQ3lPTkEzLVR3bWJjaFdmb3owMXM/*~id=db4dc46b-795b-4a6f-9c4d-f7a53c2688ad~hmac=865ac57982303ae768e75eb128f703e276c74ece7aef7a1bc5cb3ce3a79f13c6&originpath=/ondemand/hls/content/2565036/vid/mgid:arc:episode:android.playplex.mtv.com:c93afab1-62a7-11ec-a4f1-70df2f866ace/TUL/streams/7396001c-fbd4-4418-92ec-a7faa8e19405/master.m3u8 HTTP/1.1
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
traceparent: 00-b00e10e9d5a44f2b9b21cbef57016a6c-873c9c080be94b58-00
tracestate: @nr=0-2-1519094-106151434-873c9c080be94b58----1645756436638
newrelic: eyJ2IjpbMCwyXSwiZCI6eyJkLnR5IjoiTW9iaWxlIiwiZC5hYyI6IjE1MTkwOTQiLCJkLmFwIjoiMTA2MTUxNDM0IiwiZC50ciI6ImIwMGUxMGU5ZDVhNDRmMmI5YjIxY2JlZjU3MDE2YTZjIiwiZC5pZCI6Ijg3M2M5YzA4MGJlOTRiNTgiLCJkLnRpIjoxNjQ1NzU2NDM2NjM4fX0=
Host: topaz.dai.viacomcbs.digital
Connection: Keep-Alive
Accept-Encoding: gzip
X-NewRelic-ID: VQMGWFZaDBABVFdTAgADVVA=
content-length: 0
~~~
