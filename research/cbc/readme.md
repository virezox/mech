# CBC

- https://gem.cbc.ca/media/downton-abbey/s01e05
- https://github.com/ytdl-org/youtube-dl/issues/30043

> The link doesn't work in the UK

The web client [1] wont even allow you to create an account:

> Accounts cannot be created outside of Canada.

but the Android client allows you to create an account with no problems.
Further, I found the media request works, as long as you add
`--geo-bypass-country CA` (X-Forwarded-For):

~~~
GET /media/validation/v2?appCode=gem&idMedia=929078&manifestType=mobile&output=json&tech=hls HTTP/1.1
Host: services.radio-canada.ca
X-Forwarded-For: 99.246.97.250
x-claims-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiV...
~~~

1. https://gem.cbc.ca/join-now

Random IPV4:

- <https://github.com/firehol/blocklist-ipsets/blob/master/geolite2_country/country_ca.netset>
- <https://github.com/ytdl-org/youtube-dl/blob/a0068bd6/youtube_dl/utils.py#L5373-L5384>

Android:

https://play.google.com/store/apps/details?id=ca.cbc.android.cbctv

Install system certificate. Get `claimsToken` like this:

~~~
GET /ott/cbc-api/v2/profile HTTP/1.1
Host: services.radio-canada.ca
ott-access-token: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjkzQURGMUNFNDhG...
~~~
