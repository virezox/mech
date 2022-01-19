# Vimeo

- https://developer.vimeo.com/api/reference/videos
- https://github.com/httptoolkit/frida-android-unpinning/issues/8
- https://stackoverflow.com/questions/1361149
- https://vimeo.com/66531465

We start with this:

https://player.vimeo.com/video/660408476/config

Then we get this key:

~~~
Request.Files.DASH.CDNs.Fastly_Skyfire.URL
~~~

Which returns this:

~~~
https://skyfire.vimeocdn.com/
1642617968-0x7a08308e0140eef7e86b97d3e477c08de7af3769/
64a97917-f2a3-46b6-a4cc-3e55e3dd07a8/sep/
video/f18dad7b,fb8654f4,e4c45277,25c3eb1e,e34a69bd,9861e040,450eb343/
audio/e9081717,30021b64/
master.json
~~~

We can call the first URL `Config`. We can call the second URL `Master`.
