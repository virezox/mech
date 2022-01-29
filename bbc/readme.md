# BBC

## Android client

~~~
bbc.mobile.news.ww
~~~

https://github.com/ytdl-org/youtube-dl/issues/30565

Install system cert, then visit this page:

https://github.com/89z/mech/issues/15

Click this link:

https://bbc.co.uk/news/av/uk-politics-60159228

and you should see a request like this:

~~~
GET https://walter-resolver-cdn.api.bbci.co.uk/resolve?uri=https://www.bbc.co.uk/news/av/uk-politics-60159228 HTTP/2.0
accept: application/json
user-agent: BBCNews/5.19.0 (Android SDK built for x86; Android 7.0)
accept-encoding: gzip
content-length: 0
~~~

## Why does this exist?

January 28 2022:

https://github.com/ytdl-org/youtube-dl/issues/30565
