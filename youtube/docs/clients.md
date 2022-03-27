# Clients

## ANDROID

~~~
googleplay -a com.google.android.youtube
~~~

https://github.com/89z/googleplay

I was able to get YouTube working with Android API 24, using this method [1].
Doing that, I was able to test non-stock versions of YouTube. It seems anything
starting with version 14.01.51 (2019) is doing this HTTP/3 or whatever voodoo.
Anything older wont work at all, even without proxy, failing with error 400.

1.  https://android.stackexchange.com/a/245551

## ANDROID\_EMBEDDED\_PLAYER

~~~
googleplay -a com.google.android.youtube
~~~

https://github.com/89z/googleplay

## MWEB

Needed to get `publishDate`:

~~~
aN76CmldknI
~~~

## WEB\_CREATOR

~~~
GET /?approve_browser_access=true HTTP/1.1
Host: studio.youtube.com
Authorization: Bearer ya29.a0ARrdaM-2nXUrxlFNOx3hZAUNICfCwmhHKHenQkebpQFGNoYdE...
~~~
