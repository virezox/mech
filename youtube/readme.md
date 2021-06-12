# YouTube

Download from YouTube

https://pkg.go.dev/github.com/89z/youtube

## Free proxy list

https://proxy.webshare.io/register

## Process

Start with a typical video:

~~~
https://www.youtube.com/watch?v=XeojXq6ySs4
~~~

Using the same ID, construct a URL like this:

~~~
https://www.youtube.com/get_video_info?eurl=https://www.youtube.com&video_id=XeojXq6ySs4
~~~

The response will be a query string, like this (edited for readability):

~~~
innertube_api_version=v1&
innertube_context_client_version=2.20210504.09.00&
player_response=%7B%22responseContext%22%3A%7B%22serviceTrackingParams%22%3A...
ps=desktop-polymer&
root_ve_type=27240&
~~~

Extract the `player_response` value. This will be a JSON object, like this:

~~~json
{
  "streamingData": {
    "adaptiveFormats": [
      {
        "itag": 137,
        "mimeType": "video/mp4; codecs=\"avc1.640020\"",
        "height": 1080,
        "signatureCipher": "s=VZVZOq0QJ8wRgIhANWm3sPF-2hbzQQGrErjQFMNmxTfALco..."
      }
    ]
  }
}
~~~

Then extract the `signatureCipher` value, this is a query string, like this:

~~~
sp=sig&
s=VZVZOq0QJ8wRgIhANWm3sPF-2hbzQQGrErjQFMNmxTfALcoZkZ4IVR1djIpAiEA8HFKix6d4B3T...&
url=https://r3---sn-q4flrnek.googlevideo.com/videoplayback%3Fexpire%3D16201927...
~~~

The `url` is the URL to the audio or video. However before you can access the
URL, you must add an entry to the query string. The new key, is the value under
`sp` above (`sig` in this case). The new value, is the value under `s` above
(`VZVZOq0QJ8wRgIhANWm3sPF-2hbzQQGrErjQFMNmxTfALcoZkZ4IVR1djIpA...` in this case).
However before you can add the new entry, you must decode the `s` value. To
decode the value, take the following steps. First, visit the original page:

~~~
https://www.youtube.com/watch?v=XeojXq6ySs4
~~~

In the source code, will be some text like this:

~~~
/s/player/3e7e4b43/player_ias.vflset/en_US/base.js
~~~

which you can turn into:

~~~
https://www.youtube.com/s/player/3e7e4b43/player_ias.vflset/en_US/base.js
~~~

In this new page, will be some code like this:

~~~js
var uy={an:function(a){a.reverse()},
gN:function(a,b){a.splice(0,b)},
J4:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c}};
vy=function(a){a=a.split("");uy.gN(a,2);uy.J4(a,47);uy.gN(a,1);uy.an(a,49);
uy.gN(a,2);uy.J4(a,4);uy.an(a,71);uy.J4(a,15);uy.J4(a,40);return a.join("")};
~~~

Take the original `s` value, and run it through this function:

~~~js
vy('_l_lOq0QJ8wRAIgc-yNc9Z4lSO2CozG4B-W9uC5zeuTATDvqHlnQaHGNmkCICsZJGbEjKDmD...')
~~~

Result will look about the same, but scrambled:

~~~
AOq0QJ8wRAIgc-ylc9Z4lSO2CozG4B-W9uC5zeuTNTDvqH_nQaHGNmkCICsZJGbEjKDmDSnKg_atTR...
~~~

Finally you can construct the resulting URL:

~~~
https://r3---sn-q4fl6nz7.googlevideo.com/videoplayback?vprv=1&
id=o-AHThxQXyxJ3jfw5EBUJeT0IJLrdQeYpMdCsCImMfbuac&
sig=AOq0QJ8wRAIgc-ylc9Z4lSO2CozG4B-W9uC5zeuTNTDvqH_nQaHGNmkCICsZJGbEjKDmDSnKg_...
~~~

## Links

- https://github.com/iawia002/annie/issues/839
- https://github.com/kkdai/youtube/issues/186
- https://golang.org/pkg/net/http#Header.WriteSubset
- https://superuser.com/questions/773719/how-do-all-of-these-save-video
