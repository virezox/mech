# July 26 2021

## pkg.go.dev/github.com/corona10/goimagehash

- AverageHash
- DifferenceHash
- PerceptionHash

## pkg.go.dev/github.com/Nr90/imgsim

- AverageHash
- DifferenceHash

## Percentage difference between images

Lets start with a reference image:

~~~
https://ia800309.us.archive.org/9/items/
mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1/
mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg
~~~

Now lets look at an image to compare:

~~~
https://i.ytimg.com/vi/2rYQg0QmhX8/maxresdefault.jpg
~~~

First we need to crop the second image. Then, we need to scale the original
down. Then we can compare.

- <https://rosettacode.org/wiki/Percentage_difference_between_images>
- <https://wiki.musicbrainz.org/Cover_Art_Archive/API>
- https://stackoverflow.com/questions/32680834
