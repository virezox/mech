# Bandcamp

For tracks and albums, you can make a request like this:

~~~
HEAD /track/amaris-2 HTTP/1.1
Host: schnaussandmunk.bandcamp.com
~~~

and in the response should be this:

~~~
Set-Cookie: session=1 r:["nilZ0t2809477874x1633448962"]	t:1633448962
~~~

In this case, `t` is the `tralbum_type` and `2809477874` is the `tralbum_id`.
You can then make a request like this:

~~~
GET /api/mobile/24/tralbum_details?band_id=1&tralbum_type=t&tralbum_id=2809477874 HTTP/1.1
Host: bandcamp.com
~~~

For bands, you can make a request like this:

~~~
HEAD /music HTTP/1.1
Host: schnaussandmunk.bandcamp.com
~~~

## APK

~~~
com.bandcamp.android
~~~

https://github.com/httptoolkit/frida-android-unpinning
