# YouTube

Age gate videos that can be embedded:

https://www.youtube.com/embed/HtVdAasjOgU

Age gate videos that can cannot be embedded:

https://www.youtube.com/embed/bO7PgQ-DtZk

Old:

https://github.com/89z/mech/blob/160417ea/youtube/video.go#L50-L58

Tests:

<https://github.com/Hexer10/youtube_explode_dart/blob/master/test/video_test.dart>

~~~
PS D:\Desktop> youtube-dl.exe --print-traffic MeJVWBSsPAY
GET /watch?v=MeJVWBSsPAY&bpctr=9999999999&has_verified=1 HTTP/1.1
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Connection: close
Host: www.youtube.com\r\nCookie: CONSENT=YES+cb.20210328-17-p0.en+FX+512
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.6 Safari/537.36
~~~

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
