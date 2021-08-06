# August 6 2021

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
