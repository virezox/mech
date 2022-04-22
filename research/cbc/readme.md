# CBC

- https://gem.cbc.ca/media/downton-abbey/s01e05
- https://github.com/ytdl-org/youtube-dl/issues/30043

## Web client

The web client will not allow you to create an account:

> Accounts cannot be created outside of Canada.

https://gem.cbc.ca/join-now

## Android client

Android client allows you to create an account with no problems.

https://play.google.com/store/apps/details?id=ca.cbc.android.cbctv

Install system certificate. Random IPV4:

- <https://github.com/firehol/blocklist-ipsets/blob/master/geolite2_country/country_ca.netset>
- <https://github.com/ytdl-org/youtube-dl/blob/a0068bd6/youtube_dl/utils.py#L5373-L5384>

~~~
GET /media/validation/v2?appCode=gem&idMedia=929078&manifestType=mobile&output=json&tech=hls HTTP/1.1
Host: services.radio-canada.ca
X-Forwarded-For: 99.246.97.250
x-claims-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiV...
~~~

How to get media ID?
