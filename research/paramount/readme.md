# Paramount+

- https://github.com/ytdl-org/youtube-dl/issues/30491
- https://play.google.com/store/apps/details?id=com.cbs.app

Should work without Frida. Install system certificate, then might have to try a
couple of times.

## guid

Looks like this:

<https://paramountplus.com/shows/the-harper-house/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/the-harper-house-the-harper-house>

## cmsAccountId

Looks like this:

~~~
dJ5BDC
~~~

can get it like this:

~~~
paramountplus.com/apps-api/v2.0/androidphone/video/cid/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json?
at=ABAJi4xSDPXIEUKTlJ6BFQpMdL3hrvn5xbm%2BXly%2B9QZJFycgSL%2F4%2FYiDMKY4XWomRkI
~~~

## pid

Looks like this:

~~~
fNsRH_fjko5T
~~~

can get it with these:

~~~
paramountplus.com/apps-api/v2.0/androidphone/video/cid/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json?
at=ABAJi4xSDPXIEUKTlJ6BFQpMdL3hrvn5xbm%2BXly%2B9QZJFycgSL%2F4%2FYiDMKY4XWomRkI

link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU?format=preview
~~~

## aid

Looks like this:

~~~
2198311517
~~~

can get it like this:

~~~
link.theplatform.com/s/dJ5BDC/fNsRH_fjko5T?format=preview
~~~

## M3U

can get it like this:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU?formats=M3U

link.theplatform.com/s/dJ5BDC/fNsRH_fjko5T?formats=M3U

SKD:
paramountplus.com/apps-api/v2.0/iphone/video/cid/
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json?
at=ABAJi4xSDPXIEUKTlJ6BFQpMdL3hrvn5xbm%2BXly%2B9QZJFycgSL%2F4%2FYiDMKY4XWomRkI
~~~
