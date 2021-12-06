# NBC

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
