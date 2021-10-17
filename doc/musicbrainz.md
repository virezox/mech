# MusicBrainz

Use this APK as an example:

https://apkpure.com/musicbrainz/org.metabrainz.android

Using version 3.0, I can intercept the traffic with HTTP Toolkit. What I
discovered is everything between 3.0 and 5 does not work, even without running
HTTP Toolkit. So lets see if we can get version 5 working.

- https://bugs.openjdk.java.net/browse/JDK-8212111
- https://github.com/iBotPeaches/Apktool/issues/1978
- https://github.com/iBotPeaches/Apktool/issues/731
- https://stackoverflow.com/questions/52862256/charles-proxy-for-mobile-apps

~~~
$env:PATH = 'D:\Desktop\jdk-17+35\bin'

java -jar apktool_2.6.0.jar d org.metabrainz.android.apk

java -jar apktool_2.6.0.jar b org.metabrainz.android -o app_patched.apk `
--use-aapt2

keytool -genkey -alias keys -keystore keys -keyalg DSA

jarsigner -verbose -keystore keys app_patched.apk keys
~~~
