# Docs

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. `yt-dlp --proxy 127.0.0.1:8080 --no-check-certificate`
2. try navigating to the target page from home screen, instead of going directly
   to page
3. check media page for JSON requests
4. check HAR file
5. check HTML
6. check JavaScript
7. MITM Proxy APK for JSON requests
8. JADX

## Proxy

https://freeproxylists.net/br.html

## Sites

~~~
11.267 B true YouTube
4.016 B true Instagram
1.994 B true TikTok
1.259 B false Spotify: Music and Podcasts
1.250 B true Twitter
336.952 M true SoundCloud: Play Music & Songs
314.864 M false Deezer: Music & Podcast Player
277.867 M true Pandora - Music & Podcasts
250.932 M false Amazon Music: Discover Songs
166.553 M true IMDb: Your guide to movies, TV shows, celebrities
132.387 M true Tumblr – Culture, Art, Chaos
90.706 M false Reddit
84.587 M false iHeart: Radio, Music, Podcasts
31.199 M true Vimeo
29.951 M true BBC News
29.594 M true TED
21.702 M false Paramount+ | Peak Streaming
21.295 M false Napster Music
13.638 M true The NBC App - Stream TV Shows
2.601 M true Bandcamp
2.198 M false PBS Video: Live TV & On Demand
681.410 K false Qobuz
~~~

Amazon Music:

First 30 seconds is normal, rest of track is encrypted.

Apple Podcasts:

https://github.com/89z/mech/tree/12876427d9a0629de6131734c3e778daf1a62b7a

PBS:

https://github.com/89z/mech/tree/c825743ab7594025b9c70632d934820e2c68d20a

Reddit:

https://github.com/89z/mech/tree/b901af458f09e04b662dbeabc8f5880527f9adfa
