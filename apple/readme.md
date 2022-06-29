# Apple TV+

- https://github.com/TDenisM/APPLE-TV-4K-Downloader/blob/main/appletv.py
- https://github.com/edgeware/mp4ff/issues/150
- https://github.com/ytdl-org/youtube-dl/issues/30808

Workaround:

~~~
packager-win-x64 --enable_raw_key_decryption `
--keys key_id=00000000000000000000000000000000:key=22bdb0063805260307ee5045c0f3835a `
stream=video,in=enc.mp4,output=dec.mp4
~~~

https://github.com/shaka-project/shaka-packager/releases
