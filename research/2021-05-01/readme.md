# May 1 2021

Using ID `BnEn7X3Pr7o` [1] with program `manifest.go`, produces this:

~~~
https://manifest.googlevideo.com/api/manifest/dash/
expire/1619894815/
ei/v02NYIv5D__e2_gP-a-PmAs/
ip/72.181.23.38/
id/067127ed7dcfafba/
source/youtube/
requiressl/yes/
playback_host/r2---sn-q4flrnld.googlevideo.com/
mh/e2/
mm/31%2C26/
mn/sn-q4flrnld%2Csn-5uaeznyz/
ms/au%2Conr/
mv/m/
mvi/2/
pl/18/
hfr/all/
as/fmp4_audio_clear%2Cwebm_audio_clear%2Cwebm2_audio_clear%2Cfmp4_sd_hd_clear%2Cwebm2_sd_hd_clear/
initcwndbps/1887500/
vprv/1/
mt/1619872845/
fvip/2/
keepalive/yes/
fexp/24001373%2C24007246/
beids/23886219/
itag/0/
sparams/expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Chfr%2Cas%2Cvprv%2Citag/
sig/AOq0QJ8wRgIhAIhkHJzZK-f4FdEEDe_41v2xsXJPdt2VMVvNnjJaEwliAiEAzR3cHlKe8VRt-0qWaQ8TLmoE_VbVU6svQUBrbhySlw0%3D/
lsparams/playback_host%2Cmh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps/
lsig/AG3C_xAwRQIhAPTCmY8fNGtrmQEaVKnUilsJjF3m5FPSDB_Ivk6NL5ZdAiBd41s9Ezkfh256wRNVIAWk1nwNGHFKMFZ5MPEsJDEnlQ%3D%3D/
~~~

Same result with ID `9HzQvow8zF8` [2]. However, using ID `ipOogrq1m24` [3],
produces this result:

~~~
missing dashManifestUrl
~~~

These also produce a manifest:

- http://youtube.com/api/manifest/dash/id/3aa39fa2cc27967f/source/youtube?signature
- http://youtube.com/api/manifest/dash/id/bf5bb2419360daf1/source/youtube?signature

However if you replace the `id` with the one above:

~~~
youtube.com/api/manifest/dash/
id/067127ed7dcfafba/
source/youtube?signature
~~~

it fails:

~~~
Your client does not have permission to get URL
/api/manifest/dash/id/067127ed7dcfafba/source/youtube?signature from this
server. That’s all we know.
~~~

Same if we use the full URL:

~~~
youtube.com/api/manifest/dash/
id/067127ed7dcfafba/
source/youtube?as=fmp4_audio_clear,fmp4_sd_hd_clear&key=ik0&
sparams=ip,ipbits,expire,source,id,as&ip=0.0.0.0&ipbits=0&expire=19000000000&
signature=51AF5F39AB0CEC3E5497CD9C900EBFEAECCCB5C7.8506521BFC350652163895D4C26DEE124209AA9E
~~~

Same for the other three `signature`s:

~~~
youtube.com/api/manifest/dash/id/067127ed7dcfafba/
source/youtube?as=fmp4_audio_clear,fmp4_sd_hd_clear&key=ik0&
sparams=ip,ipbits,expire,source,id,as&ip=0.0.0.0&ipbits=0&expire=19000000000&
signature=A2716F75795F5D2AF0E88962FFCD10DB79384F29.84308FF04844498CE6FBCE4731507882B8307798

youtube.com/api/manifest/dash/id/067127ed7dcfafba/
source/youtube?as=fmp4_audio_clear,fmp4_sd_hd_clear&key=ik0&
sparams=ip,ipbits,expire,source,id,as&ip=0.0.0.0&ipbits=0&expire=19000000000&
signature=249B04F79E984D7F86B4D8DB48AE6FAF41C17AB3.7B9F0EC0505E1566E59B8E488E9419F253DDF413

youtube.com/api/manifest/dash/id/067127ed7dcfafba/
source/youtube?as=fmp4_audio_clear,fmp4_sd_hd_clear&key=ik0&
sparams=ip,ipbits,expire,source,id,as&ip=0.0.0.0&ipbits=0&expire=19000000000&
signature=B1C2A74783AC1CC4865EB312D7DD2D48230CC9FD.BD153B9882175F1F94BFE5141A5482313EA38E8D
~~~

Here are some more that work:

~~~
youtube.com/api/manifest/dash/
id/0894c7c8719b28a0/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature

youtube.com/api/manifest/dash/
id/48fcc369939ac96c/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature

youtube.com/api/manifest/dash/
id/d286538032258a1c/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature

youtube.com/api/manifest/dash/
id/e06c39f1151da3df/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature

youtube.com/api/manifest/dash/
id/efd045b1eb61888a/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature

youtube.com/api/manifest/dash/
id/f9a34cab7b05881a/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature
~~~

Again, it fails if we change `id`:

~~~
youtube.com/api/manifest/dash/
id/067127ed7dcfafba/
source/youtube?as=fmp4_audio_cenc,fmp4_sd_hd_cenc&signature
~~~

## References

1. https://www.youtube.com/watch?v=BnEn7X3Pr7o
2. https://www.youtube.com/watch?v=9HzQvow8zF8
3. https://www.youtube.com/watch?v=ipOogrq1m24
