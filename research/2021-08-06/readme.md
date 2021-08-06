# August 3 2021

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

## Mahalanobis distance

- <https://wikipedia.org/wiki/Mahalanobis_distance>
- https://github.com/bitterfly/emotions/blob/master/emotions/kmeans.go
- https://pkg.go.dev/gonum.org/v1/gonum/stat#Mahalanobis
- https://stats.stackexchange.com/questions/172564

## standard deviation

The standard deviation of a random variable, sample, statistical population,
data set, or probability distribution is the square root of its variance.

- <https://wikipedia.org/wiki/Standard_deviation>
- https://github.com/golang/perf/blob/master/internal/stats/sample.go

## standard score

You could calculate Z-scores for distances

<https://wikipedia.org/wiki/Standard_score>

## variance

Using the variance of the variables and assuming that queries are in the same
distributions would probably go a long way towards a reasonable answer.

Calculate the variance for each of the variable, then scale by this (the
variables and the query), then choose based on minimum Euclidean distance. This
is a reasonable, but naive implementation.

- https://pkg.go.dev/gonum.org/v1/gonum/stat#MeanVariance
- https://pkg.go.dev/gonum.org/v1/gonum/stat#Variance
- https://wikipedia.org/wiki/Variance
