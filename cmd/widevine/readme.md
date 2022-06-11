# Widevine

Example:

~~~
widevine `
-b 63363432-3534-6364-6237-656465336530 `
-a https://lic.drmtoday.com/license-proxy-widevine/cenc/?specConform=true `
-h x-dt-custom-data:eyJ1c2VySWQiOiJhbm9uaW0iLCJzZXNzaW9uSWQiOiJrRGJZZUlYeVM4VjNtNXVRMjRtU0F6cThXdDkiLCJtZXJjaGFudCI6ImNkYSJ9
~~~

## How to get key id?

If you look in the MPD file, you should see it:

~~~xml
<ContentProtection value="cenc" schemeIdUri="urn:mpeg:dash:mp4protection:2011"
cenc:default_KID="63363432-3534-6364-6237-656465336530"/>
~~~

## How to get license URL and header?

If you watch the requests on the video page, you should see a request similar
to this:

~~~
POST https://lic.drmtoday.com/license-proxy-widevine/cenc/?specConform=true HTTP/1.1
Accept-Language: en-US,en;q=0.5
Accept: */*
Connection: keep-alive
DNT: 1
Origin: https://www.cda.pl
Referer: https://www.cda.pl/video/391634853/vfilm
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
x-dt-custom-data: eyJ1c2VySWQiOiJhbm9uaW0iLCJzZXNzaW9uSWQiOiJrRGJZZUlYeVM4VjNtNXVRMjRtU0F6cThXdDkiLCJtZXJjaGFudCI6ImNkYSJ9
~~~

Sometimes, only the URL is needed, in which case you can omit the header option.
It depends on the license server.
