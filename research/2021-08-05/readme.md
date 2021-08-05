# August 5 2021

You could calculate Z-scores for distances, but still...

Using the variance of the variables and assuming that queries are in the same
distributions would probably go a long way towards a reasonable answer.

@Steven Penny Calculate the variance for each of the variable, then scale by
this (the variables and the query), then choose based on minimum Euclidean
distance. This is a reasonable, but naive implementation.

- <https://wikipedia.org/wiki/Standard_score>
- https://paulrohan.medium.com/euclidean-distance-and-normalization-of-a-vector-76f7a97abd9
- https://stats.stackexchange.com/questions/539156/distance-with-different-units
- https://wikipedia.org/wiki/Variance
- https://www.machinelearningplus.com/statistics/mahalanobis-distance
