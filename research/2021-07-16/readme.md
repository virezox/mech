# July 16 2021

## oauth

~~~json
{
  "access_token": "ya29.a0ARrdaM8IoKZHM_WiGDZpTGASVZ2-FlAEG9zmlboNF5iCovX0b7...",
  "expires_in": 3599,
  "scope": "https://www.googleapis.com/auth/youtube-paid-content https://www.googleapis.com/auth/youtube",
  "token_type": "Bearer"
}
~~~

I found a dead simple program for getting OAuth tokens [1]:

~~~
step oauth
~~~

so I might just see if I can incorporate that into my own code.

https://github.com/smallstep/cli#features

## token

I found a new way to access some videos without cookie:

~~~
POST /youtubei/v1/player HTTP/2
Host: www.youtube.com
authorization: Bearer ya29.a0ARrdaM98ywSJCBrj6BgkO9WSnPaK_XNLB-KcEKlPlS7j9o2S_...
content-type: application/json

{
  "videoId": "Cr381pDsSsA",
  "context": {
    "client": {
      "clientName": "TVHTML5",
      "clientVersion": "7.20210713.10.00"
    }
  }
}
~~~

and it works with `/get_video_info` too:

~~~
https://www.youtube.com/get_video_info?
c=TVHTML5&
cver=7.20210428.10.00&
el=detailpage&
html5=1&
video_id=Cr381pDsSsA&
access_token=ya29.a0ARrdaM_CTfzZC0eE_Yxk04oowauAra_Z8Zh6mO5UzN84XYNot32JyYqcCG...
~~~

I dont really have a good way to get it right now. You can try browsing like
this, to get you started:

~~~
originalUrl:
https://www.youtube.com/tv

userAgent one of these:
Mozilla/5.0 (Linux; Tizen 2.3; SmartTV)
Mozilla/5.0 (Linux; Tizen 2.3; SMART-TV)
~~~

I am trying to see if the process can be streamlined. If you are able to get
the token, its good for one hour, then you can refresh like this:

~~~
POST /o/oauth2/token HTTP/2
Host: www.youtube.com
content-type: application/json

{
  "grant_type": "refresh_token",
  "client_id": "...-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com",
  "client_secret": "SboVhoG9s0rNafixCSG...",
  "refresh_token": "1//0fvQm8klSPFldCgYIARAAGA8SNwF-L9Ir9T6EnZnEXuNoQt-V8UpOQ..."
}
~~~
