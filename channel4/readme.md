# Channel 4

- <https://dashif.org/identifiers/content_protection>
- https://bento4.com/downloads
- https://tools.axinom.com/generators/PsshBox

Using this video:

https://www.channel4.com/programmes/frasier/on-demand/18926-001

Download the MPD:

~~~
yt-dlp -o enc.mp4 -f video=501712 --allow-unplayable-formats `
'https://ak-jos-c4assets-com.akamaized.net/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13500989517136472855&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650915617&e=600&st=wk9eWkEeWVJzjYlveA4ysHMimJgJXGD5oCUjpu-HGqU'
~~~

Decrypt the media:

~~~
mp4decrypt --key 00000000000000000000000004246624:1da9d79c4c6a... enc.mp4 dec.mp4
~~~
