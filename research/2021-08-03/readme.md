# August 1 2021

~~~
Daft Punk - Around the world (Official Audio)
https://www.youtube.com/watch?v=dwDns8x3Jb4

Around the World
https://www.youtube.com/watch?v=Jb6gcoR266U
~~~

Lets JSON encode all the data for the entire album, so we can stop making HTTP
requests. Then we can try different sorting until we get it right. We cant use
these:

- duration
- duration size
- size

because of this:

~~~
1.36s eiN30tKzJHo The Bling Ring - Oneohtrix Point Never - Ouroboros
~~~

which leaves these:

- title
- duration title
- size title
- duration size title

## Links

- <https://wikipedia.org/wiki/Euclidean_distance>
- <https://wikipedia.org/wiki/Relative_change_and_difference>
- https://stackoverflow.com/questions/57648933/why-go-does-not-have-function
