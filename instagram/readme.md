# Instagram

## How to get User-Agent?

https://apkpure.com/instagram/com.instagram.android

## How to get `query_hash`?

If you make a request like this:

~~~
GET / HTTP/1.1
Host: www.instagram.com
User-Agent: Mozilla
~~~

in the response body, you should see something like this:

~~~html
<link rel="preload" href="/static/bundles/metro/Consumer.js/c705f96d9483.js"
as="script" type="text/javascript" crossorigin="anonymous" />
~~~

If you request that link:

~~~
GET /static/bundles/metro/Consumer.js/c705f96d9483.js HTTP/1.1
Host: www.instagram.com
~~~

in the response body, you should see something like this:

~~~js
var
   u = "2efa04f61586458cef44441f474eee7c",
   l = "6ff3f5c474a240353993056428fb851e",
   E = "ba5c3def9f75f43213da3d428f4c783a",
   p = 40,
   v = 24,
   f = 3;
~~~

The first variable should be the target `query_hash`.
