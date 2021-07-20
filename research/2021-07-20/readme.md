# July 20 2021

## deviceregistration

First make a request like this:

~~~
curl -v `
-d key=AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w `
-d rawDeviceId=1234567890123456 `
https://youtubei.googleapis.com/deviceregistration/v1/devices
~~~

Response:

~~~
{
  "key": "AP+lc7+vnzjxeNwcPRUdiO3Mw84/IqvSq+qjJvy7hKWduKU/oTyBaoCnyOO6LTE",
  "id": "AP+lc7829EbFysNk+qGtXaIl1a5ApLUBnvW2p3KVtuvsPMsosSGTNsoTWJsl+vMYy27czPr1Eu/HDs7n8wkzh1UtsiCGgRdZf6dnNmagNuPLQ+/83DYg/5yq6CBdWW5W2KTzC9Y1X2J9og"
}
~~~

- https://github.com/leptos-null/LeptosMusic/blob/1505597c/music/Services/LMApiaryDeviceCrypto.h
- https://github.com/socialAPIS/YoutubeDownloader/blob/65b0ea2a/src/Request/DeviceRegistrationRequest.php

## key

todo:

~~~
AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w
~~~

10 stars:

https://github.com/socialAPIS/YoutubeDownloader

0 stars:

https://github.com/BorisChen396/PuddingPlayer

done:

~~~
AIzaSyA64xQnVODx8qBOeSsrlfDc8gDEw_NLopk
AIzaSyA_n-CBlmsO1fOxFUZqRnQ9SX4Bh1jCjWg
AIzaSyAxmTFlJLw9-uEJ1pFJUzw8LX7veGxGUoI
AIzaSyBD1uN7sPOWjkZ3fNKv7xDlLqF7Rg_JLnk
AIzaSyC8UYZpvA2eknNex0Pjid0_eTLJoDu6los
AIzaSyCChP9IaeaDS_LLGBI0P9CDQwTzCxn1kp8
AIzaSyCTa7aViyHnB3GLIqhL5hQFZGb675SoCIA
AIzaSyCV2I1gEhkJYkd51xG7MGaZGC85zylcS74
AIzaSyCX7NVTCfWMK8eEUau8Scc2y6dZUpWfNd0
AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw
AIzaSyCqrNxCAJrrk_NQqIUp1-baqW05d3JYeOc
AIzaSyCtkvNIR1HCEwzsqK6JuE6KqpyjusIRI30
AIzaSyCymf5PAosq7hWs5DkgHy0-3uacHaY1SPE
AIzaSyD5cCj3gK6IKFQCHRf1pYAt9nDKUzfxmPg
AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8
AIzaSyDHQ9ipnphqTzDqZsbtd8_Ru4_kiKVQe2k
AIzaSyDil7P0s1hvamdVWsqFtySc1T5P1S9dHqk
AIzaSyDjSMHkZSQWmcCKsNnvZcjRc2ZaJbAXpR4
AIzaSyDtpXO8h8u8Z6N7asPxy6AczIICsqmkg64
~~~

## tv

~~~
Mozilla/5.0 (Linux; Tizen 2.3; SmartTV)
~~~

- https://developers.google.com/identity/protocols/oauth2/native-app
- https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps
- https://developers.google.com/youtube/v3/live/guides/auth/installed-apps

## x-youtube

~~~
curl -v -o youtube.txt `
-H 'x-youtube-client-name: 1' `
-H 'x-youtube-client-version: 2.20191008.04.01' `
'https://www.youtube.com/watch?v=mvDWtMVPyf8&pbj=1'
~~~

## device approach

I did some more research, and was finally able to programmatically complete a
YouTube OAuth flow. Below is the "device" approach [1], I will put the "native
app" approach in a different comment. For "device" approach, first you make a
request like this:

~~~
curl -v `
-d client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com `
-d scope=https://www.googleapis.com/auth/youtube `
https://oauth2.googleapis.com/device/code
~~~

Then you get a response like this:

~~~json
{
 "device_code": "AH-1Ng2OQIJRBmX0cs1b6hzAABUNpPvuiX4fehXraWUVNAe4oiIOkQPkcRV...",
 "user_code": "SHCN-XMKQ",
 "expires_in": 1800,
 "interval": 5,
 "verification_url": "https://www.google.com/device"
}
~~~

Then, you direct user to visit `verification_url`, and enter `user_code`. After
they do that, you make a second request like this:

~~~
curl -v `
-d client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com `
-d client_secret=SboVhoG9s0rNafixCSGGKXAT `
-d code=AH-1Ng33Q8eEB-u7rdEETLWW7tTrM_DA1KcRnaklBTbjXyqbwFK41Hc-4afQ_qzfR8Eq... `
-d grant_type=http://oauth.net/grant_type/device/1.0 `
https://oauth2.googleapis.com/token
~~~

Result:

~~~
{
  "access_token": "ya29.a0ARrdaM8HlXT53-Zijh2w5_wTI1MDQM8lgwP-bFkLliw0hp-67R...",
  "expires_in": 3599,
  "refresh_token": "1//0fQueho-bnN3JCgYIARAAGA8SNwF-L9Ir0ebXoRbstIFr6E1LDz4_...",
  "scope": "https://www.googleapis.com/auth/youtube",
  "token_type": "Bearer"
}
~~~

The `access_token` lasts one hour, at which point you can renew it with the
long lived `refresh_token`.

1. https://developers.google.com/identity/sign-in/devices

## native app approach

Update: here is the process for "native app" [1]. First, program prompts user
to visit page like this:

~~~
https://accounts.google.com/o/oauth2/v2/auth?
client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com&
redirect_uri=urn:ietf:wg:oauth:2.0:oob&
response_type=code&
scope=https://www.googleapis.com/auth/youtube
~~~

As discussed previously, for the `redirect_uri`, you can also use
`http://localhost:999` or similar. Using the "manual copy/paste" method above,
a code will be returned to the user, which they can copy/paste into the program
(PyTube). Then, PyTube can make an internal request like this:

~~~
curl -v `
-d client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com `
-d client_secret=SboVhoG9s0rNafixCSGGKXAT `
-d code=4/1AX4XfWgGbS8R7Fza-TojWaRuE0QFiN4asvmmc07VKlSjsH0ghn3Sm5... `
-d grant_type=authorization_code `
-d redirect_uri=urn:ietf:wg:oauth:2.0:oob `
https://oauth2.googleapis.com/token
~~~

Response will be `access_token` and `refresh_token`, as before. If PyTube adds
OAuth, will just need to decide to use "device" approach or "native app"
approach, and if using "native app", whether to do "manual copy/paste" or
"loopback ip" method. At this point, I think I am done with this GitHub issue.
Unless someone finds a magic way with the Android keys I put, it seems OAuth is
the best way to handle this situation.

1. https://developers.google.com/identity/protocols/oauth2/native-app
