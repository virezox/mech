# YouTube

## itag 302

If we look at video `youtube.com/watch?v=3gdfNdilGFE`:

~~~
itag 302, height 720, 3.5 mb/s, 139.4 MB, video/webm; codecs="vp9"
~~~

If we look at video `youtube.com/watch?v=XeojXq6ySs4`:

~~~
itag 247, height 720, 127.0 kb/s, 5.9 MB, video/webm; codecs="vp9"
~~~

Can we get `itag 247` on both videos, with any client? No. Can we get `itag
302` on both videos, with any client? No. For either video, can we get
`dashManifestUrl`, with any client? No.

## itag 251

If we look at video `3gdfNdilGFE`:

~~~
67201 WEB_CREATOR {decrypt:true publishDate:false search:true size:56980}
WEB_REMIX {decrypt:false publishDate:true search:true size:10221}

80760 WEB_EMBEDDED_PLAYER {decrypt:true publishDate:false search:false
size:70539} WEB_REMIX {decrypt:false publishDate:true search:true size:10221}

99252 ANDROID {decrypt:true publishDate:false search:true size:89031}

151916 MWEB {decrypt:true publishDate:true search:true size:62885}

564097 WEB {decrypt:true publishDate:true search:true size:561332}

690545 TVHTML5 {decrypt:true publishDate:false search:true size:680324}
WEB_REMIX {decrypt:false publishDate:true search:true size:10221}
~~~

If we look at video `XeojXq6ySs4`:

~~~
110560 MWEB {decrypt:false publishDate:true search:true size:57396}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53164}

120580 MWEB {decrypt:false publishDate:true search:true size:57396} ANDROID
{decrypt:true publishDate:false search:true size:63184}
~~~

## Client

Can we use one client for decrypted media, `publishDate` and search? No. Can we
use two clients? Yes. Which two are best?

~~~
120597 MWEB {decrypt:false publishDate:true search:true size:56111} ANDROID
{decrypt:true publishDate:false search:true size:64486}

135304 ANDROID {decrypt:true publishDate:false search:true size:64486} WEB
{decrypt:false publishDate:true search:true size:70818}

295700 ANDROID {decrypt:true publishDate:false search:true size:64486}
WEB_REMIX {decrypt:false publishDate:true search:true size:231214}
~~~

Can we use `MWEB` for `publishDate`? Yes. Can we use `MWEB` for search? Yes.
Can we use `ANDROID` for decrypted media? Yes.

- https://github.com/TeamNewPipe/NewPipeExtractor/issues/562
- https://github.com/yt-dlp/yt-dlp/pull/328

## publishDate

An alternative to `publishDate` is:

~~~
"text": "Published on Nov 5, 2020"
~~~

which you can get doing a `next` request with `IOS_KIDS`. Uncompressed size is
115 KB. For `publishDate`, doing a `player` request with `MWEB`, you get
uncompressed size of 57 KB.

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
