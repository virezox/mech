# September 15 2021

## `bandcamp.com/login_cb`

This doesnt work, as it requires Captcha

## `bandcamp.com/oauth_login`

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

## `bandcamp.com/oauth_token`

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
