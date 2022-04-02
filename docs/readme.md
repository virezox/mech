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
11.542 B true YouTube
4.100 B true Instagram
1.305 B false Spotify: Music and Podcasts
1.270 B true Twitter
340.410 M true SoundCloud: Play Music & Songs
317.269 M false Deezer: Music & Podcast Player
254.577 M false Amazon Music: Discover Songs
94.988 M false Reddit
85.098 M false iHeart: Music, Radio, Podcasts
84.538 M false Apple Music
31.661 M true Vimeo
25.059 M false TIDAL Music
22.338 M true Paramount+ | Peak Streaming
21.309 M false Napster Music
13.856 M true The NBC App - Stream TV Shows
12.728 M true ABC – Live TV & Full Episodes
2.656 M true Bandcamp
2.262 M true PBS Video: Live TV & On Demand
695.926 K false Qobuz
~~~

Amazon Music:

First 30 seconds is normal, rest of track is encrypted.

Apple Podcasts:

https://github.com/89z/mech/tree/12876427d9a0629de6131734c3e778daf1a62b7a

BBC:

https://github.com/89z/mech/tree/43f0466455e0204ec3319900def5e1ae638aaa8b

Bleep:

https://github.com/89z/mech/tree/f2e43b8d5557c732c7deaee78a815e6fb22378b5

CWTV:

https://github.com/89z/mech/tree/fdf4b6c0f4015d5ef2de6dcbba25bd1ac296c743

Hulu:

- https://archive.ph/CzT12
- https://github.com/chris124567/hulu/issues/14

IMDb:

https://github.com/89z/mech/tree/134506e3245b8dd2541dfae40757645a057d02a2

MTV:

https://github.com/89z/mech/tree/e90712f2f47ad49d50aaa1b32f83a1fbd541adfa

Pandora:

https://github.com/89z/mech/tree/15253dcf97be8d21cd6eccb3bfc8aac7a736725e

Reddit:

https://github.com/89z/mech/tree/b901af458f09e04b662dbeabc8f5880527f9adfa

Spotify:

Audio is encrypted.

TED:

https://github.com/89z/mech/tree/8ddaa050fbffe14a12166dae83f8f53f52de2480

TikTok:

https://github.com/89z/mech/tree/71948e302088290408a68734c641bf7dcab0b590

Tumblr:

https://github.com/89z/mech/tree/b133200505b6767f65654574fd98ab905bec2fba
