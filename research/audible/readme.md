# Audible

- https://github.com/89z/mech/issues/14
- https://play.google.com/store/apps/details?id=com.audible.application

Need to use Frida. Audio looks like this:

~~~
GET /bk_adbl_003303/2/signed/g1/bk_adbl_003303_22_64.mp4?id=8a3ac406-656e-444a-893e-3665dc9a0523&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=MD1BAJ27VEMQ8GGY7BZB%2F20220217%2Fus-east-1%2Fcloudfront%2Faws4_request&X-Amz-Date=20220217T044820Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host%3Buser-agent&X-Amz-Signature=8d59f224bcc663263b22578be1e3586e93cb710d697dfbba7c262489aa51bcc2 HTTP/1.1
Range: bytes=0-9999999
User-Agent: com.audible.playersdk.player/3.21.0 (Linux;Android 7.0) ExoPlayerLib/2.14.2
Accept-Encoding: identity
Host: d1jobzhhm62zby.cloudfront.net
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
GET /bk_adbl_003303/2/signed/g1/bk_adbl_003303_22,32,64_v7.master.mpd?ss_sec=20&iss_sec=10&isc=1&id=8a3ac406-656e-444a-893e-3665dc9a0523&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=MD1BAJ27VEMQ8GGY7BZB/20220217/us-east-1/cloudfront/aws4_request&X-Amz-Date=20220217T044820Z&X-Amz-Expires=60&X-Amz-SignedHeaders=host&X-Amz-Signature=f68f1ce44015a9bcdd58a757bbbba831b4ef7510482732a5a2c70df0c80fe9bb HTTP/1.1
User-Agent: com.audible.playersdk.player/3.21.0 (Linux;Android 7.0) ExoPlayerLib/2.14.2
Accept-Encoding: gzip
Host: d1jobzhhm62zby.cloudfront.net
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
POST /1.0/content/B00551W570/licenserequest HTTP/1.1
Accept-Charset: utf-8
Accept: application/json
x-adp-signature: MN9WDIkd3qb8E+funPEvsMzvTvZe5QUCC1uXX1VbPI24DY5QWnNzY8/fdTXowOz77MIt1URFCIFV9wNhEb5SbXpjff6h1morL35WhQa98Scu8pZTa7zDZifI/STXsgtH1BDL4KPxvlEY5QtKIeEr76GCjpdEQixYtZ1S14uXd4m3ZD+i2EF1ctN5tx950H0utIXEYYaQh8YdBgtb0YLcjjlfkhQMadGiPDAkeOzt//zBW++rC9Uey1p8TbaU+ndYrS6Hn2w+mqgzniMgffdqH/lQM6kX61AuDFmsH3gEvMwV9RqoMDRZ9zvhFJWi42TmdZkqjdev7GSkcX+M0ECdMA==:2022-02-16T22:48:21Z
x-adp-token: {enc:T1q+czVAQ/YHNSIzeOQ5fuTcZG0whnPrL1LOGou6lSh7EreQXVc+MaFuS0x/FHKYpQxOg7iPc0rZZusHLpxssOC+qrjbGd2LbnPGJ54e5TZFS2IzBvP8/DUH7nFDfCIfnmcSfnpIUgHHglW8vGWOvaZNakzvcgY7NPdHEn2l7KBlFmw/zB7zgm8flBRL+1qufER06I21F6ACNqNnHUVjjCLMEUs2RL/XRhkEH++58hZ1uJAqcqO4iuHmQej2T5DjpOE3ZAbLscDcLuB0l1r9XOq4dlJSIqnqDQVuU9mcY2YjFcjIpgNmj88T4VrhhNid7O5Q148eIM2Hzi9kQWwQjPrsue3i2vEPIDGNNzA9k+hhQc8xvzaBI0cse8KsNhXperFSJn7oyj24B+Zye9u5akPKnEo8GnqgW9hwNcXQKh1t/mMJdE9D4lgkgIA54xm8MgfIvUYwu8WjHZkb/555KSvTCCuL1tVj9I4UJF+ApYZZK0ulF6swGBiksAoFnL0W7d7vUXsWundbKDsN66yrmKEBb/y2dUDENvF4jJjg/CfS5FWWbA9Vnb8rakOLtMGH4oWW+11xupNlnUFxtFqUm2wSPQFqRgBBr8bEz7SOvj1PNo0vjoYv+bLLf+dCaAFqdAiqT45bNGnxISwwFyqI8E2CVJ1lysViC5ko5vq8EYHS4oobp0Vx3uc4aciJV9sbv+JaFK29z8+cunrea55wRqKub0FIqf6rxXyHYWTw5lfql9qZm40tifjvDhcTRlcBZgHi/ubyLJdxNMJFcbmj3/JQw5wGx6ibTwPGRDYauWH8wixOcaBEVEyNTVFqfn8flvQMfn7dlHKfwJJ9SCrflNS9zgZ70mACahvZwkMqciQh4Jno7SNq7sxbxJZ1G3DO2zytA4KH7BF9au1NaXcii9lGdJiGRgKURKzVfhyK5nmRjQpT8UkjZjuCu6CCp0IRGfrvknmd+mldNge/Kg/ZFOha9W7MqOejjl4EHyT6eMBiWrLNCvLtssbEJPiSL+BIiSnca/FRwSXXu/A+1GDaacxnO4L0ubBkAytTvejfz4s=}{key:JPUuMlehlhRVleS6cNidcxP5tMlGKSE1+N1FezIk7tyxoCvINR3c+ormG4UzUFetBeeaX33Lh7QISwfBNT4cmF6p2lh6AzYBP58EvCit2rxSn/XEf6WenDOcs9J1xFQWJKx2JsFRnm5r14WllhX473ArrCGnilE/iN4zSG/xKk/YFa8io1CvR1aLqJQaz/wJwdIKzM4g9Motdty98kOXtugIel5h4wIabGDtjU6GwGpDsB8CoDJEJmw0+aZEBjSVxTlgKfV6Q/1olzQgSnAde23cdT7OkPUwh9JNGBIYCk0XnAkArz3QsBex165AJLs9MLu6YeSHX1sT1rjhYzm18A==}{iv:gcCAyMzZ1Cn0RlPMPWy07Q==}{name:QURQVG9rZW5FbmNyeXB0aW9uS2V5}{serial:Mg==}
x-adp-alg: SHA256WithRSA:1.0
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC); com.audible.application 3.21.0 b:102028
Accept-Language: en-US
Content-Type: application/json
Content-Length: 197
Host: api.audible.com
Connection: Keep-Alive
Accept-Encoding: gzip

{"supported_drm_types":["Dash","Hls","Mpeg"],"consumption_type":"Streaming","use_adaptive_bit_rate":true,"response_groups":"content_reference,chapter_info,pdf_url,last_position_heard,ad_insertion"}
~~~
