# Paramount

how to get MPD?

<https://www.paramountplus.com/shows/video/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU/>

~~~
eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU
~~~

HLS:

~~~
> curl -I 'link.theplatform.com/s/dJ5BDC/media/guid/2198311517/z6jldNkH8VlhsBzmTyEjE7AlGi1Ros3s?formats=MPEG4,M3U&assetTypes=StreamPack'
Location: https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/04/15/1885672515948/CBS_THE_NEIGHBORHOOD_314_2398_630320_,503,4628,3128,2228,1628,848,000.mp4.csmil/master.m3u8?hdnea=acl=/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/04/15/1885672515948/CBS_THE_NEIGHBORHOOD_314_2398_630320_*~exp=1653189790~hmac=54b8bf3b28bd2b73372bd2168f75106e1abc1e59de69da069b1364aaf4a5304c
~~~

DASH:

~~~
> curl -I 'link.theplatform.com/s/dJ5BDC/media/guid/2198311517/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU?assetTypes=DASH_CENC&formats=MPEG-DASH'
Location: https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2021/08/31/1940767811923/993595_cenc_dash/stream.mpd
~~~

https://github.com/matthuisman/slyguy.addons/issues/198
