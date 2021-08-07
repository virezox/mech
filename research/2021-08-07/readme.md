# August 7 2021

~~~
POST /youtubei/v1/search HTTP/1.1
Host: www.youtube.com
x-goog-api-key: AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8

{"context":{"client":{"clientName":"MWEB", "clientVersion":"2.19700101"}},
"params":"EgIQAQ", "query":"daft punk topic around the world"}
~~~

Compare:

~~~
Daft Punk - Around the world (Official Audio)
https://www.youtube.com/watch?v=dwDns8x3Jb4

Around the World
https://www.youtube.com/watch?v=Jb6gcoR266U
~~~

We should consider the natural sort order as a factor. We cant use these:

~~~
[duration]
[duration size]
[size]
~~~

because of this:

~~~
1.36s eiN30tKzJHo The Bling Ring - Oneohtrix Point Never - Ouroboros
~~~

we cant use this:

~~~
[index]
~~~

because of this:

~~~
B3szYRzZqp4 oneohtrix point never - describing bodies
~~~

which leaves these:

~~~
[index duration]
[index size]
[index duration size]
[title]
[index title]
[duration title]
[index duration title]
[size title]
[index size title]
[duration size title]
[index duration size title]
~~~

- <https://wikipedia.org/wiki/Mahalanobis_distance>
- https://github.com/bitterfly/emotions/blob/master/emotions/kmeans.go
- https://github.com/golang/perf/blob/master/internal/stats/sample.go
