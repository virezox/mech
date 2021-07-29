# Distance between valid and invalid images

I took a "source" image, and compared it against something pretty close, and against something pretty different. here are the three images:

source:
https://ia800309.us.archive.org/9/items/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1/mbid-a40cb6e9-c766-37c4-8677-7eb51393d5a1-9261666555.jpg

good:
http://i.ytimg.com/vi/11Bvzknjo2Q/hqdefault.jpg

bad:
<http://i.ytimg.com/vi/jCMi9_6vnxk/hqdefault.jpg>

The last two images have padding, so I removed that before doing any hashing. My thinking is that, the distance value should be low for the "good" image, and high for the "bad" image. Results with `devedge/imagehash`:

good distance | bad distance
--------------|-------------
12            | 14

So what concerns me, is the difference is almost nothing. Compare to `corona10/goimagehash` [1]:

good distance | bad distance
--------------|-------------
1             | 17

1. https://github.com/corona10/goimagehash
