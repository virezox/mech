# YouTube

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

## How to get X-Goog-Api-Key

Next open your browser menu, and click Web Developer, Network or similar. Then
go to:

https://www.youtube.com

then go back to Developer Tools, and you should see a JSON request like this:

~~~
POST /youtubei/v1/guide?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8 HTTP/1.1
Host: www.youtube.com
~~~

https://cloud.google.com/apis/docs/system-parameters

## How to get `client_id` and `client_secret`

Set User-Agent to [1]:

~~~
Mozilla/5.0 (SMART-TV; LINUX; Tizen 5.0)
~~~

Then visit:

https://www.youtube.com/tv

Next open your browser menu, and click Web Developer, Network or similar. Then
go back to the page, and click "Sign in", then "Sign in with a web browser". On
the next page, dont bother with any of the instructions, just go back to
Developer Tools, and after about five seconds you should see a JSON request like
this:

~~~
POST /o/oauth2/token HTTP/1.1
Host: www.youtube.com

{"client_id":"861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com",
"client_secret":"SboVhoG9s0rNafixCSGGKXAT",
"code":"AH-1Ng14qVvEj76OeM_h14Mgklgyhchbyc67MhULhCKPY6K-0DTYJqaKng2ULVFTmTzU...",
"grant_type":"http://oauth.net/grant_type/device/1.0"}
~~~

References:

1. https://developer.samsung.com/smarttv/develop/guides/fundamentals/retrieving-platform-information.html

## Free proxy list

https://proxy.webshare.io/register
