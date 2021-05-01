# May 1 2021

Using ID `BnEn7X3Pr7o` [1] with program `manifest.go`, produces this:

~~~
https://manifest.googlevideo.com/api/manifest/dash/
expire/1619894815/
ei/v02NYIv5D__e2_gP-a-PmAs/
ip/72.181.23.38/
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
id/067127ed7dcfafba/
~~~

Same result with ID `9HzQvow8zF8` [2]:

~~~
https://manifest.googlevideo.com/api/manifest/dash/
expire/1619901023/
ei/_2WNYK3lMfHDlu8PvKWQuAk/
ip/72.181.23.38/
source/youtube/
requiressl/yes/
playback_host/r1---sn-q4fl6nsl.googlevideo.com/
mh/Vn/
mm/31%2C29/
mn/sn-q4fl6nsl%2Csn-q4flrn7s/
ms/au%2Crdu/
mv/m/
mvi/1/
pl/18/
tx/23964902/
txs/23964900%2C23964901%2C23964902%2C23964903%2C23964904%2C23964907/
hfr/all/
as/fmp4_audio_clear%2Cwebm_audio_clear%2Cwebm2_audio_clear%2Cfmp4_sd_hd_clear%2Cwebm2_sd_hd_clear/
initcwndbps/1923750/
vprv/1/
mt/1619879095/
fvip/6/
keepalive/yes/
fexp/24001373%2C24007246/
itag/0/
sparams/expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Ctx%2Ctxs%2Chfr%2Cas%2Cvprv%2Citag/
sig/AOq0QJ8wRQIgV-VXneyQGJCMwVlF01WUM91-vN6wJiGnc6kte8h6q4UCIQD_ty_6ojVd5kkLaYZdwCCRgALdW8u4Kzhd_Tgz3Bncqw%3D%3D/
lsparams/playback_host%2Cmh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps/
lsig/AG3C_xAwRgIhAIZ6lx4eMQVCpUYd-YY7_jRdwhH2QvhGwwTwZYYk-4i8AiEAu9bd_JR3t1YMuLPTlFJ6RUu_o2DZHl-K9v2y7bXLzLA%3D/
id/f47cd0be8c3ccc5f/
~~~

If we swap the two `id`s, they both fail. Using ID `ipOogrq1m24` [3], produces
this result:

~~~
missing dashManifestUrl
~~~

Here are some magic `id`s:

~~~
youtube.com/api/manifest/dash/
id/0894c7c8719b28a0/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/3aa39fa2cc27967f/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/48fcc369939ac96c/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/bf5bb2419360daf1/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/cf74bf361c2e79a3/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/d286538032258a1c/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/e06c39f1151da3df/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/efd045b1eb61888a/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear

youtube.com/api/manifest/dash/
id/f9a34cab7b05881a/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear
~~~

The URLs fail if you use a different `id` though:

~~~
youtube.com/api/manifest/dash/
id/f47cd0be8c3ccc5f/
source/youtube?signature&as=fmp4_audio_clear,fmp4_sd_hd_clear
~~~

Here is the JavaScript path:

~~~
ytplayer.config.args.raw_player_response.streamingData.dashManifestUrl
~~~

Here is source for magic:

https://github.com/google/ExoPlayer/blob/0e4d3162/demos/main/src/main/assets/media.exolist.json

## References

1. https://www.youtube.com/watch?v=BnEn7X3Pr7o
2. https://www.youtube.com/watch?v=9HzQvow8zF8
3. https://www.youtube.com/watch?v=ipOogrq1m24
