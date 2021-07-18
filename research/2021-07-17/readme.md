# July 17 2021

~~~
client_id=407408718192.apps.googleusercontent.com
client_secret=************&scope=

Mozilla/5.0 (Linux; Tizen 2.3; SmartTV)

curl -v `
-d client_id=861556708454-912i5jlic99ecvu3ro5kqirg0hldli5t.apps.googleusercontent.com `
-d client_secret=ju2WuMJMOjilz_h_1dPgFdeU `
-d code=4%2F0AX4XfWivlrlWotm2r4AWgaF6FOVkRwtOCssa6bT3vfpBqBf0QieZ5Ogl3_3VJYRuQ_jqwA `
-d grant_type=authorization_code `
-d redirect_uri=https%3A%2F%2Fdevelopers.google.com%2Foauthplayground `
https://oauth2.googleapis.com/token
~~~

## done

https://developers.google.com/oauthplayground

## todo

- https://console.developers.google.com/apis/credentials
- https://console.developers.google.com/apis/credentials/oauthclient
- https://developers.google.com/identity/protocols/oauth2/native-app
- https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps
- https://developers.google.com/youtube/v3/live/guides/auth/installed-apps

## oauth

Note: The OAuth Playground will automatically revoke refresh tokens after 24h.
You can avoid this by specifying your own application OAuth credentials using
the Configuration panel.

## step

~~~
.\step oauth -scope https://www.googleapis.com/auth/youtube

Docker Machine, dlorenc@google.com:
https://accounts.google.com/o/oauth2/v2/auth?
client_id=22738965389-8arp8bah3uln9eoenproamovfjj1ac33.apps.googleusercontent.com&
redirect_uri=urn:ietf:wg:oauth:2.0:oob&
response_type=code&
scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube

Google Container Registry, jsand@google.com:
https://accounts.google.com/o/oauth2/v2/auth?
client_id=99426463878-o7n0bshgue20tdpm25q4at0vs2mr4utq.apps.googleusercontent.com&
redirect_uri=urn:ietf:wg:oauth:2.0:oob&
response_type=code&
scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube
~~~

Exchange:

https://github.com/smallstep/cli/blob/ca448947/command/oauth/cmd.go#L875-L897
