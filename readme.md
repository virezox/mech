# Mech

> I see him behind my lids in a bright grey shirt\
> I see him tripping running and falling, covered in dirt\
> I see a lot of these things lately I know\
> I know none of it is real
>
> [Blood Orange (2013)](//youtube.com/watch?v=yP9JsIhHxSg)

Mechanize

Some users might want to make anonymous requests, because of privacy or any
number of other reasons. This module allows people to do that. Most API these
days only offically support authenticated access. This is useful for the
company providing the API, as they can use the data for their own purposes
(analytics etc). However authentication really doesnt do anything for the end
user. Its just a pointless burden to getting the information you need for a
program you may be writing. Consider that in many cases, the same information
is available via HTML on the primary website, usually without being logged in.
So why can you do that with HTML, but not with the API? Well you can, using this
module.

https://pkg.go.dev/github.com/89z/mech

## Sites

1. YouTube
2. SoundCloud
3. Deezer
4. MusicBrainz

## HTML package

Takes HTML input, and can iterate through elements by tag name or by attribute
name and value. Content from text nodes can be returned. Also, you can check if
an element has a certain attribute, and return an attribute value given an
attribute name. Finally, you can indent and write the HTML to some output.

## JavaScript package

Takes JavaScript input, and will return a `map`. Keys are the variable names,
and values are the variable values. The values are returned as `byte` slices, to
make it easy to `json.Unmarshal`.

## Author

Steven Penny
