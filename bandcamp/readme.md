# Bandcamp

https://stackoverflow.com/questions/65085046/how-do-i-pull-album-art

## Mobile

~~~
POST /api/mobile/22/band_details HTTP/1.1
Host: bandcamp.com

{"band_id":"1196681540"}
~~~

- <http://bandcamp.com/api/mobile/24/band_details?band_id=2853020814>
- <http://bandcamp.com/api/mobile/24/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>

## Track

<http://bandcamp.com/api/track/3/info?key=veidihundr&track_id=714939036>

## URL

- http://bandcamp.com/api/url/2/info?key=veidihundr&url=duststoredigital.com
- http://bandcamp.com/api/url/1/info?key=veidihundr&url=duststoredigital.com/album/silenced
- http://bandcamp.com/api/url/1/info?key=veidihundr&url=duststoredigital.com/track/sudden-intake
