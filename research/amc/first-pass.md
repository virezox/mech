# AMC

Using this video:

https://amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152

Download the MPD:

~~~
yt-dlp -o enc.mp4 --allow-unplayable-formats `
-f edfa136e-6e76-4043-bec6-e41e772dff2f `
https://ssaimanifest.prod.boltdns.net/us-east-1/playback/once/v1/dash/live-timeline/bccenc/6245817279001/74263065-1285-4d00-842d-73a66554716a/aff72f30-a546-4e44-99bf-d630f41c1adf/be9bd8c1-b573-46ac-ac9f-4f35c24243e3/manifest.mpd?bc_token=NjI5MDMxMTdfZWI0NjlmNTlkMjNjMmFkZGUyMTkyY2I2M2Y3MDAxYWMzNjMzNWM3NTI4MTU5NjI0NjQxYTNjNzFmODM0YzI2Zg%3D%3D
~~~

Next we need the Widevine [1] PSSH from the MPD file:

~~~xml
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed" bc:licenseAcquisitionUrl="https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6245817279001/aff72f30-a546-4e44-99bf-d630f41c1adf/c0e598b2-47fa-4435-9029-9d5ef47da32c?fastly_token=NjI5MDMzMDBfZjM4ODViZWY1ZWZmOGUyODQ3NzBiNTQ1ZDAxY2ZjMDc4OWE3YjlkYmI4MGM3ZmRlYzIwNzdjNzFlZWNmNzQ5NA%3D%3D">
   <cenc:pssh>
   AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQwOWYskf6RDWQKZ1e9H2jLBoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=
   </cenc:pssh>
</ContentProtection>
~~~

Now go back to the video page, and you should see a request like this:

~~~
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
