# YouTube

I think I am about finished with my research. I didnt make a lot of progress
sadly. One addition, and several changes:

~~~diff
+TVHTML5_AUDIO;2.0

-ANDROID_TV;2.13
+ANDROID_TV;2.16.032

-ANDROID_TV_KIDS;1.0.0
+ANDROID_TV_KIDS;1.15.03

-ANDROID_VR;0.1
+ANDROID_VR;1.28.63

-MWEB;2.20220404
+MWEB;2.20220405.01.00

-TVHTML5;7.20220325
+TVHTML5;7.20220404.09.00

-TVHTML5_CAST;1.1
+TVHTML5_CAST;1.1.426206631

-TV_UNPLUGGED_ANDROID;0.1
+TV_UNPLUGGED_ANDROID;1.13.02

-WEB;2.20220404
+WEB;2.20220405.00.00

-WEB_CREATOR;1.20220330
+WEB_CREATOR;1.20220405.02.00

-WEB_KIDS;2.20220404
+WEB_KIDS;2.20220405.00.00

-WEB_REMIX;1.20220330
+WEB_REMIX;1.20220330.01.00

-WEB_UNPLUGGED;1.20220330
+WEB_UNPLUGGED;1.20220403.00.00
~~~

I checked every single YouTube Android app in the Google Play store, as well as
every single old YouTube app at ApkMirror.com. So I think these are a lost
cause:

~~~
ANDROID_CASUAL(54),
ANDROID_GAMING(24),
ANDROID_INSTANT(20),
ANDROID_PRODUCER(91),
ANDROID_SPORTS(36),
ANDROID_WITNESS(34),
IOS_GAMING(25),
IOS_INSTANT(17),
IOS_SPORTS(37),
IOS_WITNESS(35),
~~~

I dont have a way to hack IOS or XBOX, so these are out for me:

~~~
IOS_DIRECTOR(40),
IOS_PILOT_STUDIO(53),
IOS_TABLOID(22),
XBOX(11),
~~~

which leaves these:

~~~
CLIENTX(12),
UNKNOWN_INTERFACE(0),
WEB_GAMING(32),
WEB_LIVE_STREAMING(73),
WEB_MUSIC_EMBEDDED_PLAYER(86),
~~~

These might be possible, but we need to find the correct request to have them
returned. For example, here is a request that returns the current MWEB client:

~~~
GET / HTTP/1.1
Host: m.youtube.com
User-Agent: iPad
~~~
