# Peacock

https://github.com/ytdl-org/youtube-dl/issues/26018

Using this video:

<https://peacocktv.com/watch/playback/vod/GMO_00000000102539_01>

Capture MPD request:

~~~
GET /pub/global/FcS/xvQ/PCK_1616077455260.896_01/cmaf/mpeg_cenc_2sec/master_cmaf.mpd?c3.ri=3777436256273240854 HTTP/1.1
Host: g003-vod-us-cmaf-prd-cl.cdn.peacocktv.com
~~~

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f video_353422 --allow-unplayable-formats `
https://g003-vod-us-cmaf-prd-cl.cdn.peacocktv.com/pub/global/FcS/xvQ/PCK_1616077455260.896_01/cmaf/mpeg_cenc_2sec/master_cmaf.mpd
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
   <cenc:pssh>
   AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAWylyu1Ra6Obew32huzP9I49yVmwY=
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST https://ovp.peacocktv.com/drm/widevine/acquirelicense?bt=43-ZU8f9RX99J26WMHcSeoet2y5R9JqRXCSQxUE8r4SlCUOkFXg_iRz6yauwLz6yrS-n-CsDA28kePZkgXsaK4RR7c5WIokXlv12bqfUpwUrZwXMK3YGNXZ1hTMcco-oEA_6VP377URbD_5K8vTWRMAv-Pd2tdE0MmtW67rU8oWWviPjKZ3WBfHgwOU57Bs1L1CEmy8DpxRv5TFIUi4i-htiT9oWtrnft6Ut-UQU7bgIhbPvFbe1lrgk_nupT8Oq3j3UltL4aBFI_NdYPyKwW4T7Ot2KS-OCcO3Te7bq9tGy6Ry_ipZ571uET4OcYZV5d5Ym5NStYDki-7AN3fxJjbd0Adm1nWYNdYJFGxg3hMdXJzV-8bkyGjnX43sYQ== HTTP/2.0
x-sky-signature: SkyOTT client="NBCU-WEB-v6",signature="98Iw36ODpVPDS6Ial//I2/eGsrA=",timestamp="1650811663",version="1.0"
~~~

Now go to Get Widevine Keys [2], and enter the information from above:

~~~
PSSH:
AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAWylyu1Ra6Obew32huzP9I49yVmwY=

License:
https://ovp.peacocktv.com/drm/widevine/acquirelicense?bt=43-0Bzfu3_lu2EG4AlSU4eQZPavwvjM-lP5h7cwPYI5YHEHWYzH8lqxPDqB6SJMtSwbi0_HDKaAsOjzDBwMLIvsYBKmwLhrhdpcQ_MJZXpSIR5e52sohHDwx9VyX3rtJlv8X3vOB6Fkn77yMYTzF5R2YnvXanei895p9hf9nb8hrBKY7DWMNSx03Qy7NqAKrgwZzRc00_RoolxslOVKZ2yWhvPUhCOECwnxEwHa07zNGfOBT6znd6v_gjyYs2s3YdTjK8URKKMHl8P7esyxB5Bwl6ln0svU55jTYs4V81FxUbMfjRjn49isEWBJaCkwnd4sorzkazTiXAN2g4HxexXuwISsZ3CWbOJM5MzFHGdZ8jh7ox8ZOCTvS0VPPQ==

Headers:
x-sky-signature: SkyOTT client="NBCU-WEB-v6",signature="IPhXJwk6nYiwXJuo7S3aK92VLOc=",timestamp="1650813143",version="1.0"
~~~

You should get a result like this:

~~~
883780c4a981494cb9949e93792ff8a7:680a46ebd6cf2b9a6a0b05a24dcf944a
~~~

Finally, you can decrypt [3] the media:

~~~
mp4decrypt --key 883780c4a981494cb9949e93792ff8a7:680a46ebd6cf... enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://getwvkeys.cc
3. https://bento4.com/downloads
