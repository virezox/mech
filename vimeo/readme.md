# Vimeo

## Android client

- https://github.com/httptoolkit/frida-android-unpinning
- https://play.google.com/store/apps/details?id=com.vimeo.android.videoapp
- https://play.google.com/store/apps/details?id=tv.vhx.moonflix

Requires SDK version 26

## Parse URL

We have to write our own URL parser, since the current ones dont work in all
cases. This URL fails with first parser:

~~~
GET /api/oembed.json?url=https://vimeo.com/477957994?unlisted_hash=2282452868
Host: vimeo.com
~~~

and second parser:

~~~
GET /videos?links=https://vimeo.com/477957994?unlisted_hash=2282452868 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMj...
~~~

## files

If we use the Vimeo URL:

https://player.vimeo.com/video/762862842/config

we get direct links to videos:

https://vod-progressive.akamaized.net/exp=1659985751~acl=%2Fvimeo-prod-skyfir...

If we use the VHX URL:

https://api.vhx.tv/videos/17901

we then have to request again:

https://api.vhx.tv/videos/17901/files

before getting the direct link:

https://gcs-vimeo.akamaized.net/exp=1659999606~acl=%2A%2F1253374903.mp4%2A~hm...
