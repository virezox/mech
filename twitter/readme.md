# Twitter

These requests seem to do the trick:

~~~
POST /1.1/guest/activate.json HTTP/1.1
Host: api.twitter.com
Authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA

GET /1.1/statuses/show/1577759019724333058.json?tweet_mode=extended HTTP/1.1
Host: api.twitter.com
Authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
X-Guest-Token: 1578461005738397697
~~~
