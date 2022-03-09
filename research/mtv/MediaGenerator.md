# MediaGenerator

Using this:

~~~
mtv -i -a https://www.mtv.com/episodes/scyb0g/aeon-flux-utopia-or-deuteranopia-season-1-ep-1
~~~

I get this result:

~~~
Resolution:320x240 Bandwidth:268860 Codecs:avc1.4d401e,mp4a.40.2 Audio:audio0
Resolution:480x360 Bandwidth:410154 Codecs:avc1.4d401f,mp4a.40.2 Audio:audio0
Resolution:576x432 Bandwidth:659235 Codecs:avc1.4d401f,mp4a.40.2 Audio:audio0
Resolution:640x480 Bandwidth:910392 Codecs:avc1.4d401f,mp4a.40.2 Audio:audio0
Resolution:768x576 Bandwidth:1287890 Codecs:avc1.4d401f,mp4a.40.2 Audio:audio0
~~~

YT-DLP gets this:

~~~
#EXTM3U
#EXT-X-VERSION:4
## Created by Viacom Delivery Service(version=2.0.7)
#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=594790,BANDWIDTH=634075,FRAME-RATE=23.976,CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=480x360
https://dlvrsvc.mtvnservices.com/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_480x360_534749.m3u8?tk=st=1646699602~exp=1646714002~acl=/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_480x360_534749.m3u8*~hmac=0fa90af738deff43732acc9702591991e2b0f161442401059b2424e794552482&account=mtv.com&cdn=ns1
#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=369952,BANDWIDTH=404182,FRAME-RATE=23.976,CODECS="avc1.4D401E,mp4a.40.2",RESOLUTION=320x240
https://dlvrsvc.mtvnservices.com/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_320x240_315417.m3u8?tk=st=1646699602~exp=1646714002~acl=/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_320x240_315417.m3u8*~hmac=9fdb5a907ea705298cfd6366519fa7e2e800445ba37424ca4ff02a642e947f3b&account=mtv.com&cdn=ns1
#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2136273,BANDWIDTH=2291064,FRAME-RATE=23.976,CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=768x576
https://dlvrsvc.mtvnservices.com/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_768x576_2041574.m3u8?tk=st=1646699602~exp=1646714002~acl=/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_768x576_2041574.m3u8*~hmac=0c22ce20163bfa6b7a31e12237115db12b43eaf664a84f3a11e0eb236003368a&account=mtv.com&cdn=ns1
#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=1627949,BANDWIDTH=1736653,FRAME-RATE=23.976,CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=640x480
https://dlvrsvc.mtvnservices.com/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_640x480_1544131.m3u8?tk=st=1646699602~exp=1646714002~acl=/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_640x480_1544131.m3u8*~hmac=ab22a944b1188d5a5062e0151afc9ab1aee05e8e0a3d101cebb616d810141703&account=mtv.com&cdn=ns1
#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=1158789,BANDWIDTH=1233071,FRAME-RATE=23.976,CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=576x432
https://dlvrsvc.mtvnservices.com/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_576x432_1086196.m3u8?tk=st=1646699602~exp=1646714002~acl=/api/playlist/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/stream_576x432_1086196.m3u8*~hmac=ccb81c7cbf83cbe67e0646a0e4790e735b3e9b866dcabc2767ec8df464e1ed96&account=mtv.com&cdn=ns1
~~~

How are they getting this:

~~~
BANDWIDTH=2291064,FRAME-RATE=23.976,CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=768x576
~~~

First:

~~~
GET /feeds/mrss/?uri=mgid%3Aarc%3Aepisode%3Amtv.com%3A96defb28-d238-11e1-a549-0026b9414f30 HTTP/1.1
Host: www.mtv.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

Then:

~~~
GET /services/MediaGenerator/mgid:arc:video:mtv.com:d2eca5e3-a0c9-4058-9fb9-fd7459787e52?&arcPlatforms=7ac7942e-6481-457f-b39a-2b1aedb29f29,b995f21c-e76f-4e58-8d0f-0964dc76efd3,6caa8f01-72e5-4707-abd6-608a0146e2ee,39dfe10c-cc2a-40f9-a20d-962d2604d543,0a16f611-d105-436a-8188-33ea8871171e,f17d0e9b-657a-4785-b3ca-dae9e78563a1&acceptMethods=hls HTTP/1.1
Host: media-utils.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

Then:

~~~
GET /api/gen/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/,stream_320x240_315417,stream_768x576_2041574,stream_640x480_1544131,stream_480x360_534749,stream_576x432_1086196/master.m3u8?account=mtv.com&cdn=ns1&tk=st=1646699601~exp=1646786001~acl=/api/gen/gsp.alias/mediabus/mtv.com/2018/11/28/11/06/52/ab596fbe5aba46b0acfc19c6046d34c8/974095/0/*~hmac=01361c6b53ffc56b952a759c90c532fb71c4400ef874ab6964ea4e020631101b HTTP/1.1
Host: dlvrsvc.mtvnservices.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
