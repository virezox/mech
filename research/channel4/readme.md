# Channel 4

https://forum.videohelp.com/threads/405455-Need-help-downloading-from-Channel-4

## Android client

https://play.google.com/store/apps/details?id=com.channel4.ondemand

~~~
POST /online/v2/auth/token?client=android-mod HTTP/1.1
Connection: Keep-Alive
Content-Type: application/x-www-form-urlencoded
Host: api.channel4.com
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC) C4oD_Android/9.4.5 (uid:73978094-edd6-40f9-abe0-e4d487154024; tid:-; did:Google_Android SDK built for x86_24;)
X-C4-App-Version: android_app:9.4.5
X-C4-Device-Name: Google Android SDK built for x86 (sdk_google_phone_x86)
X-C4-Device-Type: mobile
X-C4-Optimizely-Datafile: 1693
X-C4-Platform-Name: android
X-Correlation-Id: ANDROID-f32aa536-e08a-431d-9b5f-5e3f0f0f5807
X-Forwarded-For: 25.0.0.0

grant_type=client_credentials&client_id=36UUCt98VMQvBAgQ27Au8zGHl31N9LQ1&client_secret=JYswyHvGe62VlikW

HTTP/1.1 403 Forbidden
~~~

## Web client

Using this video:

https://www.channel4.com/programmes/frasier/on-demand/18926-001

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f video=501712 --allow-unplayable-formats `
'https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13500989517136472855&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650915617&e=600&st=wk9eWkEeWVJzjYlveA4ysHMimJgJXGD5oCUjpu-HGqU'
~~~

Next we need the KID from the MPD file:

~~~xml
<ContentProtection
   schemeIdUri="urn:mpeg:dash:mp4protection:2011"
   value="cenc"
   cenc:default_KID="00000000-0000-0000-0000-000003323603">
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
POST /wvlicenceproxy-service/widevine/acquire HTTP/1.1
Host: c4.eme.lp.aws.redbeemedia.com

{
  "request_id": 5273616,
  "token": "QVJDUm94UXVYVLYcO0q52OVAEcZqaxNZzLYaHy2ePGpGwsm0K4k37r4E55bv8G-i0C7UfUVtzmPxRb_XYp1hnXZGqPjdIH8FhOKQ7I5Asa-FzAkKjiBfvL9EDSdL5z-dCbznzXwHkIbczQP9B8VMypXxQRpxfT_x",
  "video": {
    "type": "ondemand",
    "url": "https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13500145088601714632&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650917217&e=600&st=EQehVvoFEAq_hXMQHqYy9IZhJpLINev_xyrX_FuMsVQ"
  },
  "message": "CAQ="
}
~~~

Next download the Channel 4 script [1], and update using the request body above:

~~~py
self.json_payloads = {
  "request_id": 5273616,
  "token": "QVJDUm94UXVYVLYcO0q52OVAEcZqaxNZzLYaHy2ePGpGwsm0K4k37r4E55bv8G-i0C7UfUVtzmPxRb_XYp1hnXZGqPjdIH8FhOKQ7I5Asa-FzAkKjiBfvL9EDSdL5z-dCbznzXwHkIbczQP9B8VMypXxQRpxfT_x",
  "video": {
    "type": "ondemand",
    "url": "https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13500145088601714632&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650917217&e=600&st=EQehVvoFEAq_hXMQHqYy9IZhJpLINev_xyrX_FuMsVQ"
  },
  "message": "CAQ="
}
~~~

Run the script, and you should get a result like this:

~~~
--key 00000000000000000000000004246624:1da9d79c4c6adbd5a0658e712ccf7f7e
~~~

Finally, you can decrypt [2] the media:

~~~
mp4decrypt --key 00000000000000000000000004246624:1da9d79c4c6a... enc.mp4 dec.mp4
~~~

1. https://getwvkeys.cc/scripts
2. https://bento4.com/downloads
