# Vimeo

If we use the VHX API, we get these videos:

~~~
https://api.vhx.tv/videos/28599/files?codec=h264&format=mp4&quality=360p
https://api.vhx.tv/videos/28599/files?codec=h264&format=mp4&quality=540p
https://api.vhx.tv/videos/28599/files?codec=h264&format=mp4&quality=720p
https://api.vhx.tv/videos/28599/files?codec=h264&format=m3u8&quality=adaptive
https://api.vhx.tv/videos/28599/files?codec=h264&format=mpd&quality=adaptive
~~~

Is 720p the highest?

## HTML

pass:

~~~
GET /videos/1264265?vimeo=1 HTTP/1.1
Host: embed.vhx.tv
user-agent: curl/7.78.0
accept: */*

GET /videos/1264265 HTTP/1.1
Host: api.vhx.tv
authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E...

HTTP/2.0 200 OK
~~~

JWT fails:

~~~
GET /_next/jwt HTTP/1.1
Host: vimeo.com
X-Requested-With: XMLHttpRequest

GET /videos/1264265 HTTP/1.1
Host: api.vhx.tv
authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NTk4MjY5O...

HTTP/2.0 401 Unauthorized
~~~

api.vimeo.com fails:

~~~
GET /_next/jwt HTTP/1.1
Host: vimeo.com
X-Requested-With: XMLHttpRequest

GET /videos/1264265?fields=duration,download,name,pictures,release_time,user HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NTk4MjcxN...

HTTP/1.1 404 Not Found
~~~

read write fails:

~~~
POST https://api.vhx.tv/oauth/token/ HTTP/2.0
accept: application/json
x-ott-agent: android site/40903 android-app/7.206.1
user-agent: Moonflix/7.206.2(Google Android SDK built for x86, Android 7.0 (API 24))
ott-client-version: 7.206.1
content-type: application/json
accept-encoding: identity

{
  "client_id": "85c89e1f5a386b54dbad29a60be04d64b45dfeb6fe710f408e55b0c6f1dedddc",
  "client_secret": "efec91dc214518ff812c46405e1025c3d0259defa6969313115a33effe716fd7",
  "grant_type": "client_credentials",
  "scope": "read write"
}

GET https://api.vhx.tv/videos/1264265 HTTP/2.0
user-agent: Moonflix/7.206.2(Google Android SDK built for x86, Android 7.0 (API 24))
authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E...
x-ott-agent: android site/40903 android-app/7.206.1
accept-encoding: identity
content-length: 0

HTTP/2.0 403 Forbidden
~~~

public fails:

~~~
POST https://api.vhx.tv/oauth/token/ HTTP/2.0
accept: application/json
x-ott-agent: android site/40903 android-app/7.206.1
user-agent: Moonflix/7.206.2(Google Android SDK built for x86, Android 7.0 (API 24))
ott-client-version: 7.206.1
content-type: application/json
accept-encoding: identity

{
  "client_id": "85c89e1f5a386b54dbad29a60be04d64b45dfeb6fe710f408e55b0c6f1dedddc",
  "client_secret": "efec91dc214518ff812c46405e1025c3d0259defa6969313115a33effe716fd7",
  "grant_type": "client_credentials",
  "scope": "public"
}

GET https://api.vhx.tv/videos/1264265 HTTP/2.0
user-agent: Moonflix/7.206.2(Google Android SDK built for x86, Android 7.0 (API 24))
authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E...
x-ott-agent: android site/40903 android-app/7.206.1
accept-encoding: identity
content-length: 0

HTTP/2.0 403 Forbidden
~~~

- https://developer.vimeo.com/api/authentication
- https://play.google.com/store/apps/details?id=tv.vhx.moonflix
