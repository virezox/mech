# YouTube

## Client

Can we use one client for decrypted media, `publishDate` and search? No. Can we
use two clients? Yes. Which two are best?

~~~
110690 MWEB {decrypt:false publishDate:true search:true size:56784}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53906}

119227 ANDROID {decrypt:true publishDate:false search:true size:62443} MWEB
{decrypt:false publishDate:true search:true size:56784}

123978 WEB {decrypt:false publishDate:true search:true size:70072}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53906}

132515 ANDROID {decrypt:true publishDate:false search:true size:62443} WEB
{decrypt:false publishDate:true search:true size:70072}

289237 WEB_REMIX {decrypt:false publishDate:true search:true size:235331}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53906}

297774 ANDROID {decrypt:true publishDate:false search:true size:62443}
WEB_REMIX {decrypt:false publishDate:true search:true size:235331}
~~~

Can we use `MWEB` for `publishDate`? Yes. Can we use `MWEB` for search? Yes.
Can we use `ANDROID_CREATOR` for decrypted media? Yes.

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
