# YouTube

## Client

For `publishDate` and search, seems best option is `WEB_REMIX`. For decrypted
media, seems best option is `ANDROID_MUSIC` or `WEB_EMBEDDED_PLAYER`. If we
look at video `HtVdAasjOgU`:

~~~
99988 WEB_EMBEDDED_PLAYER {decrypt:true publishDate:false search:false
size:89580} WEB_REMIX {decrypt:false publishDate:true search:true size:10408}

104193 WEB_EMBEDDED_PLAYER {decrypt:true publishDate:false search:false
size:89580} WEB {decrypt:false publishDate:true search:true size:14613}

102311 WEB_EMBEDDED_PLAYER {decrypt:true publishDate:false search:false
size:89580} MWEB {decrypt:false publishDate:true search:true size:12731}
~~~

If we look at video `XeojXq6ySs4`:

~~~
116361 MWEB {decrypt:false publishDate:true search:true size:53801} ANDROID
{decrypt:true publishDate:false search:true size:62560}

107769 MWEB {decrypt:false publishDate:true search:true size:53801}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53968}

132827 WEB {decrypt:false publishDate:true search:true size:70267} ANDROID
{decrypt:true publishDate:false search:true size:62560}

124235 WEB {decrypt:false publishDate:true search:true size:70267}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53968}

123103 WEB_REMIX {decrypt:false publishDate:true search:true size:60543}
ANDROID {decrypt:true publishDate:false search:true size:62560}

114511 WEB_REMIX {decrypt:false publishDate:true search:true size:60543}
ANDROID_MUSIC {decrypt:true publishDate:false search:true size:53968}
~~~

Can we use one client for decrypted media, `publishDate` and search? No. Can we
use two clients? Yes. Which two are best?

## publishDate

An alternative to `publishDate` is:

~~~
"text": "Published on Nov 5, 2020"
~~~

which you can get doing a `next` request with `IOS_KIDS`. Uncompressed size is
115 KB. For `publishDate`, doing a `player` request with `MWEB`, you get
uncompressed size of 57 KB.

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
- https://github.com/yt-dlp/yt-dlp/pull/328

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
