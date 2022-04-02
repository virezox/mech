# Paramount

- <https://paramountplus.com/shows/the-harper-house/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/the-harper-house-the-harper-house>
- https://github.com/matthuisman/slyguy.addons/tree/master/slyguy.paramount.plus
- https://kodi.tv/download/windows

~~~
https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/
video_robot/CBS_Production_Entertainment/2020/06/08/1747974723591/
MTV_AEONFLUX_110_206218_,503,4628,3128,2228,1628,848,000.mp4.csmil/master.m3u8?
hdnea=acl=/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/
CBS_Production_Entertainment/2020/06/08/1747974723591/
MTV_AEONFLUX_110_206218_*~exp=1648914311~hmac=e6d1a344f851f39b9ab63ed77719359733e2f234eeb51470a65c996f3274c719
~~~

deep:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://www.paramountplus.com/shows/the-harper-house/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/the-harper-house-the-harper-house/
~~~

app:

~~~
com.cbs.app
~~~

This looks promising:

~~~
https://www.paramountplus.com/apps-api/v2.0/androidphone/videos/section/
292602.json?begin=0&rows=30&locale=en-us&
at=ABB%2FAZwa9fcxXKGHkYznjNfYetF8Jj1wgZkllraTIgOa3PIaL1aa4qmzBBlQErkhjYw%3D
~~~

response:

~~~
{manifest:none}
   https://entclips.cbsaavideo.com/2021/07/14/1921578563637/732460_dash_ta/stream.mpd
{manifest:m3u}
   https://entclips.cbsaavideo.com/{slistFilePath}2021/07/14/1921578563637/732460_dash_ta/stream.mpd{slistFilePath}?hdnea=st={date:-30:true}~exp={date:259200:true}~acl=/{slistFilePath}*~hmac={token:AkamaiEdgeAuth:2f85d2369afa3c0efbd4b2ff7d5edca2}
{manifest}
~~~
