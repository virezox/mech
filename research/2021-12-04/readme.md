# December 5 2021

- <https://github.com/yt-dlp/yt-dlp/blob/818faa3a/yt_dlp/downloader/fragment.py#L357-L369>
- <https://nowtilus.gitbook.io/serverside-ai/getting-started/onboarding-document/prepare_content_source/ad-marker-specification/ssai-live-scte35>
- <https://scte-cms-resource-storage.s3.amazonaws.com/ANSI_SCTE-35-2019a-1582645390859.pdf>
- <https://wagtail-prod-storage.s3.amazonaws.com/documents/ANSI_SCTE-35-2020-1619708851007.pdf>
- <https://webstore.ansi.org/preview-pages/SCTE/preview_ANSI+SCTE+35+2019.pdfo>
- https://adobe.com/content/dam/acom/en/devnet/primetime/PrimetimeDigitalProgramInsertionSignalingSpecification.pdf
- https://datatracker.ietf.org/doc/html/rfc8216
- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://www.nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358

## Approach 2

~~~
GET /s/NnzsPC/media/guid/2410887629/9000194212?manifest=m3u HTTP/1.1
Host: link.theplatform.com
~~~

## Progressive

If you make a request like this:

http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL

You get results like this:

~~~html
<video
src="http://nbcmpx-vh.akamaihd.net/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_7830k.mp4?hdnea=st=1638675793~exp=1638688423~acl=/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_*~id=f46974e2-ba14-44ef-8e3e-5764799dc427~hmac=e1e983df3726aad47d33f5d328216f96d068c8fc67386a86efdb806a774af0bb"
system-bitrate="7932473" height="1080" width="1920"/>
~~~

However it appears none of the links actually work:

~~~
Access Denied
You don't have permission to access
"http://nbcmpx-vh.akamaihd.net/z/prod/video/udN/MbI/9000210182/Y5xLPDQ92UZQXliQ79_zS/HD_TVE_THEBLACKLIST_10212021_7830k.mp4?"
on this server.
Reference #18.100ec617.1638675890.1159b69a 
~~~

Also, these requests:

- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL&switch=http
- http://link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000210182?format=SMIL&switch=progressive

Return empty links:

~~~html
<video src="" system-bitrate="7932473" height="1080" width="1920"/>
~~~

https://github.com/ytdl-org/youtube-dl/issues/7806

## AES

[ThePlatform] 9000194212: Downloading SMIL data:

~~~
GET /s/NnzsPC/media/guid/2410887629/9000194212?mbr=true&manifest=m3u&format=SMIL HTTP/1.1
Host: link.theplatform.com
~~~

[ThePlatform] 9000194212: Downloading m3u8 information:

~~~
GET /i/prod/video/ZHU/RUI/9000194212/eMWA3j2InLZkHf6h7ry5p/HD_TVE_LABREA_09282021_V2_,185,783,483,300,86,35,0k.mp4.csmil/master.m3u8?hdnea=st=1638647763~exp=1638660393~acl=/i/prod/video/ZHU/RUI/9000194212/eMWA3j2InLZkHf6h7ry5p/HD_TVE_LABREA_09282021_V2_*~id=b8268349-1a36-4a6e-b94d-ed942450a243~hmac=0f7240907fb00ee4393d07974890b0abd155e4d22520502bb20381bab868414b HTTP/1.1
Host: nbcmpx-vh.akamaihd.net
~~~

[hlsnative] Downloading m3u8 manifest:

~~~
GET /i/prod/video/ZHU/RUI/9000194212/eMWA3j2InLZkHf6h7ry5p/HD_TVE_LABREA_09282021_V2_,185,783,483,300,86,35,0k.mp4.csmil/index_5_av.m3u8?null=0&id=AgBItRcmaHVClG7Gq2HZUUyl1qZdj4KJELImX0HjnAsH4hpnN9fe%2fl+p46%2fSo1llzQ%2fAnyv+kyKHrg%3d%3d&hdntl=exp=1638733806~acl=%2fi%2fprod%2fvideo%2fZHU%2fRUI%2f9000194212%2feMWA3j2InLZkHf6h7ry5p%2fHD_TVE_LABREA_09282021_V2_*~data=hdntl~hmac=3cf354717cb39041d13a6733a12a819dc9c144a1ff3d2e42a08f2b977039ef42 HTTP/1.1
Host: nbcmpx-vh.akamaihd.net
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3717.2 Safari/537.36
~~~
