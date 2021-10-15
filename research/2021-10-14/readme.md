# October 14 2021

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358

Example video link:

<https://vod-lf-oneapp-prd.akamaized.net/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/4830k_720_hls/_1927821650-2_00000.ts>

Comes from here:

<https://east.manifest.na.theplatform.com/m/NnzsPC/mf_htCaWEEbX,IBiRGzW1xuHR,9cQum5DswcH2,0zDLGDs3SnDq,cMb56n9qo7UZ,jUzY0qw3LRjj/4.m3u8?sid=1ed55084-70ac-42cf-a738-0655f867ee0c&policy=188569381&date=1634266602515&ip=72.181.23.38&schema=1.1&cid=25e25682-f115-4b92-b592-22104d09add1&host=vod-lf-oneapp-prd.akamaized.net&meta=false&manifest=M3U&switch=HLSServiceSecure&am_sdkv=null&_fw_did=a858dd5f-f1a2-464a-934a-fd5c679a6c15&_fw_h_referer=www.nbc.com&siteSectionId=oneapp_desktop_computer_web_ondemand&nw=169843&am_extmp=default&tracking=true&uuid=a858dd5f-f1a2-464a-934a-fd5c679a6c15&am_appv=null&mode=on-demand&uoo=0&mparticleid=2746124636930297065&us_privacy_string=1YNN&sfid=9244655&player=%5Bv2%5D+OneApp+-+PDK6+NBC.com&am_buildv=null&am_abvrtd=0&csid=oneapp_desktop_computer_web_ondemand&metr=1023&am_cpsv=4.0.0-2&userAgent=Mozilla%2F5.0+%28Windows+NT+10.0%3B+Win64%3B+x64%3B+rv%3A86.0%29+Gecko%2F20100101+Firefox%2F86.0&afid=200265138&prof=nbcu_web_svp_js_https&fallbackSiteSectionId=9244655&vpaid=script&_fw_vcid2=169843%3A2746124636930297065&am_crmid=2746124636930297065&rdid=28575ed3-15f8-d6e5-5c86-c92bb7058035&am_playerv=null&sdk=PDK+6.4.3&did=28575ed3-15f8-d6e5-5c86-c92bb7058035&am_stitcherv=poc&am_abtestid=0&sig=eeffb779f9fd0e443270197a3f419a2c77d5afadb1a1e9ecf880a49353e41fab>

You can download without query string, but result file will be basically empty.

~~~
[NBC] 9000199358: Downloading JSON metadata
GET /v2/graphql?query=query+bonanzaPage%28%0A++%24app%3A+NBCUBrands%21+%3D+nbc%0A++%24name%3A+String%21%0A++%24oneApp%3A+Boolean%0A++%24platform%3A+SupportedPlatforms%21+%3D+web%0A++%24type%3A+EntityPageType%21+%3D+VIDEO%0A++%24userId%3A+String%21%0A%29+%7B%0A++bonanzaPage%28%0A++++app%3A+%24app%0A++++name%3A+%24name%0A++++oneApp%3A+%24oneApp%0A++++platform%3A+%24platform%0A++++type%3A+%24type%0A++++userId%3A+%24userId%0A++%29+%7B%0A++++metadata+%7B%0A++++++...+on+VideoPageData+%7B%0A++++++++description%0A++++++++episodeNumber%0A++++++++keywords%0A++++++++locked%0A++++++++mpxAccountId%0A++++++++mpxGuid%0A++++++++rating%0A++++++++resourceId%0A++++++++seasonNumber%0A++++++++secondaryTitle%0A++++++++seriesShortTitle%0A++++++%7D%0A++++%7D%0A++%7D%0A%7D&variables=%7B%22name%22%3A+%22http%3A%2F%2Fwww.nbc.com%2Fsaturday-night-live%2Fvideo%2Foctober-2-owen-wilson%2F9000199358%22%2C+%22oneApp%22%3A+true%2C+%22userId%22%3A+%220%22%7D HTTP/1.1
Host: friendship.nbc.co
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: en-us,en;q=0.5
Connection: close

