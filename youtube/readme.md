# YouTube

Is `maxres1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/maxres1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/maxres1.jpg

Is `sd1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/sd1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/sd1.jpg

If `hq1` always available? Yes:

http://i.ytimg.com/vi/hq2KgzKETBw/hq1.jpg

## Clients

Need Android:

client  | MeJVWBSsPAY
--------|-------
Android | pass
Embed   | fail
Mweb    | fail

Need Embed:

client  | QWlNyzzwgcc
--------|-------
Android | fail
Embed   | pass
Mweb    | fail

Need Mweb:

client  | aN76CmldknI publishDate
--------|-------
Android | fail
Embed   | fail
Mweb    | pass

## Cover art

- <https://wiki.musicbrainz.org/Cover_Art_Archive/API>
- http://hackerfactor.com/blog?archives/529-Kind-of-Like-That.html
- http://mdpi.com/2313-433X/7/3/48/htm

## Device OAuth

- https://datatracker.ietf.org/doc/html/rfc8628
- https://developers.google.com/identity/sign-in/devices

## How to get X-Goog-Api-Key

Make a request like this:

~~~
GET / HTTP/2
Host: www.youtube.com
~~~

in the response you should see something like this:

~~~
"INNERTUBE_API_KEY":"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
~~~

https://cloud.google.com/apis/docs/system-parameters

## How to get `client_id` and `client_secret`

Set User-Agent to [1]:

~~~
Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version
~~~

Then visit:

https://www.youtube.com/tv

Next open your browser menu, and click Web Developer, Network or similar. Then
go back to the page, and click "Sign in", then "Sign in with a web browser". On
the next page, dont bother with any of the instructions, just go back to
Developer Tools, and after about five seconds you should see a JSON request like
this:

~~~
POST /o/oauth2/token HTTP/1.1
Host: www.youtube.com

{"client_id":"861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com",
"client_secret":"SboVhoG9s0rNafixCSGGKXAT",
"code":"AH-1Ng14qVvEj76OeM_h14Mgklgyhchbyc67MhULhCKPY6K-0DTYJqaKng2ULVFTmTzU...",
"grant_type":"http://oauth.net/grant_type/device/1.0"}
~~~

References:

1. <https://github.com/youtube/cobalt/blob/master/src/cobalt/browser/user_agent_string.cc>

## other repos

- <https://github.com/Hexer10/youtube_explode_dart/issues>
- https://github.com/Athlon1600/youtube-downloader/issues
- https://github.com/Tyrrrz/YoutubeDownloader/issues
- https://github.com/Tyrrrz/YoutubeExplode/issues
- https://github.com/iawia002/annie/issues
- https://github.com/kkdai/youtube/issues
- https://github.com/yt-dlp/yt-dlp/issues
- https://github.com/ytdl-org/youtube-dl/issues
