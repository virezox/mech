# August 23 2021

<https://github.com/protocolbuffers/protobuf-go/blob/master/testing/protopack/pack_test.go>

To be fair, a couple of people did it already in Python [1] and Crystal [2], so
I just translated the code to Go. However doing so was quite difficult, as I
only found one implementation for Protobuf Wire in the standard library [3],
and the functions arent exported, so I would have had to vendor them. After
much looking, I found the Protopack package [4], which does exactly what is
needed here. I overlooked it at first, because the Protobuf module is huge, and
since it was under the `testing` folder, I didnt think it was anything.

1. https://github.com/user234683/youtube-local/blob/master/youtube/comments.py
2. https://github.com/iv-org/invidious/blob/master/src/invidious/comments.cr
3. https://github.com/golang/go/blob/master/src/runtime/pprof/protobuf.go
4. https://pkg.go.dev/google.golang.org/protobuf/testing/protopack

## next

~~~
next
request
continuationCommand

next
response
commentThreadRenderer commentRenderer
~~~

YouTube API for comments?

https://www.youtube.com/watch?v=q5UnT4Ik6KU

So we can search for a comment that identifies the end song.

~~~
The unbiased journalism we need
~~~

- https://github.com/golang/go/blob/master/src/runtime/pprof/protobuf.go
- https://github.com/iv-org/invidious/blob/master/src/invidious/comments.cr
- https://github.com/user234683/youtube-local/blob/master/youtube/comments.py
- https://github.com/user234683/youtube-local/blob/master/youtube/proto.py
- https://godocs.io/google.golang.org/protobuf/encoding/protowire

## clients

Check different clients
