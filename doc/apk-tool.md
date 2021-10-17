# APK Tool

~~~
$env:PATH = 'D:\Desktop\jdk-17+35\bin'

java -jar apktool_2.6.0.jar d org.metabrainz.android.apk

java -jar apktool_2.6.0.jar b org.metabrainz.android -o app_patched.apk `
--use-aapt2

keytool -genkey -alias keys -keystore keys -keyalg DSA

jarsigner -verbose -keystore keys app_patched.apk keys
~~~

- https://bugs.openjdk.java.net/browse/JDK-8212111
- https://github.com/iBotPeaches/Apktool/issues/1978
- https://github.com/iBotPeaches/Apktool/issues/731
- https://stackoverflow.com/questions/52862256/charles-proxy-for-mobile-apps

## MusicBrainz

10.8 MB. Version 5 works with self-signed certificate. Version 4.1 - 4.5 dont
work, even without a proxy.

https://apkpure.com/musicbrainz/org.metabrainz.android

## Bandcamp

14.1 MB. Works with self-signed certificate.

https://apkpure.com/bandcamp/com.bandcamp.android

## Vimeo

35.5 MB. Requires Android 26.

https://apkpure.com/vimeo/com.vimeo.android.videoapp

## Spotify

40.0 MB

https://apkpure.com/spotify-music-i/com.spotify.music

## Instagram

40.8 MB

https://apkpure.com/instagram/com.instagram.android

## YouTube

42.9 MB

https://apkpure.com/youtube/com.google.android.youtube

## Google Play

45.4 MB

https://apkpure.com/google-play-store/com.android.vending

## SoundCloud

83.5 MB

https://apkpure.com/sound-cloud-android/com.soundcloud.android

## Reddit

113.3 MB

https://apkpure.com/reddit/com.reddit.frontpage
