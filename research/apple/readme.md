# Apple TV

Using this video:

https://tv.apple.com/us/show/ted-lasso/umc.cmc.vtoh0mn0xn7t3c643xqonfzy

Download encrypted media:

~~~
ffmpeg -i `
'https://play.itunes.apple.com/WebObjects/MZPlay.woa/hls/subscription/stream/playlist.m3u8?cc=US&g=230&cdn=vod-ap2-aoc.tv.apple.com&a=1484589502&p=461374806&st=1821682575&a=1625486472&p=461370051&st=1821645191&a=1622268591&p=461372307&st=1821659224&a=1613450761&p=461480166&st=1822491467&a=1522961240&p=377679659&st=1490983814&a=1524197777&p=368330428&st=1449517784&a=1524197722&p=368330432&st=1449518254&a=1524198082&p=368330370&st=1449517587&a=1525078430&p=368329706&st=1449509871&a=1524197604&p=368330236&st=1449518699&a=1524197554&p=368330322&st=1449518442&a=1524197773&p=368330253&st=1449517917&a=1539152595&p=368283705&st=1449199894' `
-c copy enc.mp4
~~~

Next we need the Widevine [1] PSSH from the HLS file:

~~~xml
#EXT-X-KEY:URI="data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgu8rYzAgICAgICBI88aJmwY=",
KEYFORMAT="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed",KEYFORMATVERSIONS="1",METHOD=SAMPLE-AES
~~~

Now go back to the video page, and you should see a request like this:

~~~
FIXME
POST https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6245817279001/aff72f30-a546-4e44-99bf-d630f41c1adf/c0e598b2-47fa-4435-9029-9d5ef47da32c?fastly_token=NjI5MDMxMTdfZWZjNDViYzJmN2VhNTRiM2M3ZWZhM2NjMWE2MjNlMzg5ODJjYmFiMjY3ODBiMGNmYjBkODIxMjQ0YWJlMjFkNQ%3D%3D HTTP/2.0
bcov-auth: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJlNTllNjI5Ni1iZGM0LTRiMzItYThkMy1iMTIyMjI5MDU0MzQiLCJhY2NpZCI6IjYyNDU4MTcyNzkwMDEiLCJwcmlkIjoiODJkMWI0MmEtMWQ0Mi00ZGZiLTg2MmUtNTNmZDhkNWU2NmE4IiwiY29uaWQiOiI2MjYzNDM1MTA4MDAxIiwiY2xpbWl0Ijo0LCJleHAiOjE2NTM2Nzc5ODYsImNiZWgiOiJCTE9DS19ORVdfVVNFUiIsImlhdCI6MTY1MzU5MTU4NiwianRpIjoiNzE4MzI4OWItMWEyNy00YjU4LWFiMGYtZmEyMTRkOWZhMzFiIiwic2lkIjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NDsgcnY6ODguMCkgR2Vja28vMjAxMDAxMDEgRmlyZWZveC84OC4wIC0gMjE3MDk3MDAwOSJ9.MuTrQMNdfxuqn5rD7ydbc5yuzfnL9DFsPlbHXV5Xsg8-1eQXJmZ7HQeBO-PMps5eF0SoG4FM1B--uif0Hae_D2aqUmxZoGKe8xQDxCnMYaTzwtmpMZYZfiPtAyUyXY80DMgaOxcH8o680MVvRuV-98dm6boHfx-yvAVNzIVJXNDGYGBoKqL5bHm7PmrEXXaRXGZTf-8Ep09eSqpou_eQeiqmFiZhPtNEXXuFNHB-E68SqZIIv2nfkRH4stwHzUpD6fjyuMuMnlRaRbEY_955I-8lCy68QI8uxCIGyBaPAb6LpLbtptpQtdCGpUumuR2Qt-yWwqakTKa-LAnUAe7umg
~~~

Now go to Get Widevine Keys, and enter the information from above:

~~~
PSSH:
AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQwOWYskf6RDWQKZ1e9H2jLBoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=

License:
https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6245817279001/aff72f30-a546-4e44-99bf-d630f41c1adf/c0e598b2-47fa-4435-9029-9d5ef47da32c?fastly_token=NjI5MDMxMTdfZWZjNDViYzJmN2VhNTRiM2M3ZWZhM2NjMWE2MjNlMzg5ODJjYmFiMjY3ODBiMGNmYjBkODIxMjQ0YWJlMjFkNQ%3D%3D

Headers:
bcov-auth: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJlNTllNjI5Ni1iZGM0LTRiMzItYThkMy1iMTIyMjI5MDU0MzQiLCJhY2NpZCI6IjYyNDU4MTcyNzkwMDEiLCJwcmlkIjoiODJkMWI0MmEtMWQ0Mi00ZGZiLTg2MmUtNTNmZDhkNWU2NmE4IiwiY29uaWQiOiI2MjYzNDM1MTA4MDAxIiwiY2xpbWl0Ijo0LCJleHAiOjE2NTM2Nzc5ODYsImNiZWgiOiJCTE9DS19ORVdfVVNFUiIsImlhdCI6MTY1MzU5MTU4NiwianRpIjoiNzE4MzI4OWItMWEyNy00YjU4LWFiMGYtZmEyMTRkOWZhMzFiIiwic2lkIjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NDsgcnY6ODguMCkgR2Vja28vMjAxMDAxMDEgRmlyZWZveC84OC4wIC0gMjE3MDk3MDAwOSJ9.MuTrQMNdfxuqn5rD7ydbc5yuzfnL9DFsPlbHXV5Xsg8-1eQXJmZ7HQeBO-PMps5eF0SoG4FM1B--uif0Hae_D2aqUmxZoGKe8xQDxCnMYaTzwtmpMZYZfiPtAyUyXY80DMgaOxcH8o680MVvRuV-98dm6boHfx-yvAVNzIVJXNDGYGBoKqL5bHm7PmrEXXaRXGZTf-8Ep09eSqpou_eQeiqmFiZhPtNEXXuFNHB-E68SqZIIv2nfkRH4stwHzUpD6fjyuMuMnlRaRbEY_955I-8lCy68QI8uxCIGyBaPAb6LpLbtptpQtdCGpUumuR2Qt-yWwqakTKa-LAnUAe7umg
~~~

You should get a result like this:

~~~
c0e598b247fa443590299d5ef47da32c:a66a5603545ad206c1a78e160a6710b1
~~~

Finally, you can decrypt [2] the media:

~~~
mp4decrypt --key 1:a66a5603545ad206c1a78e160a6710b1 enc.mp4 dec.mp4
~~~

1. <https://dashif.org/identifiers/content_protection>
2. https://bento4.com/downloads
