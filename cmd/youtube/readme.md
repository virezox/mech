# YouTube

## Age gate

Video that can be embedded:

https://www.youtube.com/embed/HtVdAasjOgU

Other age gate video exist, that cannot be embedded:

https://www.youtube.com/embed/bO7PgQ-DtZk

## Sort

If we look at video `youtube.com/watch?v=3gdfNdilGFE`:

~~~
itag 302, height 720, 3.5 mb/s, 139.4 MB, video/webm; codecs="vp9"
~~~

If we look at video `youtube.com/watch?v=Mfu_iFS-UY8`:

~~~
itag 247, height 720, 833.4 kb/s, 9.6 MB, video/webm; codecs="vp9"
~~~

If we look at video `youtube.com/watch?v=XeojXq6ySs4`:

~~~
itag 247, height 720, 127.0 kb/s, 5.9 MB, video/webm; codecs="vp9"
~~~

Can we get `itag 247` on all videos, with any client? No. Can we get `itag 302`
on all videos, with any client? No. For any video, can we get
`dashManifestUrl`, with any client? No.
