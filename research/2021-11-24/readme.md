# November 24 2021

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://www.nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358

Start with this:

~~~
POST /v2/graphql HTTP/1.1
Host: friendship.nbc.co
Content-Type: application/json

{
  "extensions": {
    "persistedQuery": {
      "sha256Hash": "c594b18a26f54ad76a6d74f94f4e3d03a48db0615dfb9115f1f894a0b918dac2"
    }
  },
  "variables": {
    "app": "nbc",
    "name": "saturday-night-live/video/october-2-owen-wilson/9000199358",
    "platform": "web",
    "type": "VIDEO",
    "userId": ""
  }
}
~~~

In the response should be this:

~~~
data	
bonanzaPage	
metadata	
mpxAccountId	"2410887629"
~~~

Then you can do this:

~~~
GET /s/NnzsPC/media/guid/2410887629/9000199358?format=preview HTTP/1.1
Host: link.theplatform.com
~~~

or:

~~~
GET /s/NnzsPC/media/guid/2410887629/9000199358?manifest=m3u HTTP/1.1
Host: link.theplatform.com
~~~

This is interesting:

~~~
GET /s/NnzsPC/media/DEFwDJCQMino?manifest=m3u&policy=188569381 HTTP/1.1
Host: link.theplatform.com
~~~

Here are the different values:

key            | value
---------------|------
`contentPid`   | `NnzsPC`       
`mediaPid`     | `DEFwDJCQMino`
`mpxAccountId` | `2410887629`
`mpxGuid`      | `9000199358`

Check this out:

~~~
data, attributes, mediaUrl pass
https://api.nbc.com/v3.14/videos/0ac67dfc-07d8-4cc4-a52f-8a98537b9c05

https://api.nbc.com/v3.14/videos?
filter[permalink]=http://www.nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358

mediaUrl fail
https://api.nbc.com/v3.14/videos/63c301b6-5bd7-4f9f-97fb-536cf7c7a9ac

404
https://api.nbc.com/v4.28.0/videos/63c301b6-5bd7-4f9f-97fb-536cf7c7a9ac
~~~

Version 3 and 4 are in the HAR file.
