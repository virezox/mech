# SoundCloud

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

The `client_id` seems to last at least a year:

https://github.com/rrosajp/soundcloud-archive/commit/c02809dc

## Other `client_id`s

~~~
mobi_browser
https://m.soundcloud.com/

pc_browser
https://soundcloud.com/

widget & widget2
https://w.soundcloud.com/player/

win10_app
https://soundcouch.soundcloud.com/#/

win10_app_beta
https://soundcouch-beta.soundcloud.com/#/
~~~

- https://archive.ph/IOglb
- https://github.com/inkuringu-ika/inkuringu-ika.github.io/issues/1
