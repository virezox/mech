# Facebook

OK I got it. First you do this:

~~~
GET /login.php HTTP/1.1
Host: m.facebook.com
~~~

in the response header should be this:

~~~
set-cookie: datr=FxZfYrBtEgZeGMkKP-rsd6RH
~~~

in the response body should be this:

~~~html
<form method="post"
action="/login/device-based/regular/login/?refsrc=deprecated&amp;lwv=100&amp;refid=9"
class="u v" id="login_form" novalidate="1"><input type="hidden"
name="lsd" value="AVp_Z4_s3ms" autocomplete="off" />
~~~

use both of those to make a new request:

~~~
POST /login/device-based/regular/login/ HTTP/1.1
Host: m.facebook.com
Content-Type: application/x-www-form-urlencoded
Cookie: datr=MxJfYrG9o7FrP2k9iHQ2uhm9

email=YOUR_EMAIL&lsd=AVoiuEvJgiA&pass=YOUR_PASSWORD
~~~

then you can use response cookie to get gated videos.
