# December 17 2021

## How to get request body?

- https://github.com/mcrute/pydora/blob/master/pandora/transport.py
- https://github.com/nlowe/mousiki/blob/master/pandora/api/legacy.go

## How to get `auth_token`?

~~~
POST /services/json/?method=auth.partnerLogin HTTP/1.1
Host: android-tuner.pandora.com

{"username":"android","password":"AC7IBG09A3DTSYM4R41UJWL07VLN8JI7"}
~~~

## How to get `user_id`?

~~~
POST /services/json/?partner_id=42&auth_token=VAzrFQTtsy3BSe9w6TEEEejwPulcDkRLMA&method=auth.userLogin HTTP/1.1
Host: android-tuner.pandora.com

cacc58f238d86eeb823381388dc1e3955f711602e6979219601f028e4c8497f873e200af9bcf7...
~~~

--------------------------------------------------------------------------------

## How to get `audioUrl`?

~~~
POST /services/json/?partner_id=42&user_id=1901383005&auth_token=VIUUlTskRgDbBySduQEY343ZwoyVPZ1yLQeapGMYNSBZXIt8dLFCIA8w%3D%3D&method=onDemand.getAudioPlaybackInfo HTTP/1.1
Host: android-tuner.pandora.com
Content-Length: 1072

0b5b3f806abef32879a802a0749e65e9bea1623d9ff53d4c47e1db0a11135f61e8de2089919ef...
~~~
