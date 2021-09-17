# September 16 2021

## /api/album

- <http://bandcamp.com/api/album/2/info?key=veidihundr&album_id=79940049>
- <http://bandcamp.com/api/album/1/info?key=veidihundr&album_id=79940049>

## /api/band

- <http://bandcamp.com/api/band/3/info?key=veidihundr&band_id=2853020814>
- <http://bandcamp.com/api/band/2/info?key=veidihundr&band_id=2853020814>
- <http://bandcamp.com/api/band/1/info?key=veidihundr&band_url=duststoredigital.com>
- <http://bandcamp.com/api/band/1/info?key=veidihundr&band_id=2853020814>

## /api/mobile

~~~
POST /api/mobile/22/band_details HTTP/1.1
Host: bandcamp.com

{"band_id":"1196681540"}
~~~

- <http://bandcamp.com/api/mobile/24/band_details?band_id=2853020814>
- <http://bandcamp.com/api/mobile/24/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/23/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/22/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/21/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/20/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/18/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/17/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/16/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/15/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/14/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/13/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/12/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/11/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/10/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/9/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>
- <http://bandcamp.com/api/mobile/8/tralbum_details?tralbum_type=t&band_id=2853020814&tralbum_id=714939036>

## /api/track

- <http://bandcamp.com/api/track/3/info?key=veidihundr&track_id=714939036>
- <http://bandcamp.com/api/track/1/info?key=veidihundr&track_id=714939036>

## /api/url

~~~
POST /api/url/2/info HTTP/1.1
Host: bandcamp.com

{"key":"veidihundr","url":"duststoredigital.com"}
~~~

- http://bandcamp.com/api/url/2/info?key=veidihundr&url=duststoredigital.com
- http://bandcamp.com/api/url/1/info?key=veidihundr&url=duststoredigital.com/track/sudden-intake
- http://bandcamp.com/api/url/1/info?key=veidihundr&url=duststoredigital.com/album/silenced
- http://bandcamp.com/api/url/1/info?key=veidihundr&url=duststoredigital.com

## key

~~~
thrjozkaskhjastaurrtygitylpt
throtaudvinroftignmarkreina
ullrettkalladrhampa
veidihundr
~~~

## /login\_cb

This doesnt work, as it requires Captcha

## /oauth\_login

This works:

~~~
POST /oauth_login HTTP/1.1
host: bandcamp.com
x-bandcamp-dm: 8f38339869c3003e9f1c8b1c13fe48530f74e3c6

client_id=134
client_secret=1myK12VeCL3dWl9o%2FncV2VyUUbOJuNPVJK6bZZJxHvk%3D
grant_type=password
password=PASSWORD
username=4095486538
username_is_user_id=1
~~~

We can get `x-bandcamp-dm` from Android, but its only good for three minutes. I
found an implementation online, but it seems BandCamp has changed the algorithm:

https://github.com/the-eater/camp-collective/issues/5

## /oauth\_token

We can try this:

~~~
POST /oauth_token HTTP/1.1
host: bandcamp.com

client_id=134&
client_secret=1myK12VeCL3dWl9o%2FncV2VyUUbOJuNPVJK6bZZJxHvk%3D&
grant_type=client_credentials
~~~

Result:

~~~
Only third-party clients can use client_credentials
~~~
