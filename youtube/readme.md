# YouTube

## Client

Can we use one client for decrypted media, `publishDate` and search? No. Can we
use two clients? Yes. Which two are best?

~~~
67796 IOS_CREATOR MWEB
67796 MWEB IOS_CREATOR
78242 ANDROID_CREATOR MWEB
78242 MWEB ANDROID_CREATOR
81549 IOS_CREATOR WEB
81549 WEB IOS_CREATOR
89389 IOS_MUSIC MWEB
89389 MWEB IOS_MUSIC
91995 ANDROID_CREATOR WEB
91995 WEB ANDROID_CREATOR
97358 IOS MWEB
97358 MWEB IOS
103142 IOS_MUSIC WEB
103142 WEB IOS_MUSIC
110019 ANDROID_MUSIC MWEB
110019 MWEB ANDROID_MUSIC
111111 IOS WEB
111111 WEB IOS
121632 ANDROID MWEB
121632 MWEB ANDROID
123772 ANDROID_MUSIC WEB
123772 WEB ANDROID_MUSIC
135385 ANDROID WEB
135385 WEB ANDROID
245595 IOS_CREATOR WEB_REMIX
245595 WEB_REMIX IOS_CREATOR
256041 ANDROID_CREATOR WEB_REMIX
256041 WEB_REMIX ANDROID_CREATOR
267188 IOS_MUSIC WEB_REMIX
267188 WEB_REMIX IOS_MUSIC
275157 IOS WEB_REMIX
275157 WEB_REMIX IOS
287818 ANDROID_MUSIC WEB_REMIX
287818 WEB_REMIX ANDROID_MUSIC
299431 ANDROID WEB_REMIX
299431 WEB_REMIX ANDROID
~~~

An alternative to `publishDate` is:

~~~
"text": "Published on Nov 5, 2020"
~~~

which you can get doing a `next` request with `IOS_KIDS`. Uncompressed size is
115 KB. For `publishDate`, doing a `player` request with `MWEB`, you get
uncompressed size of 57 KB.

- https://github.com/TeamNewPipe/NewPipeExtractor/issues/562
- https://github.com/yt-dlp/yt-dlp/pull/328

## MWEB

https://github.com/thanhphongdo/youtube-noads/blob/master/server/apis/youtube.ts

## Image

Given a video ID, return all the possible image links. Leave it up to end user
to make sure links are valid for a given video. Also add a test to ensure all
options are up to date.

## Search

~~~
https://github.com/yuliskov/MediaServiceCore/blob/master/youtubeapi/src/test/
resources/youtube-requests.http
~~~

## youtubei

Since ANDROID is smaller, we want to pull as much from it as possible. Here is
what we need to check:

~~~go
StreamingData struct {
   AdaptiveFormats []struct {
      Bitrate         int64  // pass
      ContentLength   int64  // pass
      Height          int    // pass
      Itag            int    // pass
      MimeType        string // pass
      URL             string // pass
   }
}
VideoDetails struct {
   Author           string // pass
   ShortDescription string // pass
   Title            string // pass
   ViewCount        int    // pass
}
Microformat struct {
   PlayerMicroformatRenderer struct {
      AvailableCountries []string // fail
      PublishDate        string   // fail
   }
}
~~~

- https://github.com/TeamNewPipe/NewPipeExtractor/issues/562
- https://github.com/TeamNewPipe/NewPipeExtractor/issues/568
- https://github.com/iv-org/invidious/issues/1981
- https://github.com/iv-org/invidious/pull/1985

## watch

desktop:

~~~
curl -o index.html https://www.youtube.com/watch?v=UpNXI3_ctAc
~~~

Next:

~~~html
<script nonce="GWQS4dROIhbOWa4QpveqWw">var ytInitialPlayerResponse = {"respons...
...ta":false,"viewCount":"11059","category":"Music","publishDate":"2020-10-02"...
...1"}},"adSlotLoggingData":{"serializedSlotAdServingDataEntry":""}}}]};</script>
~~~

Next:

~~~html
<script nonce="GWQS4dROIhbOWa4QpveqWw">var ytInitialPlayerResponse = {"respons...
...u0026sp=sig\u0026url=https://r4---sn-q4flrner.googlevideo.com/videoplayback...
...1"}},"adSlotLoggingData":{"serializedSlotAdServingDataEntry":""}}}]};</script>
~~~

mobile good:

~~~
Never Gonna Reach Me
curl -o index.html -A iPad https://m.youtube.com/watch?v=UpNXI3_ctAc
~~~

mobile bad:

~~~
Goon Gumpas
curl -o index.html -A iPad https://m.youtube.com/watch?v=NMYIVsdGfoo
~~~

## Free proxy list

https://proxy.webshare.io/register

## Links

- https://github.com/iawia002/annie/issues/839
- https://github.com/kkdai/youtube/issues/186
- https://golang.org/pkg/net/http#Header.WriteSubset
- https://superuser.com/questions/773719/how-do-all-of-these-save-video
