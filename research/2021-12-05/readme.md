# December 5 2021

## Approach 1

Using this video:

https://nbc.com/la-brea/video/pilot/9000194212

and this request:

~~~
POST /access/vod/nbcuniversal/9000194212 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
User-Agent: Mozilla/5
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=43d1ebdb5dfe3a21b9d76f38d370cd83d8316076987d581be74d43562840aca1,time=1638665836328
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2410887629}}
~~~

