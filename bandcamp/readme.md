# Bandcamp

~~~
Date: Fri, 17 Sep 2021 21:31:29 +0000
From: Bandcamp Support <support@bandcamp.com>
Subject: Re: (Case 1268890) API Access

Thanks for your interest. Our APIs are currently limited to Label Accounts
(bandcamp.com/labels) looking to pull physical order and general sales
information, though that might change down the road. Please keep an eye on
bandcamp.com/developer for updates.
~~~

## Resolve

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

https://github.com/masterT/bandcamp-scraper/issues/59
