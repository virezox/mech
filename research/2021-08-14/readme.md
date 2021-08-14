# August 14 2021

First make request 0:

~~~
GET /api/reference/videos HTTP/1.1
Host: developer.vimeo.com
~~~

then in response:

~~~
"x-playground-token":"81f5441838f52d92c8b5de970abf819fbfc843fb38c6337a5119958..."
~~~

Then make request 1, then in response:

~~~
"Authorization":"jwt eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2Mjg5NjU..."
~~~

Then make request 2.

https://developer.vimeo.com/api/reference/videos
