# YouTube

## Client

~~~
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
- https://github.com/tombulled/innertube
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
