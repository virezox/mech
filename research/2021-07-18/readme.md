# July 18 2021

## key

What `key`s are available? First, we have `youtube.com`:

~~~
AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8
~~~

What about `youtube.com/tv`:

~~~
Mozilla/5.0 (Linux; Tizen 2.3; SmartTV)
AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8
~~~

## oauth

> I'd be more interested in trying to find the flow for generating a token in
> the first place.

You cannot get an OAuth token in the way you are decribing, that I am aware of.
The flow MUST go through the browser. If someone finds a way around that, I
will be shocked. Using the browser, you have two options. For the first option,
the client program (PyTube), can direct user to visit a page like the following
(`client_id` borrowed from SmallStep [1]):

~~~
https://accounts.google.com/o/oauth2/v2/auth?
client_id=1087160488420-8qt7bavg3qesdhs6it824mhnfgcfe8il.apps.googleusercontent.com&
redirect_uri=urn:ietf:wg:oauth:2.0:oob&
response_type=code&
scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube
~~~

Then after your user completes the flow, they get an OAuth code like this:

~~~
4/1AX4XfWisI0eOngHMlafYoucs5EQqrO0GF6LI7A5_kQYfu6SGd_PD1X9hV24
~~~

which PyTube can use to make a request internally and get an `access_token`. A
better option is to instead use a `redirect_uri` such as
`http://localhost:999`. This way, once user completes OAuth flow, the OAuth
code is instead delivered the localhost address, where PyTube could intercept
it without needing the user to copy paste. If you want to see a demonstration
of this process, get the SmallStep tool [2], and run a command like this:

~~~
step oauth -scope https://www.googleapis.com/auth/youtube
~~~

1. https://github.com/smallstep/cli/blob/ca448947/command/oauth/cmd.go#L45
2. https://github.com/smallstep/cli/releases

## step

> What I'm a little unclear on is how I can allow people to generate that token
> programmatically (i.e. pass in a username and password as part of the
> application, and generate the oauth token that way).

Still working on this. You can generate the access token here pretty easy:

https://developers.google.com/oauthplayground

You just choose `YouTube Data > https://www.googleapis.com/auth/youtube`, then
click `Authorize APIs`. However the problem is anything created from the
playground:

> Note: The OAuth Playground will automatically revoke refresh tokens after
> 24h. You can avoid this by specifying your own application OAuth credentials
> using the Configuration panel.

So for a long lived refresh token, similar to long lived cookie, someone has to
create their own client. You can create your own client here:

https://console.developers.google.com/apis/credentials

but that has drawbacks too, as unless you verify your client with Google,
anyone using your client get the "unverified application" scary screen. It
still works, its just annoying. And the verification process can take two weeks
I read, and you have to submit a video to Google explaining what your app does.
Whats also weird, is the OAuth YouTube requests require FULL scope, not just
readonly. So any app that someone grants this access to, has FULL access to
their YouTube account.

After thinking about all that, I thought maybe its better to just give
instructions to people on how to set up their own app/client. Whats weird also
is even if you sign into your own account, with YOUR OWN client, you still get
"unverified app" warning. I dont know. I am still thinking about the different
options to decide what is best. Finally I did notice one thing when testing
today. When using `/youtubei/v1/player` endpoint, with either `Authorization`
header or `access_token` query parameter, the `key` parameter is not needed. So
maybe some special key is out there that can bypass all this. I decompiled the
YouTube Android app yesterday. I didnt find anything crazy, but I wasnt looking
for that either, so I might look again.

## methods

Can we do `/get_video_info` with `Authorization` header? Yes:

~~~
curl -v -H 'Authorization: Bearer ya29.a0ARrdaM9gbcgLGqyyJDKV3agV60hXypXRdj...' `
'https://www.youtube.com/get_video_info?c=TVHTML5&cver=7.20210428.10.00&' +
'el=detailpage&html5=1&video_id=Cr381pDsSsA'
~~~

Can we do `/get_video_info` with `Cookie` header? Yes:

~~~
curl -v `
-H 'Cookie: __Secure-3PSID=_wdNhho4fTGOQgwrBshGc4QT3onztCB6xvfzwFh7OrWNTMjn...' `
('https://www.youtube.com/get_video_info?c=TVHTML5&cver=7.20210428.10.00&' +
'el=detailpage&html5=1&video_id=Cr381pDsSsA')
~~~

Can we do `/get_video_info` with query parameter? Yes:

~~~
curl -v ('https://www.youtube.com/get_video_info?c=TVHTML5&' +
'cver=7.20210428.10.00&el=detailpage&html5=1&video_id=Cr381pDsSsA&' +
'access_token=ya29.a0ARrdaM9gbcgLGqyyJDKV3agV60hXypXRdjspznK-4N2JF9nzW6z7BS9...')
~~~

Can we do `/youtubei/v1/player` with `Authorization` header? Yes:

~~~
curl -v `
-H 'Authorization: Bearer ya29.a0ARrdaM-R9c0crGrwdmS-FRugGQbDdRtkVEM5YFOm6d...' `
-H 'content-type: application/json' `
-d '@youtube.json' `
https://www.youtube.com/youtubei/v1/player
~~~

Can we do `/youtubei/v1/player` with `Cookie` header? Yes:

~~~
curl -v `
-H 'Authorization: SAPISIDHASH 1626641976_8b25e923dce0638caf2ef1d7dbac1e253...' `
-H 'Content-Type: application/json' `
-H 'Cookie: __Secure-3PSID=_wdNhnFIjWLd...; __Secure-3PAPISID=Y35R-aq0Bud-...' `
-H 'X-Origin: https://www.youtube.com' `
-d '@youtube.json' `
'https://www.youtube.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8'
~~~

Can we do `/youtubei/v1/player` with query parameter? Yes:

~~~
curl -v -H 'Content-Type: application/json' -d '@youtube.json' `
'https://www.youtube.com/youtubei/v1/player?access_token=ya29.a0ARrdaM8XMbTHQ...'
~~~

## Links

- https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps
- https://github.com/pytube/pytube/issues/1057
