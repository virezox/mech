# Instagram

## Android client

~~~
com.instagram.android
~~~

<https://github.com/itsMoji/Instagram_SSL_Pinning>

## How to get User-Agent?

https://github.com/89z/googleplay

## How to get `query_hash`?

Visit this page:

https://instagram.com

Then View Page Source, you should see a link like this:

https://instagram.com/static/bundles/es6/Consumer.js/341626c79aac.js

In the JavaScript source, you should see something like this:

~~~js
const {data: o} = await r(d[7]).query(E, {
   shortcode: n,
   child_comment_count: f,
   fetch_comment_count: p,
   parent_comment_count: v,
   has_threaded_comments: !0
});
~~~

Then find the variable that corresponds to the first argument to the `query`
method (`E` in the example above):

~~~js
const E = "7d4d42b121a214d23bd43206e5142c8c",
   _ = "6ff3f5c474a240353993056428fb851e",
   u = "ba5c3def9f75f43213da3d428f4c783a",
   p = 40,
   v = 24,
   f = 3;
~~~

## Why does this exist?

January 28 2022.

I use it myself.

https://instagram.com/p/CT-cnxGhvvO

## Why not use other APIs?

`/api/v1/media` API requires separate calls for media and comments:

~~~
GET /api/v1/media/2755022163816059161/info/ HTTP/2.0
Host: i.instagram.com
user-agent: Instagram 206.1.0.34.121 Android
Authorization: Bearer IGT:2:eyJkc191c2VyX2lkIjoiNDkzNzgxNzEzMzQiLCJzZXNzaW9ua...

GET /api/v1/media/2755652849306967814/comments/ HTTP/2.0
Host: i.instagram.com
user-agent: Instagram 215.0.0.27.359 Android
Authorization: Bearer IGT:2:eyJkc191c2VyX2lkIjoiNDkzNzgxNzEzMzQiLCJzZXNzaW9ua...
~~~

`/embed` API does not return URLs in all cases:

<https://instagram.com/p/CY-Wwq_O6S0/embed>

`__a=1` API is known to send a different response depending on the country. For
example, the `dfw5` location produces this result:

~~~
graphql.shortcode_media.display_resources
~~~

while the `hel3` location produces this result:

~~~
items[0].carousel_media[0].image_versions2.candidates
~~~
