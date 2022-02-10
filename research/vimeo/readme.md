# Vimeo

## Pass pass

~~~
GET /videos?links=https://vimeo.com/581039021/9603038895 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMjAsInVzZXJfaWQiOm51bGwsImFwcF9pZCI6NTg0NzksInNjb3BlcyI6InB1YmxpYyIsInRlYW1fdXNlcl9pZCI6bnVsbH0.mtroxz3fs6nb7wx50JnmwM2Nq7Hd-A2d5pDysuNMLZ0
~~~

https://vimeo.com/api/oembed.json?url=https://vimeo.com/581039021/9603038895

## Pass fail

~~~
GET /videos?links=https://vimeo.com/477957994/2282452868 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMjAsInVzZXJfaWQiOm51bGwsImFwcF9pZCI6NTg0NzksInNjb3BlcyI6InB1YmxpYyIsInRlYW1fdXNlcl9pZCI6bnVsbH0.mtroxz3fs6nb7wx50JnmwM2Nq7Hd-A2d5pDysuNMLZ0
~~~

https://vimeo.com/api/oembed.json?url=https://vimeo.com/477957994/2282452868

## Fail pass

~~~
GET /videos?links=https://vimeo.com/581039021?unlisted_hash=9603038895 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMjAsInVzZXJfaWQiOm51bGwsImFwcF9pZCI6NTg0NzksInNjb3BlcyI6InB1YmxpYyIsInRlYW1fdXNlcl9pZCI6bnVsbH0.mtroxz3fs6nb7wx50JnmwM2Nq7Hd-A2d5pDysuNMLZ0
~~~

<https://vimeo.com/api/oembed.json?url=https://vimeo.com/581039021?unlisted_hash=9603038895>

## Fail fail

~~~
GET /videos?links=https://vimeo.com/477957994?unlisted_hash=2282452868 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMjAsInVzZXJfaWQiOm51bGwsImFwcF9pZCI6NTg0NzksInNjb3BlcyI6InB1YmxpYyIsInRlYW1fdXNlcl9pZCI6bnVsbH0.mtroxz3fs6nb7wx50JnmwM2Nq7Hd-A2d5pDysuNMLZ0
~~~

<https://vimeo.com/api/oembed.json?url=https://vimeo.com/477957994?unlisted_hash=2282452868>

## How to get Basic Authentication?

https://github.com/httptoolkit/frida-android-unpinning

## Issue

https://github.com/ytdl-org/youtube-dl/issues/30622
