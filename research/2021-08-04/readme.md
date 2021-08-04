# August 4 2021

Distance with different units

Say I am a detective, and I want to find the suspect that most closely matches a
description:

~~~
height: 70 inch
weight: 170 lbs
~~~

and here are the suspects:

name  | height | weight
------|--------|-------
Adam  | 60     | 160
Bob   | 65     | 180
Chris | 70     | 200

I first thought about using squared euclidian distance [1], such as:

~~~
Adam = (70-60)^2 + (170-160)^2
~~~

but since the vectors are different units, I dont think this works. I also read
about L2-normalised Euclidean distance [2], but I cant figure out how to
"normalise" the values.

1. <https://wikipedia.org/wiki/Euclidean_distance>
2. <https://wikipedia.org/wiki/Cosine_similarity>

## Answer

Mahalanobis distance

https://paulrohan.medium.com/euclidean-distance-and-normalization-of-a-vector-76f7a97abd9
