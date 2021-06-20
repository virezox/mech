# June 20 2021

api    | client      | format | size    | media | date
-------|-------------|--------|---------|-------|-----
next   | android     | JSON   | 1.3 MB  | no    | Published on Oct 24, 2009
next   | android     | proto  | 319 KB  | no    | Published on Oct 24, 2009
next   | web         | JSON   | 522 KB  | no    | Oct 24, 2009
next   | web         | proto  | 131 KB  | no    | Oct 24, 2009
player | `web_remix` | JSON   | 242 KB  | yes   | 2009-10-24
player | android     | JSON   | 98.3 KB | yes   | no
player | web         | JSON   | 228 KB  | yes   | 2009-10-24
player | web         | proto  | 160 KB  | yes   | 2009-10-24

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
