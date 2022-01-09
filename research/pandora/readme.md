# Pandora

Using this example:

https://pandora.com/artist/the-black-dog/radio-scarecrow/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV

How do we get from this:

~~~
TRddpp5JJ2hqnVV
~~~

to this:

~~~
TR:1168891
~~~

Base64 encode `1168891` gives `MTE2ODg5MQ`. Base64 encode `S1168891` gives
`UzExNjg4OTE`. Base64 encode `TR:1168891` gives `VFI6MTE2ODg5MQ`. Searching with
JaDx for these:

~~~
DetailUrl
SeoToken
~~~

Doesnt return anything helpful. Turning to MITM Proxy, some of the request
bodies are encrypted, so we should focus on the responses. Response header
search returns nothing. Response body search fails as well. Request URL search
fails as well. Request header search fails as well. Request body search fails as
well. Next turning to the web app. Request search fails. Response search fails
as well. So we will need to parse HTML.
