# July 12 2021

Age gate videos that can be embedded:

https://www.youtube.com/embed/HtVdAasjOgU

Age gate videos that can cannot be embedded:

https://www.youtube.com/embed/bO7PgQ-DtZk

Old:

https://github.com/89z/mech/blob/160417ea/youtube/video.go#L50-L58

I think I figured it out! So currently most people are using `/get_video_info`
endpoint with the `TVHTML5` client [1], which returns something like this:

~~~json
{
  "streamingData": {
    "adaptiveFormats": [
      {
        "signatureCipher": "s=%3Dgz%3DgzXlPC_hz9AwbkcIP73w7qqqnVzWsANhh2nhY7J50PnAEiAukA6dg2vte5oi3O0Jm4W37PZPdF9rJh6yBdFmUho1NNAhIgRw8JQ0qOBOSBOS&sp=sig&url=https://r3---sn-q4flrnl6.googlevideo.com/videoplayback%3Fexpire%3D1626131386%26ei%3DWnfsYPOJBcLWigSRvrjAAw%26ip%3D72.181.23.38%26id%3Do-AN84JAdDnsjZ_LCOgV7ZJQghJ9X_h-43rkODyCDGtZKK%26itag%3D137%26aitags%3D133%252C134%252C135%252C136%252C137%252C160%252C242%252C243%252C244%252C247%252C248%252C278%26source%3Dyoutube%26requiressl%3Dyes%26mh%3DES%26mm%3D31%252C26%26mn%3Dsn-q4flrnl6%252Csn-5uaezn6y%26ms%3Dau%252Conr%26mv%3Dm%26mvi%3D3%26pl%3D18%26initcwndbps%3D1863750%26vprv%3D1%26mime%3Dvideo%252Fmp4%26ns%3D4jt-ZWY4Y-jKkDEIQeP0fTQG%26gir%3Dyes%26clen%3D116413993%26dur%3D260.226%26lmt%3D1623845851267731%26mt%3D1626109037%26fvip%3D3%26keepalive%3Dyes%26fexp%3D24001373%252C24007246%26c%3DTVHTML5%26txp%3D5535434%26n%3DnsWDB-laC4YqsKUEyU%26sparams%3Dexpire%252Cei%252Cip%252Cid%252Caitags%252Csource%252Crequiressl%252Cvprv%252Cmime%252Cns%252Cgir%252Cclen%252Cdur%252Clmt%26lsparams%3Dmh%252Cmm%252Cmn%252Cms%252Cmv%252Cmvi%252Cpl%252Cinitcwndbps%26lsig%3DAG3C_xAwRAIgCjOdKSmt5rDGUdUh02qF6vxuUK4BmdoksI4bCVh07XQCIFCV4M-Y7nnlGVzPUFfaKpkZTeG0XPvVRoxhFt0IovLR"
~~~

but, just like with the `/youtubei/v1/player` endpoint, you can change the
client. So if you instead use the `ANDROID` client:

~~~
https://www.youtube.com/get_video_info?c=ANDROID&cver=16.05&eurl=https%3A%2F%2Fwww.youtube.com&html5=1&video_id=bO7PgQ-DtZk
~~~

You get this:

~~~
{
  "streamingData": {
    "adaptiveFormats": [
      {
        "itag": 137,
        "url": "https://r3---sn-q4flrnl6.googlevideo.com/videoplayback?expire=1626131842&ei=InnsYIvYDZuMlu8P7qaRoAU&ip=72.181.23.38&id=o-APUbeSrtxOTu5eQlCMdszApT6RvaaSLnBQhoCEXux_Z8&itag=137&source=youtube&requiressl=yes&mh=ES&mm=31%2C29&mn=sn-q4flrnl6%2Csn-q4fl6ns7&ms=au%2Crdu&mv=m&mvi=3&pl=18&initcwndbps=1850000&vprv=1&mime=video%2Fmp4&gir=yes&clen=116413993&dur=260.226&lmt=1623845851267731&mt=1626110006&fvip=3&keepalive=yes&fexp=24001373%2C24007246&c=ANDROID&txp=5535434&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AOq0QJ8wRAIgPi8M6zwc1VDe2hUUzFlPZcMCuy8vRRfTkgm7cFxEnDUCIEXuf0_rcGAuSrP1BF3RUBLSTQgXxl7OROCUe-giG5OY&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRQIhAPili0SMqtDzFnlkMTkMwsLj0AV_Mt445Phsfm6tWf7sAiAM0fuPWKninLVO1gEXhVjeGqjl9FpaLUHcLzAuR-SHkQ%3D%3D"
~~~

With this method, not only is the URL already decrypted (no JavaScript needed),
but the method even works with age gate videos that are not embeddable! Maybe
you can give it a try to see if any problems.

1. <https://github.com/ytdl-org/youtube-dl/blob/a8035827/youtube_dl/extractor/youtube.py#L1508-L1516>
