# Web client

This is it:

~~~
GET /gsp.originmusicstor/reencode/mtv.com/onair/cribs/0/seg_320x240_449098_1.ts HTTP/1.1
Accept-Encoding: identity
Host: mtv-ns1.ts.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /api/playlist/gsp.originmusicstor/reencode/mtv.com/onair/cribs/0/stream_320x240_449098.m3u8?tk=st=1645751662~exp=1645766062~acl=/api/playlist/gsp.originmusicstor/reencode/mtv.com/onair/cribs/0/stream_320x240_449098.m3u8*~hmac=c67a1a4c6c86759d1e49636b84a70c5f2f8d0296a941a96e515da771567f7b96&account=mtv.com&cdn=ns1 HTTP/1.1
Host: dlvrsvc.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /api/gen/gsp.originmusicstor/reencode/mtv.com/onair/cribs/,1/stream_320x240_171299_1045700095,0/stream_320x240_449098/master.m3u8?account=mtv.com&cdn=ns1&tk=st=1645751662~exp=1645838062~acl=/api/gen/gsp.originmusicstor/reencode/mtv.com/onair/cribs/*~hmac=ea88f00476c51e4b7dbbc633895f37ce1c2a66918e2a37a344cb0922a6e8c200 HTTP/1.1
Host: dlvrsvc.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /services/MediaGenerator/mgid:arc:video:mtv.com:d26f2b22-097d-11e3-8a73-0026b9414f30?&arcPlatforms=7ac7942e-6481-457f-b39a-2b1aedb29f29,b995f21c-e76f-4e58-8d0f-0964dc76efd3,6caa8f01-72e5-4707-abd6-608a0146e2ee,39dfe10c-cc2a-40f9-a20d-962d2604d543,0a16f611-d105-436a-8188-33ea8871171e,f17d0e9b-657a-4785-b3ca-dae9e78563a1&acceptMethods=hls HTTP/1.1
Host: media-utils.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /feeds/mrss/?uri=mgid%3Aarc%3Ashowvideo%3Amtv.com%3Ad26f2b22-097d-11e3-8a73-0026b9414f30 HTTP/1.1
Host: www.mtv.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
