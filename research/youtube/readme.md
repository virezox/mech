# YouTube

~~~
youtube-dl -x -o audio.opus GhJNXlMsPG4
~~~

I dont think `--extract-audio` makes sense in this case, as it produces an
`.opus` file (which is correct). AtomicParsley works with MPEG-4 files. So you
either need to use a different YouTube-DL command, or use a different post
processing tool

~~~
ffmpeg -i audio.opus `
-i https://i.ytimg.com/vi/GhJNXlMsPG4/mqdefault.jpg `
-c:a copy `
-c:v mjpeg `
-disposition:v attached_pic `
out.mp4
~~~

## done

https://trac.ffmpeg.org/ticket/9770

## to do

- https://askubuntu.com
- https://superuser.com
- https://unix.stackexchange.com
- https://video.stackexchange.com
