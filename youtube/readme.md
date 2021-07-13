# YouTube

To anyone interested, I figured out how to do an authenticated request to
`/youtubei/v1/player`. First you have the normal body:

~~~json
{
   "context": {
      "client": {
         "clientName": "WEB",
         "clientVersion": "2.20210708.06.00",
      }
   },
   "videoId": "bO7PgQ-DtZk"
}
~~~

probably can adjust the client as needed. You still want to add the `key` to
querystring as normal. Next, you need these headers:

~~~
"X-Origin","https://www.youtube.com"
"Authorization","SAPISIDHASH 1625981319_5d6eb..."
~~~

Finally, you need these cookies:

~~~
Name: "__Secure-3PSID",
Value: "_gdNhnpLL2zVVA9c-gj53X-bZQipAWuXEccm0..."

Name: "__Secure-3PAPISID"
Value: "VFKYV_f44SBoEuOa/AxyyZj1QZKPY..."
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
