# YouTube-DL

first:

~~~
GET https://api.config.watchabc.go.com/appsconfig/prod/abcv3/031_04/10.23.0/config.json HTTP/2.0
user-agent: com.disney.datg.videoplatforms.android.abc/10.23.1 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
datg-usertz: -0500
accept-encoding: gzip
content-length: 0
~~~

then:

~~~
https://api.contents.watchabc.go.com/vp2/ws/s/contents/3001/videos/001/031_04/
lf,es,mp,sf/%SHOW%/%SEASON%/%VIDEO%/%START%/%LIMIT%.json
~~~

then:

~~~
GET /vp2/ws/contents/3000/videos/001/001/-1/-1/-1/VDKA26847512/-1/-1.json HTTP/1.1
Host: api.contents.watchabc.go.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.11 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close
content-length: 0
~~~

then:

~~~
POST /vp2/ws-secure/entitlement/2020/authorize.json HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 56
Host: api.entitlement.watchabc.go.com
X-Forwarded-For: 6.128.0.30
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.111 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close

video_id=VDKA26847512&video_type=lf&brand=001&device=001
~~~

then:

~~~
GET /ext/d874124ecca24c88a3c9575e78686acf/5e6dda216a5e43309543689d3fd85f91.m3u8?cqs=v52pCbLlPtngdpG8c9oApMWDzg1zuLCANHsI9k1CTdacMV8lSP6Pi9PBNMczXDGTeOUv135TmHogDbx2QHPeYuMkb_A5_6L00-Hb2V6Lzap2xeQ49yLs7hGGom762RFHcVo__iCSYbP5hu0G8XB_CsgIRBf7Cgk_lUhTT7fHQOSyfHr7d1H_RswYY47ERk8UjYo8XYw66qxJjbOBXeFFKNSdyzlOBZNKrMYCz-gQrtXI-DaaqbI75ibL4Gbit1Q4NWILzi1Ym8ZJBelofkhy645RgOBzJyeab5VnARdJ4nt9rk3YcU21OHil8pMCE8QFNDpPkeBGJwuYeTtjIXkvmizy_O1hYLZCHxhxDkaaHdizMUK-5Jngg0pEhgjkcFDF9PZh_s462qppFosFcdJFP6huLsraNNs2U4Hi5L6aE_HVbAeGkkIqA3nyLfw0n1sTf6pl1VozhlE_rUavlujgTrS6Rq3NWIM2vd2Q8kE0Auk=&kid=4098a6a720374bfcbb4e362b652bcd51 HTTP/1.1
Host: content-dtci.uplynk.com
X-Forwarded-For: 6.128.0.30
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.111 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close
content-length: 0
~~~
