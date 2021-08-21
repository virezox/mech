# Vimeo

## title and date

~~~
GET /66531465 HTTP/1.1
Host: vimeo.com
~~~

can get title and date from HTML:

~~~
[
  {
    "name": "Gnarls Barkley - Who&#039;s Gonna Save My Soul - From the Basement",
    "uploadDate": "2013-05-19T21:57:42-04:00"
  }
]
~~~

## title and URL

~~~
GET /video/66531465/config HTTP/1.1
Host: player.vimeo.com
~~~

can get title and URL from JSON:

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

- https://developer.vimeo.com/api/reference/videos
- https://github.com/silentsokolov/go-vimeo/issues/17
- https://stackoverflow.com/questions/1361149
- https://vimeo.com/66531465
