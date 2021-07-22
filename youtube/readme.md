# YouTube

## Device OAuth

- https://datatracker.ietf.org/doc/html/rfc8628
- https://developers.google.com/identity/sign-in/devices

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
