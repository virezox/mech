# Docs

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. `--proxy 127.0.0.1:8080 --no-check-certificate`
2. try navigating to the target page from home screen, instead of going directly
   to page
3. check media page for JSON requests
4. MITM Proxy Firefox
5. MITM Proxy APK
6. JADX

## Sites

~~~
11.552 B true YouTube
4.103 B true Instagram
1.306 B false Spotify: Music and Podcasts
1.271 B true Twitter
340.504 M true SoundCloud: Play Music & Songs
317.349 M false Deezer: Music & Podcast Player
254.711 M false Amazon Music: Discover Songs
95.262 M false Reddit
85.114 M false iHeart: Music, Radio, Podcasts
84.655 M false Apple Music
31.678 M true Vimeo
25.069 M false TIDAL Music
22.381 M true Paramount+ | Peak Streaming
21.310 M false Napster Music
13.859 M true The NBC App - Stream TV Shows
12.731 M true ABC – Live TV & Full Episodes
2.659 M true Bandcamp
2.266 M true PBS Video: Live TV & On Demand
696.441 K false Qobuz
~~~

Amazon Prime Video:

https://github.com/ytdl-org/youtube-dl/issues/1753

Apple TV:

https://github.com/ytdl-org/youtube-dl/issues/30808
