# August 23 2021

~~~
protobuf wire in:readme language:go stars:>0
~~~

https://pkg.go.dev/google.golang.org/protobuf/testing/protopack

we need to call `proto.Marshal`:

https://pkg.go.dev/google.golang.org/protobuf/proto#Marshal

to do that, we need any type that has a `ProtoReflect` method:

https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#ProtoMessage

what types have a `ProtoReflect` method?

https://github.com/protocolbuffers/protobuf-go

## next

~~~
next
request
continuationCommand

next
response
commentThreadRenderer commentRenderer -language:json
~~~

YouTube API for comments?

https://www.youtube.com/watch?v=q5UnT4Ik6KU

So we can search for a comment that identifies the end song.

~~~
The unbiased journalism we need
~~~

- https://github.com/DeDiS/protobuf
- https://github.com/golang/go/blob/master/src/runtime/pprof/protobuf.go
- https://github.com/iv-org/invidious/blob/master/src/invidious/comments.cr
- https://github.com/user234683/youtube-local/blob/master/youtube/comments.py
- https://github.com/user234683/youtube-local/blob/master/youtube/proto.py
- https://godocs.io/google.golang.org/protobuf/encoding/protowire

## size difference

next:

~~~
Content-Encoding: gzip
Content-Length: 34976
~~~

watch:

~~~
Content-Encoding: gzip
79,940 bytes
~~~

make sure "next next" actually works before deleting proto stuff.

## clients

Check different clients
