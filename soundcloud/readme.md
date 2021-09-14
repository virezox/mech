# SoundCloud

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
