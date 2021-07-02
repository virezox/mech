# Mech

Mechanize

https://pkg.go.dev/github.com/89z/mech

## Sites

1. Deezer
2. MusicBrainz
3. Rotten Tomatoes
4. YouTube

## Transport

How can I tell if response is Gzip encoded? With cURL, I can do this:

~~~
PS C:\> curl -v -H 'Accept-Encoding: gzip' https://github.com/manifest.json
< content-encoding: gzip
~~~

and how can I see the Gzipped size? Same cURL command:

~~~
PS C:\> curl -v -H 'Accept-Encoding: gzip' https://github.com/manifest.json
< content-length: 345
~~~

Now with Go, how can I tell if response is Gzip encoded? With Go, how can I see
the Gzipped size? I dont like how Go is deleting response headers:

- <https://github.com/golang/go/blob/go1.16.5/src/net/http/response_test.go#L638-L641>
- https://github.com/golang/go/blob/go1.16.5/src/net/http/transport.go#L2186-L2192
- https://golang.org/pkg/net/http#Transport.DisableCompression
