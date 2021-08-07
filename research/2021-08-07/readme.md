# August 7 2021

Using the variance of the variables and assuming that queries are in the same
distributions would probably go a long way towards a reasonable answer.

Calculate the variance for each of the variable, then scale by this (the
variables and the query), then choose based on minimum Euclidean distance. This
is a reasonable, but naive implementation.

- <https://wikipedia.org/wiki/Mahalanobis_distance>
- https://github.com/bitterfly/emotions/blob/master/emotions/kmeans.go
- https://github.com/golang/perf/blob/master/internal/stats/sample.go
- https://stats.stackexchange.com/questions/172564
- https://wikipedia.org/wiki/Variance
