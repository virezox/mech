# Paramount+

~~~
yt-dlp --proxy 127.0.0.1:8080 --no-check-certificate `
paramountplus.com/shows/star-trek-prodigy/video/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy/star-trek-prodigy-dreamcatcher
~~~

- https://github.com/ytdl-org/youtube-dl/issues/30491
- https://play.google.com/store/apps/details?id=com.cbs.app

This is it:

~~~
GET /i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_,2228,4628,3128,1628,848,503,000.mp4.csmil/segment2_1_av.ts?null=0&id=AgBItRcmFya85M6kHmKe34Eu2IzJ1IHI2He5etlGwZvVWDnuRjMBNgeWRd%2fmV2csh0f3yRlamjyDXg%3d%3d&hdntl=exp=1646261838~acl=/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_*~data=hdntl~hmac=c6fd17c4197dae64d2b9a055624536c851f2788e4e00524c6a8e391e511d4d04 HTTP/1.1
Accept-Encoding: identity
Host: cbsios-vh.akamaihd.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Cookie: _alid_=MbBNIhoN43TV17fPr98BXQ==; hdntl=exp=1646261838~acl=%2fi%2ftemp_hd_gallery_video%2fCBS_Production_Outlet_VMS%2fvideo_robot%2fCBS_Production_Entertainment%2f2021%2f10%2f18%2f1963091011554%2fNICKELODEON_STARTREKPRODIGY_104_HD_985058_*~data=hdntl~hmac=c6fd17c4197dae64d2b9a055624536c851f2788e4e00524c6a8e391e511d4d04
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_,2228,4628,3128,1628,848,503,000.mp4.csmil/index_1_av.m3u8?null=0&id=AgBItRcmFya85M6kHmKe34Eu2IzJ1IHI2He5etlGwZvVWDnuRjMBNgeWRd%2fmV2csh0f3yRlamjyDXg%3d%3d&hdntl=exp=1646261838~acl=%2fi%2ftemp_hd_gallery_video%2fCBS_Production_Outlet_VMS%2fvideo_robot%2fCBS_Production_Entertainment%2f2021%2f10%2f18%2f1963091011554%2fNICKELODEON_STARTREKPRODIGY_104_HD_985058_*~data=hdntl~hmac=c6fd17c4197dae64d2b9a055624536c851f2788e4e00524c6a8e391e511d4d04 HTTP/1.1
Host: cbsios-vh.akamaihd.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Cookie: _alid_=MbBNIhoN43TV17fPr98BXQ==; hdntl=exp=1646261838~acl=%2fi%2ftemp_hd_gallery_video%2fCBS_Production_Outlet_VMS%2fvideo_robot%2fCBS_Production_Entertainment%2f2021%2f10%2f18%2f1963091011554%2fNICKELODEON_STARTREKPRODIGY_104_HD_985058_*~data=hdntl~hmac=c6fd17c4197dae64d2b9a055624536c851f2788e4e00524c6a8e391e511d4d04
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_,2228,4628,3128,1628,848,503,000.mp4.csmil/master.m3u8?hdnea=acl=/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2021/10/18/1963091011554/NICKELODEON_STARTREKPRODIGY_104_HD_985058_*~exp=1646175557~hmac=d431ae3ab4581e4e089a0aa9ec76dc3a364c4130a4359b42fb6d9398ae2b7c0c HTTP/1.1
Host: cbsios-vh.akamaihd.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~

which comes from:

~~~
GET /s/dJ5BDC/media/guid/2198311517/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy?format=SMIL&formats=MPEG4%2CM3U HTTP/1.1
Host: link.theplatform.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.20 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Sec-Fetch-Mode: navigate
Connection: close
content-length: 0
~~~
