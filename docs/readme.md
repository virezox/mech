# Docs

## Apps

Updated December 15 2021.

- YouTube (10.906 B)
   - com.google.android.youtube
- Instagram (3.902 B)
   - com.instagram.android
- TikTok (1.910 B)
   - com.zhiliaoapp.musically
- Twitter (1.226 B)
   - com.twitter.android
- SoundCloud (332.603 M)
   - com.soundcloud.android
- Pandora (276.415 M)
   - com.pandora.android
- Amazon Music (245.537 M)
   - com.amazon.mp3
- iHeart (83.974 M)
   - com.clearchannel.iheartradio.controller
- Vimeo (30.496 M)
   - com.vimeo.android.videoapp
- TIDAL Music (23.447 M)
   - com.aspiro.tidal
- Napster Music (21.278 M)
   - com.rhapsody
- The NBC App (13.396 M)
   - com.nbcuni.nbc
- Bandcamp (2.539 M)
   - com.bandcamp.android
- PBS Video (2.101 M)
   - com.pbs.video
- Qobuz (656.191 k)
   - com.qobuz.music

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. see how YT-DLP does it

2. try navigating to the target page from home screen, instead of going directly
   to page

3. check media page for JSON requests
4. check HAR file
5. check HTML
6. check JavaScript
7. MITM Proxy APK for JSON requests
8. APK tool
9. JADX

## How to update topics?

~~~
PUT /repos/89z/mech/topics HTTP/1.1
Host: api.github.com
Authorization: Basic ODl6OmE1NzYxMjZlNzVlZjZiY2Y5ZDljNzEyZWIyN2RmZjFmOGFhZmQ...

{"names":[
   "youtube", "instagram"
]}
~~~

## Other sites

I have implementations for these other sites, which I removed for now, as I am
busy with other stuff.

BBC:

https://github.com/89z/mech/tree/ae41cbd605d887b03876861ca2109d00be967ca3

PBS:

https://github.com/89z/mech/tree/c825743ab7594025b9c70632d934820e2c68d20a

Reddit:

https://github.com/89z/mech/tree/b901af458f09e04b662dbeabc8f5880527f9adfa
