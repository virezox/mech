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

https://godocs.io/github.com/89z/mech

## Sites

- Bandcamp
- GitHub
- Google Play
- Instagram
- MusicBrainz
- Reddit
- SoundCloud
- Vimeo
- YouTube

## Deezer

I have an implementation here:

https://github.com/89z/mech/tree/9dadd39c

However I have removed it for now, as I am busy with other stuff.

## Android

First get Android Studio [1]. Click More Actions, AVD Manager, Create Virtual
Device. Use the default device defintion, then click Next. Click "x86 Images"
and download:

Release Name | API Level | ABI | Target
-------------|-----------|-----|------------
Nougat       | 25        | x86 | Google APIs

Then Download HTTP Toolkit [2]. Just extract, dont need to install. Click
Android device connected via ADB. Then download APK [3]. Drag APK to device
home screen to install.

1. https://developer.android.com/studio#downloads
2. https://httptoolkit.tech
3. https://apkpure.com

## Author

Steven Penny
