# October 5 2021

## new

track:

~~~
PS C:\> curl -I https://schnaussandmunk.bandcamp.com/track/amaris-2
HTTP/1.1 200 OK
Set-Cookie: session=1	r:["nilZ0t2809477874x1633448962"]	t:1633448962; domain=.bandcamp.com; path=/; expires=Fri, 19 Nov 2021 15:49:22 -0000
~~~

album:

~~~
PS C:\> curl -I https://schnaussandmunk.bandcamp.com/album/passage-2
HTTP/1.1 200 OK
Set-Cookie: session=1	r:["nilZ0a1670971920x1633449063"]	t:1633449063; domain=.bandcamp.com; path=/; expires=Fri, 19 Nov 2021 15:51:03 -0000
~~~

## old

track:

~~~
bandcamp.com/api/url/2/info?key=veidihundr&url=schnaussandmunk.bandcamp.com/track/amaris-2
{"track_id":2809477874,"band_id":3454424886}
~~~

album:

~~~
bandcamp.com/api/url/2/info?key=veidihundr&url=schnaussandmunk.bandcamp.com/album/passage-2
{"album_id":1670971920,"band_id":3454424886}
~~~
