# Peacock

- <https://peacocktv.com/watch/playback/vod/GMO_00000000441654_01>
- https://github.com/ytdl-org/youtube-dl/issues/26018

## Native Libraries

<https://ragingrock.com/AndroidAppRE/reversing_native_libs>

## Android client

download:

https://play.google.com/store/apps/details?id=com.peacocktv.peacockandroid

Then create Virtual Device using API 23 or higher. Install user certificate. It
seems some request dont get captured, but from my testing that is also true when
using system certificate and even Frida. If you start the app and Sign In, this
request:

~~~
POST https://rango.id.peacocktv.com/signin/service/international HTTP/2.0
content-type: application/x-www-form-urlencoded
x-skyott-device: MOBILE
x-skyott-proposition: NBCUOTT
x-skyott-provider: NBCU
x-skyott-territory: US

userIdentifier=MY_EMAIL&password=MY_PASSWORD
~~~

will fail:

~~~
HTTP/2.0 429
~~~

You can fix this problem by removing this request header before starting the
app:

~~~
set modify_headers '/~u signin.service.international/x-skyott-device/'
~~~

Header needs to be removed from that request only, since other requests need the
header. Now start the app and Sign In:

~~~
Email:
peacocktv2@proton.me

Password:
rWRci6atDD#j3r

Email:
mediadownloaddummy@gmail.com

Password:
!pCoCkduMMyT35tr!
~~~

Under Who's watching, click Account Holder.
