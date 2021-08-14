# SoundCloud

## How to get `client_id`

First, make a request like this:

~~~
GET / HTTP/2
Host: soundcloud.com
~~~

In the HTML response, you should see several JavaScript assets, like this:

~~~html
<script crossorigin src="https://a-v2.sndcdn.com/assets/2-b0e52b4d.js"></script>
<script crossorigin src="https://a-v2.sndcdn.com/assets/49-4b976e4f.js"></script>
~~~

If you download asset 2, you should see this in the response:

~~~
?client_id=fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj&
~~~

If you download asset 49, you should see this in the response:

~~~
client_id:"fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj"
~~~

From my testing, asset 2 was smaller. The `client_id` seems to last at least a
year:

https://github.com/flyingrub/scdl/commit/08317287
