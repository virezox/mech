# July 16 2021

If it helps, I found a new way to access some videos without cookie:

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

> @89z We can already download all age-gated videos with logged in cookies.

So what? Cookies are not the end all solution to authentication. For example,
in addition to suggestions I have put forward, other options might be
available, such as monitoring the YouTube Android app. I know other apps have
special keys hidden inside the APK. For example the Android deezer.com app
hides a special key inside a PNG file [1]. YouTube could be doing something
similar.

> **If** someone can figure out a way to download them without authentication,
> then we can think about implementing it.

Its not possible, people need to accept that.

> The discussion of any method that requires auth is useless since **it can
> already be done** with the existing code

If you still think that after reading my paragraph above, then suit yourself
and close the issue.

> I agree. But when there are workarounds, it is good to implement them. If
> there are none, then we just need to accept that fact and start blaming
> youtube instead of youtube-dl for it.

The only workaround at this point is authentication. That can be done different
ways, if you just want to only support the cookie method, thats your choice.
But better auth methods might be available.

> This is why this issue is marked as "enhancement" and not "bug" and this is
> the same reason why @colethedj originally closed this issue. This **is not a
> bug**.

I dont care what you call it, and I dont care if you close the issue. I am
actively doing research on the topic, so I thought it would be good for people
to work together to a common goal. But if your goal is different than mine, or
if your goal already done, then close issue and I will continue on my own.

1. <https://github.com/89z/mech/blob/master/research/2021-05-12/android_gw_key.py>
