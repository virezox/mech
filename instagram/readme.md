# Instagram

One response looks like this:

~~~
items[0].image_versions2.candidates[0].url
Item
~~~

the other looks like this:

~~~
graphql.shortcode_media.display_url
Media
~~~

## Android client

<https://github.com/itsMoji/Instagram_SSL_Pinning>

## How to get User-Agent?

https://github.com/89z/googleplay

## Why does this exist?

January 28 2022.

I use it myself.

https://instagram.com/p/CT-cnxGhvvO

## Why not use other APIs?

`/api/v1/media/` uses ID instead of shortcode.

`/embed` API does not return URLs in all cases:

<https://instagram.com/p/CY-Wwq_O6S0/embed>

`/graphql/query/` API gets images up to 1080p, but not original quality.
