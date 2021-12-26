# December 25 2021

- https://github.com/ytdl-org/youtube-dl/issues/29145
- https://twitter.com/andrewbrown/status/1468949573565657094

First:

~~~
GET /andrewbrown/status/1468949573565657094 HTTP/1.1
Host: twitter.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Cookie: G_ENABLED_IDPS=google; _ga=GA1.2.1632657088.1631928682; kdt=ToSEFNCXkLCovFL8vrBKCAOoQpu3TedvePSAM89K; dnt=1; personalization_id="v1_3Jp0sUHoSZHukx+6Hhdl+Q=="; guest_id=v1%3A163483047089284027; guest_id_marketing=v1%3A163483047089284027; guest_id_ads=v1%3A163483047089284027; external_referer=padhuUp37zj9xuUOXCNFvOO7FrlI6ZSWG30xkU1yfXM4z9yB11B5mkHAtNWWp236lqUY4iig%2F2Q%3D|0|8e8t2xd8A2w%3D; ct0=fa99a0c003cc026aee2b1585264a5b4a; gt=1474894925598703618; _gid=GA1.2.552981454.1640469200; lang=en; _twitter_sess=BAh7CSIKZmxhc2hJQzonQWN0aW9uQ29udHJvbGxlcjo6Rmxhc2g6OkZsYXNo%250ASGFzaHsABjoKQHVzZWR7ADoPY3JlYXRlZF9hdGwrCBkLp9B9AToMY3NyZl9p%250AZCIlNzIwMDA5NzQyMDNhOGE0ZTJmMDgzMzdjMDYxMzUyNTM6B2lkIiU4OWU1%250AMTAwZDQ2MDZhNWQ2NTRmMGRjZDZhNGY0YzEyNA%253D%253D--f0ff3763641acbd79caddb606c523dd3b7b27d36
Upgrade-Insecure-Requests: 1
Pragma: no-cache
Cache-Control: no-cache
~~~

Then:

~~~
GET /i/api/graphql/MwoNOssr8CR7CxUWbBQO9w/TweetDetail?variables=%7B%22focalTweetId%22%3A%221468949573565657094%22%2C%22with_rux_injections%22%3Afalse%2C%22includePromotedContent%22%3Atrue%2C%22withCommunity%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withTweetQuoteCount%22%3Atrue%2C%22withBirdwatchNotes%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withBirdwatchPivots%22%3Afalse%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Afalse%7D HTTP/1.1
Host: twitter.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://twitter.com/andrewbrown/status/1468949573565657094
content-type: application/json
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
x-twitter-client-language: en
x-twitter-active-user: yes
x-csrf-token: fa99a0c003cc026aee2b1585264a5b4a
Connection: keep-alive
Cookie: G_ENABLED_IDPS=google; _ga=GA1.2.1632657088.1631928682; kdt=ToSEFNCXkLCovFL8vrBKCAOoQpu3TedvePSAM89K; dnt=1; personalization_id="v1_3Jp0sUHoSZHukx+6Hhdl+Q=="; guest_id=v1%3A163483047089284027; guest_id_marketing=v1%3A163483047089284027; guest_id_ads=v1%3A163483047089284027; external_referer=padhuUp37zj9xuUOXCNFvOO7FrlI6ZSWG30xkU1yfXM4z9yB11B5mkHAtNWWp236lqUY4iig%2F2Q%3D|0|8e8t2xd8A2w%3D; ct0=fa99a0c003cc026aee2b1585264a5b4a; gt=1474894925598703618; _gid=GA1.2.552981454.1640469200; lang=en; _twitter_sess=BAh7CSIKZmxhc2hJQzonQWN0aW9uQ29udHJvbGxlcjo6Rmxhc2g6OkZsYXNo%250ASGFzaHsABjoKQHVzZWR7ADoPY3JlYXRlZF9hdGwrCBkLp9B9AToMY3NyZl9p%250AZCIlNzIwMDA5NzQyMDNhOGE0ZTJmMDgzMzdjMDYxMzUyNTM6B2lkIiU4OWU1%250AMTAwZDQ2MDZhNWQ2NTRmMGRjZDZhNGY0YzEyNA%253D%253D--f0ff3763641acbd79caddb606c523dd3b7b27d36
Pragma: no-cache
Cache-Control: no-cache
~~~

Then:

