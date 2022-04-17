# CBC

- https://gem.cbc.ca/media/downton-abbey/s01e05
- https://github.com/ytdl-org/youtube-dl/issues/30043

> The link doesn't work in the UK

The web client [1] wont even allow you to create an account:

> Accounts cannot be created outside of Canada.

but the Android client [2] allows you to create an account with no problems.
Further, I found the media request works, as long as you add
`--geo-bypass-country CA` (X-Forwarded-For):

~~~
GET /media/validation/v2?appCode=gem&idMedia=929078&manifestType=mobile&output=json&tech=hls HTTP/1.1
Host: services.radio-canada.ca
x-claims-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiV...
X-Forwarded-For: 99.246.97.250
~~~

1. https://gem.cbc.ca/join-now
2. https://play.google.com/store/apps/details?id=ca.cbc.android.cbctv
