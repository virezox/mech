# ABC

## Locked videos

With locked videos, it seems you can only FairPlay.

## Android client

~~~
com.disney.datg.videoplatforms.android.abc
~~~

Install system certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://abc.com/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you
~~~

## Authorize

This URL:

http://api.entitlement.watchabc.go.com/vp2/ws-secure/entitlement/2020/authorize.json

Comes from here:

<https://api.config.watchabc.go.com/appsconfig/prod/abcv3/031_04/10.23.0/config.json>

## Why does this exist?

June 2 2022

https://github.com/89z/mech/issues/58
