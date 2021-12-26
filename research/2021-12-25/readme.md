# December 25 2021

- https://github.com/ytdl-org/youtube-dl/issues/29145
- https://twitter.com/andrewbrown/status/1468949573565657094

First:

~~~
GET /1.1/statuses/show/1468949573565657094.json?tweet_mode=extended HTTP/1.1
Host: api.twitter.com
Authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
X-Guest-Token: 1474951992976019457
~~~

Then:

~~~
GET /i/api/graphql/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById?variables=%7B%22id%22%3A%221OdKrBnaEPXKX%22%2C%22isMetatagsQuery%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withBirdwatchPivots%22%3Afalse%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withReplays%22%3Atrue%2C%22withScheduledSpaces%22%3Atrue%7D HTTP/1.1
Host: twitter.com
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
~~~

Then:

~~~
GET /i/api/1.1/live_video_stream/status/28_1468947428984328193?client=web&use_syndication_guest_id=false&cookie_set_host=twitter.com HTTP/1.1
Host: twitter.com
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
~~~
