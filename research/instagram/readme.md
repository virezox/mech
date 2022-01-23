# Instagram

Old API returns different results in different countries:

https://github.com/89z/mech/issues/13#issuecomment-1019303135

New API does not return comments:

https://github.com/89z/mech/issues/13#issuecomment-1019300734

Not sure how this one works:

https://github.com/honkkki/goins/issues/1

Looks like comments has a different request:

- https://github.com/codeuniversity/smag-mvp/blob/master/insta/scraper/comments/comments-scraper.go#L191
- https://github.com/codeuniversity/smag-mvp/blob/master/insta/scraper/comments/comments-scraper.go#L234

What does Android client do?

~~~
com.instagram.android
~~~

Info endpoint:

~~~
GET /api/v1/media/2755022163816059161/info/ HTTP/2.0
Host: i.instagram.com
user-agent: Instagram 206.1.0.34.121 Android
Authorization: Bearer IGT:2:eyJkc191c2VyX2lkIjoiNDkzNzgxNzEzMzQiLCJzZXNzaW9ua...
~~~

Comment endpoint:

~~~
GET /api/v1/media/2755652849306967814/comments/ HTTP/2.0
Host: i.instagram.com
user-agent: Instagram 215.0.0.27.359 Android
Authorization: Bearer IGT:2:eyJkc191c2VyX2lkIjoiNDkzNzgxNzEzMzQiLCJzZXNzaW9ua...
~~~

With the Android client, this would be the struct:

~~~go
type info struct {
   Items []struct {
      Carousel_Media []struct {
         Image_Versions2 struct {
            Candidates []struct {
               URL string
            }
         }
      }
   }
}
~~~

With the Web client, this would be the struct:

~~~go
type info struct {
   Data struct {
      Shortcode_Media struct {
         Display_URL string
      }
   }
}
~~~
