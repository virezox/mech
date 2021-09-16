bandcamp: bad key

Using this key:

https://github.com/tomahawk-player/tomahawk-resolvers/blob/7f827bbe410ccfdb0446f7d6a91acc2199c9cc8d/bandcamp/content/contents/code/bandcamp.js#L21

If you try to make a request such as this:

http://bandcamp.com/api/url/2/info?key=inganwbxhyy&url=duststoredigital.com

You get this result:

~~~json
{
  "error_message": "bad key",
  "error": true
}
~~~
