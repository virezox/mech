# September 22 2021

This method is severely rate limited:

~~~
GET /tv/CT-cnxGhvvO/?__a=1 HTTP/1.1
Host: www.instagram.com
User-Agent: Mozilla
~~~

This requires authentication:

~~~
GET /api/v1/media/2665693907534674894/info/ HTTP/1.1
Host: i.instagram.com
~~~

This works, although it returns HTML, so not ideal:

~~~
GET /p/CT-cnxGhvvO/embed/ HTTP/1.1
Host: www.instagram.com
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36
~~~
