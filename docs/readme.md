# Docs

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. `yt-dlp --proxy 127.0.0.1:8080 --no-check-certificate`
2. try navigating to the target page from home screen, instead of going directly
   to page
3. check media page for JSON requests
4. check HAR file
5. check HTML
6. check JavaScript
7. MITM Proxy APK for JSON requests
8. APK tool
9. JADX

## How to update topics?

~~~
PUT /repos/89z/mech/topics HTTP/1.1
Host: api.github.com
Authorization: Basic ODl6OmE1NzYxMjZlNzVlZjZiY2Y5ZDljNzEyZWIyN2RmZjFmOGFhZmQ...

{"names":[
   "youtube", "instagram"
]}
~~~
