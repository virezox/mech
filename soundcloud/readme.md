# SoundCloud

## Android client

~~~
com.soundcloud.android
~~~

Newer versions are broken. Even with no proxy, I just get this over and over:

~~~
Signing you in
~~~

This one seems to work:

~~~
2020.03.19-release
~~~

API 24? Small fail. Big pass.

## How to get `client_id`

First, make a request like this:

~~~
GET / HTTP/2
Host: m.soundcloud.com
~~~

In the HTML response, you should see something like this:

~~~
"clientId":"iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
~~~

You can also get it with JADX, but more difficult:

~~~
com\soundcloud\android\api\di\a.java
return new kt.a(cVar, "dbdsA8b6V6Lw7wzu1x0T4CLxt58yd4Bf", iVar.deobfuscateString("NykCWyEEEyUrRCd2AQAtEAUdfy9HKAAkKRwjJh4cMSk="));
~~~

The `client_id` seems to last at least a year:

https://github.com/rrosajp/soundcloud-archive/commit/c02809dc

## Image

artworks:

~~~
https://soundcloud.com/oembed?format=json&url=https://soundcloud.com/western_vinyl/jessica-risker-cut-my-hair
https://i1.sndcdn.com/artworks-000308141235-7ep8lo-t500x500.jpg
~~~

placeholder:

~~~
https://soundcloud.com/oembed?format=json&url=https://soundcloud.com/pdis_inpartmaint/harold-budd-perhaps-moss
https://soundcloud.com/images/fb_placeholder.png
~~~

avatars:

~~~
https://soundcloud.com/oembed?format=json&url=https://soundcloud.com/pdis_inpartmaint
https://i1.sndcdn.com/avatars-000274827119-0dxutu-t500x500.jpg
~~~

## Why does this exist?

January 28 2022.

I use the site myself.

https://soundcloud.com/afterhour-sounds/premiere-ele-bisu-caradamom-coffee
