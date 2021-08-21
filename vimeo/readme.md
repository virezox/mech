# Vimeo

## vimeo.com/66531465

~~~
GET /66531465 HTTP/1.1
Host: vimeo.com
~~~

title and date:

~~~html
<script>
window.vimeo.clip_page_config = {
  "clip": {
    "title": "Gnarls Barkley - Who’s Gonna Save My Soul - From the Basement",
    "uploaded_on": "2013-05-19 21:57:42"
  }
};
</script>
~~~

## api.vimeo.com/videos/66531465

~~~
GET /videos/66531465 HTTP/1.1
Host: api.vimeo.com
Authorization: jwt eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2Mjk2NjMwMD...
~~~

title, date and URL:

~~~json
{
  "files": [
    {
      "link": "https://player.vimeo.com/play/165631714?s=66531465_1629581623_054cc505a3abcab54b1b519be11bb55d&sid=8dda996379d8b7d5d6983d1983e172b25c1bc2d11629570823&oauth2_token_id="
    }
  ],
  "name": "Gnarls Barkley - Who's Gonna Save My Soul - From the Basement",
  "release_time": "2013-05-20T01:57:42+00:00"
}
~~~

## player.vimeo.com/video/66531465/config

~~~
GET /video/66531465/config HTTP/1.1
Host: player.vimeo.com
~~~

title and URL:

~~~json
{
  "request": {
    "files": {
      "progressive": [
        {
          "url": "https://vod-progressive.akamaized.net/exp=1629581401~acl=%2A%2F165631714.mp4%2A~hmac=8f2c0423a3211987f7a35b164037e8a764c0bcdefbb7b3ed7a9f584cc11e7efc/vimeo-prod-skyfire-std-us/01/3306/2/66531465/165631714.mp4"
        }
      ]
    }
  },
  "video": {
    "title": "Gnarls Barkley - Who's Gonna Save My Soul - From the Basement"
  }
}
~~~

## vimeo.com/66531465?action=load_download_config

~~~
GET /66531465?action=load_download_config HTTP/1.1
Host: vimeo.com
X-Requested-With: XMLHttpRequest
~~~

URL:

~~~json
{
  "files": [
    {
      "download_url": "https://player.vimeo.com/play/165631714?s=66531465_1629578449_6b3d64d53abdf46f091a63188e973c2e&loc=external&context=Vimeo%5CController%5CClipController.main&download=1"
    }
  ]
}
~~~

- https://developer.vimeo.com/api/reference/videos
- https://github.com/silentsokolov/go-vimeo/issues/17
- https://stackoverflow.com/questions/1361149
- https://vimeo.com/66531465
