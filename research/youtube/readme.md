# YouTube

~~~
youtube-dl -x -o audio.opus GhJNXlMsPG4
~~~

I dont think `--extract-audio` makes sense in this case, as it produces an
`.opus` file (which is correct). AtomicParsley works with MPEG-4 files. So you
either need to use a different YouTube-DL command, or use a different post
processing tool

~~~
ffmpeg -i ignore.opus `
-i https://i.ytimg.com/vi/GhJNXlMsPG4/mqdefault.jpg `
-c:a copy `
-c:v mjpeg `
-disposition:v attached_pic `
out.mp4

ffmpeg -i ignore.opus `
-i image.png `
-c:a copy `
-c:v png `
-disposition:v attached_pic `
1.mp4

ffmpeg -i audio.ogg `
-i Cover.jpg `
-c:a copy `
-c:v mjpeg `
-disposition:v attached_pic `
2.mp4
~~~

## done

- https://askubuntu.com
- https://superuser.com
- https://trac.ffmpeg.org/ticket/9770
- https://unix.stackexchange.com
- https://video.stackexchange.com

## to do

https://github.com/ytdl-org/youtube-dl