[NBC] 9000199358: Downloading JSON metadata
send: b'GET /s/NnzsPC/media/guid/2410887629/9000199358?format=preview HTTP/1.1\r\nHost: link.theplatform.com\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nConnection: close\r\n\r\n'

[ThePlatform] 9000199358: Downloading SMIL data
send: b'GET /s/NnzsPC/media/guid/2410887629/9000199358?mbr=true&manifest=m3u&format=SMIL HTTP/1.1\r\nHost: link.theplatform.com\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nConnection: close\r\n\r\n'

[ThePlatform] 9000199358: Downloading m3u8 information
send: b'GET /i/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/HD_TVE_SATURDAYNIGHTLIVE_10022021_V2_,185,783,483,300,86,35,0k.mp4.csmil/master.m3u8?hdnea=st=1634267857~exp=1634280487~acl=/i/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/HD_TVE_SATURDAYNIGHTLIVE_10022021_V2_*~id=4720ba91-93f5-409b-b5ce-1dfc38d1db99~hmac=f0ece90bce58633b57d4e52f07266e75076d5c001de310838bc78718a69f1622 HTTP/1.1\r\nHost: nbcmpx-vh.akamaihd.net\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nConnection: close\r\n\r\n'

[ThePlatform] 9000199358: Downloading JSON metadata
send: b'GET /s/NnzsPC/media/guid/2410887629/9000199358?format=preview HTTP/1.1\r\nHost: link.theplatform.com\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nConnection: close\r\n\r\n'

[info] 9000199358: Downloading 1 format(s): hls-7932
[hlsnative] Downloading m3u8 manifest
send: b'GET /i/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/HD_TVE_SATURDAYNIGHTLIVE_10022021_V2_,185,783,483,300,86,35,0k.mp4.csmil/index_1_av.m3u8?null=0&id=AgBItRcmFy81YfDyaGGl3NjtcNaCFh6+fmwlBodz3fqvrLfZSL23EHhV5ddT2krYepBAD1Co6TTp9w%3d%3d&hdntl=exp=1634354288~acl=%2fi%2fprod%2fvideo%2f1j7%2fmXg%2f9000199358%2fCEyiDynfCGvaJjdJ0uGUx%2fHD_TVE_SATURDAYNIGHTLIVE_10022021_V2_*~data=hdntl~hmac=b01b784a50885f78d11b4ae56632d53e51762c6fd7dc38aae3178e98d37f94a6 HTTP/1.1\r\nHost: nbcmpx-vh.akamaihd.net\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: en-us,en;q=0.5\r\nCookie: _alid_=EUV/aRsdyXujP9BV/pdP2A==\r\nConnection: close\r\n\r\n'

[hlsnative] Total fragments: 398
[download] Destination: October 2 - Owen Wilson [9000199358].mp4
send: b'GET /i/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/HD_TVE_SATURDAYNIGHTLIVE_10022021_V2_,185,783,483,300,86,35,0k.mp4.csmil/segment1_1_av.ts?null=0&id=AgBItRcmFy81YfDyaGGl3NjtcNaCFh6+fmwlBodz3fqvrLfZSL23EHhV5ddT2krYepBAD1Co6TTp9w%3d%3d&hdntl=exp=1634354288~acl=/i/prod/video/1j7/mXg/9000199358/CEyiDynfCGvaJjdJ0uGUx/HD_TVE_SATURDAYNIGHTLIVE_10022021_V2_*~data=hdntl~hmac=b01b784a50885f78d11b4ae56632d53e51762c6fd7dc38aae3178e98d37f94a6 HTTP/1.1\r\nAccept-Encoding: identity\r\nHost: nbcmpx-vh.akamaihd.net\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.9 Safari/537.36\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-us,en;q=0.5\r\nCookie: _alid_=EUV/aRsdyXujP9BV/pdP2A==\r\nConnection: close\r\n\r\n'
~~~
