# YouTube

## Client

Can we use one client for decrypted media, `publishDate` and search? No.

Can we use two clients? Yes. Which two are best?

~~~
len 2390 ANDROID_KIDS
len 2698 IOS_KIDS
len 3023 WEB_KIDS
len 4565 ANDROID_EMBEDDED_PLAYER
len 5752 WEB_EMBEDDED_PLAYER
len 11813 IOS_CREATOR
len 22444 ANDROID_CREATOR
len 32702 IOS_MUSIC
len 40312 IOS
len 54491 ANDROID_MUSIC
len 54609 WEB_CREATOR
len 58280 MWEB
len 63598 ANDROID
len 70718 WEB
len 233260 WEB_REMIX
len 669127 TVHTML5
~~~

And alternative to `publishDate` is:

~~~
"text": "Published on Nov 5, 2020"
~~~

which you can get doing a `next` request with `IOS_KIDS`. Uncompressed size is
115 KB. For `publishDate`, doing a `player` request with `MWEB`, you get
uncompressed size of 57 KB.

~~~
publishDate fail ANDROID_CREATOR
publishDate fail ANDROID_EMBEDDED_PLAYER
publishDate fail ANDROID_KIDS
publishDate fail ANDROID_MUSIC
publishDate fail IOS
publishDate fail IOS_CREATOR
publishDate fail IOS_KIDS
publishDate fail IOS_MUSIC
publishDate fail TVHTML5
publishDate fail WEB_CREATOR
publishDate fail WEB_EMBEDDED_PLAYER
publishDate fail WEB_KIDS
publishDate pass MWEB
publishDate pass WEB
publishDate pass WEB_REMIX

decrypt fail ANDROID_EMBEDDED_PLAYER
decrypt fail ANDROID_KIDS
decrypt fail IOS_KIDS
decrypt fail MWEB
decrypt fail TVHTML5
decrypt fail WEB
decrypt fail WEB_CREATOR
decrypt fail WEB_EMBEDDED_PLAYER
decrypt fail WEB_KIDS
decrypt fail WEB_REMIX
decrypt pass ANDROID
decrypt pass ANDROID_CREATOR
decrypt pass ANDROID_MUSIC
decrypt pass IOS
decrypt pass IOS_CREATOR
decrypt pass IOS_MUSIC

player 200 OK ANDROID
player 200 OK ANDROID_CREATOR
player 200 OK ANDROID_EMBEDDED_PLAYER
player 200 OK ANDROID_KIDS
player 200 OK ANDROID_MUSIC
player 200 OK IOS
player 200 OK IOS_CREATOR
player 200 OK IOS_KIDS
player 200 OK IOS_MUSIC
player 200 OK MWEB
player 200 OK TVHTML5
player 200 OK WEB
player 200 OK WEB_CREATOR
player 200 OK WEB_EMBEDDED_PLAYER
player 200 OK WEB_KIDS
player 200 OK WEB_REMIX

search 200 OK ANDROID
search 200 OK ANDROID_EMBEDDED_PLAYER
search 200 OK ANDROID_KIDS
search 200 OK ANDROID_MUSIC
search 200 OK IOS
search 200 OK IOS_KIDS
search 200 OK IOS_MUSIC
search 200 OK MWEB
search 200 OK TVHTML5
search 200 OK WEB
search 200 OK WEB_CREATOR
search 200 OK WEB_KIDS
search 200 OK WEB_REMIX
search 400 Bad Request ANDROID_CREATOR
search 400 Bad Request IOS_CREATOR
search 400 Bad Request WEB_EMBEDDED_PLAYER
~~~

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
