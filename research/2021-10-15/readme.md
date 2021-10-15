# October 15 2021

- https://github.com/ytdl-org/youtube-dl/issues/29191
- https://nbc.com/saturday-night-live/video/october-2-owen-wilson/9000199358

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