~~~
GET /i/api/graphql/Uv5R_-Chxbn1FEkyUkSW2w/AudioSpaceById?variables=%7B%22id%22%3A%221OdKrBnaEPXKX%22%2C%22isMetatagsQuery%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withBirdwatchPivots%22%3Afalse%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withReplays%22%3Atrue%2C%22withScheduledSpaces%22%3Atrue%7D HTTP/1.1
Host: twitter.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://twitter.com/andrewbrown/status/1468949573565657094
content-type: application/json
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
x-twitter-client-language: en
x-twitter-active-user: yes
x-csrf-token: fa99a0c003cc026aee2b1585264a5b4a
Connection: keep-alive
Cookie: G_ENABLED_IDPS=google; _ga=GA1.2.1632657088.1631928682; kdt=ToSEFNCXkLCovFL8vrBKCAOoQpu3TedvePSAM89K; dnt=1; personalization_id="v1_3Jp0sUHoSZHukx+6Hhdl+Q=="; guest_id=v1%3A163483047089284027; guest_id_marketing=v1%3A163483047089284027; guest_id_ads=v1%3A163483047089284027; external_referer=padhuUp37zj9xuUOXCNFvOO7FrlI6ZSWG30xkU1yfXM4z9yB11B5mkHAtNWWp236lqUY4iig%2F2Q%3D|0|8e8t2xd8A2w%3D; ct0=fa99a0c003cc026aee2b1585264a5b4a; gt=1474894925598703618; _gid=GA1.2.552981454.1640469200; lang=en; _twitter_sess=BAh7CSIKZmxhc2hJQzonQWN0aW9uQ29udHJvbGxlcjo6Rmxhc2g6OkZsYXNo%250ASGFzaHsABjoKQHVzZWR7ADoPY3JlYXRlZF9hdGwrCBkLp9B9AToMY3NyZl9p%250AZCIlNzIwMDA5NzQyMDNhOGE0ZTJmMDgzMzdjMDYxMzUyNTM6B2lkIiU4OWU1%250AMTAwZDQ2MDZhNWQ2NTRmMGRjZDZhNGY0YzEyNA%253D%253D--f0ff3763641acbd79caddb606c523dd3b7b27d36
Pragma: no-cache
Cache-Control: no-cache
~~~

Then:

~~~
GET /i/api/1.1/live_video_stream/status/28_1468947428984328193?client=web&use_syndication_guest_id=false&cookie_set_host=twitter.com HTTP/1.1
Host: twitter.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://twitter.com/i/spaces/1OdKrBnaEPXKX
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
x-twitter-client-language: en
x-twitter-active-user: yes
x-csrf-token: fa99a0c003cc026aee2b1585264a5b4a
Connection: keep-alive
Cookie: G_ENABLED_IDPS=google; _ga=GA1.2.1632657088.1631928682; kdt=ToSEFNCXkLCovFL8vrBKCAOoQpu3TedvePSAM89K; dnt=1; personalization_id="v1_3Jp0sUHoSZHukx+6Hhdl+Q=="; guest_id=v1%3A163483047089284027; guest_id_marketing=v1%3A163483047089284027; guest_id_ads=v1%3A163483047089284027; external_referer=padhuUp37zj9xuUOXCNFvOO7FrlI6ZSWG30xkU1yfXM4z9yB11B5mkHAtNWWp236lqUY4iig%2F2Q%3D|0|8e8t2xd8A2w%3D; ct0=fa99a0c003cc026aee2b1585264a5b4a; gt=1474894925598703618; _gid=GA1.2.552981454.1640469200; lang=en; _twitter_sess=BAh7CSIKZmxhc2hJQzonQWN0aW9uQ29udHJvbGxlcjo6Rmxhc2g6OkZsYXNo%250ASGFzaHsABjoKQHVzZWR7ADoPY3JlYXRlZF9hdGwrCBkLp9B9AToMY3NyZl9p%250AZCIlNzIwMDA5NzQyMDNhOGE0ZTJmMDgzMzdjMDYxMzUyNTM6B2lkIiU4OWU1%250AMTAwZDQ2MDZhNWQ2NTRmMGRjZDZhNGY0YzEyNA%253D%253D--f0ff3763641acbd79caddb606c523dd3b7b27d36
Pragma: no-cache
Cache-Control: no-cache
~~~

Then:

~~~
GET /Transcoding/v1/hls/JVkxIZDjXFVW1HNMy76mnF384DlHI37k_rZp1Sk3P3NCWoLioMbH6lEo4QPm-ogfeNGrYpfSRkHZBx2lJgWkIw/non_transcode/us-east-1/periscope-replay-direct-prod-us-east-1-public/audio-space/playlist_16807240958558906262.m3u8?type=replay HTTP/1.1
Host: prod-fastly-us-east-1.video.pscp.tv
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://twitter.com/
Origin: https://twitter.com
Connection: keep-alive
Pragma: no-cache
Cache-Control: no-cache
~~~
