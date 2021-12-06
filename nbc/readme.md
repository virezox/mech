# NBC


## SCTE-35

- <https://wagtail-prod-storage.s3.amazonaws.com/documents/ANSI_SCTE-35-2020-1619708851007.pdf>
- <https://webstore.ansi.org/preview-pages/SCTE/preview_ANSI+SCTE+35+2019.pdfo>
- https://adobe.com/content/dam/acom/en/devnet/primetime/PrimetimeDigitalProgramInsertionSignalingSpecification.pdf
- https://datatracker.ietf.org/doc/html/rfc8216
- <https://scte-cms-resource-storage.s3.amazonaws.com/ANSI_SCTE-35-2019a-1582645390859.pdf>
- <https://nowtilus.gitbook.io/serverside-ai/getting-started/onboarding-document/prepare_content_source/ad-marker-specification/ssai-live-scte35>
- <https://github.com/yt-dlp/yt-dlp/blob/818faa3a/yt_dlp/downloader/fragment.py#L357-L369>

Using this video:

https://nbc.com/la-brea/video/pilot/9000194212

and this request:

~~~
POST /access/vod/nbcuniversal/9000194212 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
User-Agent: Mozilla/5
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=43d1ebdb5dfe3a21b9d76f38d370cd83d8316076987d581be74d43562840aca1,time=1638665836328
content-type: application/json

{"device":"android","deviceId":"android","externalAdvertiserId":"NBC",
"mpx":{"accountId":2410887629}}
~~~

Then filter:

~~~
rg -e SCTE35 -e DISCONTINUITY -e ^https:// 2.m3u8 > filter.m3u8
~~~

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

Get APK:

~~~
com.nbcuni.nbc
~~~

MITM Proxy gives this:

~~~
POST /access/vod/nbcuniversal/9000194212 HTTP/1.1
Host: access-cloudpath.media.nbcuni.com
user-agent: Mozilla/5
authorization: NBC-Security key=android_nbcuniversal,version=2.4,hash=008d72003a9440598d06faac0715d2c8aa214e6a60549a55d8e168baf70911ad,time=1638589119812

{
   "device": "android",
   "deviceId": "2b7cd9d6f310611",
   "externalAdvertiserId": "NBC_VOD_9000194212",
   "mpx": {"accountId": "2304985974"}
}
~~~

hash comes from this request:

~~~
GET /config/android/player/v10/prod/mobile.json HTTP/1.1
Host: nbcapp.nbc.co
~~~

Response:

~~~json
{
   "name": "VODPlayer",
   "class": "com.nbc.cpc.player.cloudpath.CloudpathPlayer",
   "chromecastClass": "com.nbc.cpc.player.ottchromecast.OTTCastPlayer",
   "appKey": "android_nbcuniversal",
   "secretKey": "2b84a073ede61c766e4c0b3f1e656f7f",
   "accessUrl": "https://access-cloudpath.media.nbcuni.com/access/vod/nbcuniversal/",
   "accessMetadataURL": "https://access-cloudpath.media.nbcuni.com/access/vod/%s/%s/metadata",
   "enableAccessMetadata": false,
   "platformNameSpace": "pl1",
   "specificConfig": {
      "freewheel": {
         "default": {
            "network_id": "169843"
         }
      }
   }
}
~~~

`accountId` can also be found in the response:

~~~json
"channelsConfig": [
   {
      "id": "nbc",
      "Title": "NBC",
      "resourceID": "nbcentertainment",
      "contentSecurityLevel": "full",
      "live": {
         "playerModule": "LocalPlayer",
         "bionicModule": "LinearBionicNBC"
      },
      "vod": {
         "playerModule": "VODPlayer",
         "mpxAccountId": "2304985974",
         "metadataUrl": "https://feed.theplatform.com/f/HNK2IC/%sd_app_adstitch_v3_prod?byGUID=%s",
         "brightlineID": "1022"
      }
~~~

Androguard reveals this:

~~~java
v11_4.append("NBC-Security key=");
v11_4.append(this.appkey);
v11_4.append(",version=");
v11_4.append(com.nbc.cpc.cloudpathshared.CloudpathShared.getConfigServerVersion());
v11_4.append(",hash=");
v11_4.append(v0_5);
v11_4.append(",time=");
v11_4.append(v4_20);
v10_6.put("authorization", v11_4.toString());
~~~

Where time is:

~~~java
v4_20 = String.valueOf(new java.util.Date().getTime());
~~~

and hash is:

~~~
v0_5 = this.generateHash(v4_20, this.secretKey.getBytes());
~~~

Function here:

~~~java
private String generateHash(String p3, byte[] p4)
{
   javax.crypto.spec.SecretKeySpec v0_1 = new javax.crypto.spec.SecretKeySpec(p4, "HmacSHA256");
   javax.crypto.Mac v4_2 = javax.crypto.Mac.getInstance(v0_1.getAlgorithm());
   v4_2.init(v0_1);
   return com.nbc.cpc.core.network.AccessVOD.toHexString(v4_2.doFinal(p3.getBytes()));
}
~~~

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://www.nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358
