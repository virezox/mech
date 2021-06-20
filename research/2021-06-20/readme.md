# June 20 2021

Since ANDROID is smaller, we want to pull as much from it as possible. Here is
what we need to check:

~~~go
StreamingData struct {
   AdaptiveFormats []struct {
      Bitrate         int64  // pass
      ContentLength   int64  // pass
      Height          int    // pass
      Itag            int    // pass
      MimeType        string // pass
      URL             string // pass
   }
}
VideoDetails struct {
   Author           string // pass
   ShortDescription string // pass
   Title            string // pass
   ViewCount        int    // pass
}
Microformat struct {
   PlayerMicroformatRenderer struct {
      AvailableCountries []string // fail
      PublishDate        string   // fail
   }
}
~~~

api    | client      | format | size    | media | date
-------|-------------|--------|---------|-------|-----
player | android     | JSON   | 98.3 KB | yes   | no
player | web         | JSON   | 228 KB  | yes   | 2009-10-24

`next` with `ANDROID` doesnt return any media, so we will need to call `player`
with `ANDROID` regardless. To get the date, we could do `next` with `ANDROID`,
but format is not machine readable:

~~~
Published on Oct 24, 2009
~~~

Same for `next` with `WEB`:

~~~
Oct 24, 2009
~~~

better would be `player` with `WEB` (JSON 228 KB proto 160 KB):

~~~
2009-10-24
~~~

- https://github.com/TeamNewPipe/NewPipeExtractor/issues/562
- https://github.com/TeamNewPipe/NewPipeExtractor/issues/568
- https://github.com/tombulled/innertube/blob/main/innertube/infos.py
