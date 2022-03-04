# TED

## Android client

~~~
com.ted.android
~~~

Install system certificate. Then go here with Android Chrome:

https://ted.com

and click on one of the videos. A prompt should come up that says "Open with".
Click "TED", then "JUST ONCE". The video should open in the app, and if you are
monitoring, you should see this request:

~~~
GET https://devices.ted.com/api/v2/videos/rha_goddess_and_deepa_purushothaman_4_ways_to_redefine_power_at_work_to_include_women_of_color/react_native_v2.json HTTP/2.0
accept: application/json
cache-control: no-cache
accept-encoding: gzip
user-agent: okhttp/4.9.1
content-length: 0
~~~

## Why does this exist?

February 11 2022:

https://github.com/ytdl-org/youtube-dl/issues/30561
