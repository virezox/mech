# Roku

## Android client

just a remote control:

https://play.google.com/store/apps/details?id=com.roku.remote

## Web free

Clear:

https://therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76

Protected:

https://therokuchannel.roku.com/watch/934102524ea6524eb7d1faa29233695a

## Web premium

https://therokuchannel.roku.com/watch/32c95b576307502b98f7fe32c4aa0a22

We can create free account, but login uses CAPTCHA:

~~~
POST https://my.roku.com/api/auth/1/login HTTP/2.0
csrf-token: sVzM79JV-kKC0kN2Jlz_PlI6vHLZ3NhqOqFk
content-type: application/json
cookie: _csrf=LHBM-wxg8GRExB8JboGxeJCC

{"email":"EMAIL","password":"PASSWORD","rememberMe":false}

{
  "error": "captcha"
}
~~~
