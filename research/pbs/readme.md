# PBS

https://github.com/yt-dlp/yt-dlp/issues/3092

This is it:

~~~
GET /videos/sanditon/0bb6c46b-c0da-4b5d-8fce-169b5ee9bd75/2000285295/hd-16x9-mezzanine-1080p/mast5216-hls-16x9-1080p-270p-365k_00010.ts HTTP/1.1
User-Agent: PbsPlayer/5.5.5 (Linux;Android 7.0) ExoPlayerLib/2.10.6
Accept-Encoding: identity
Host: ga.video.cdn.pbs.org
Connection: Keep-Alive
content-length: 0
~~~

comes from:

~~~
GET /videos/sanditon/0bb6c46b-c0da-4b5d-8fce-169b5ee9bd75/2000285295/hd-16x9-mezzanine-1080p/mast5216-hls-16x9-1080p-270p-365k.m3u8 HTTP/1.1
User-Agent: PbsPlayer/5.5.5 (Linux;Android 7.0) ExoPlayerLib/2.10.6
Host: ga.video.cdn.pbs.org
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

comes from:

~~~
GET /videos/sanditon/0bb6c46b-c0da-4b5d-8fce-169b5ee9bd75/2000285295/hd-16x9-mezzanine-1080p/mast5216-hls-16x9-1080p_830.m3u8 HTTP/1.1
User-Agent: PbsPlayer/5.5.5 (Linux;Android 7.0) ExoPlayerLib/2.10.6
Host: ga.video.cdn.pbs.org
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

comes from:

~~~
GET /redirect/1c3c2955a0254cee8d30490f4b5db148/ HTTP/1.1
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Host: urs.pbs.org
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

comes from:

~~~
GET /v3/android/screens/video-assets/episode-1-eyn9m9/?station_id=b3291387-78a4-41e1-beb0-da2f61a96a3e HTTP/1.1
X-PBS-PlatformVersion: 5.5.5
Authorization: Basic YW5kcm9pZDpiYVhFN2h1bXVWYXQ=
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Host: content.services.pbs.org
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~
