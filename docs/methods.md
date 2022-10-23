# Methods

## anti\_vm

https://github.com/rednaga/APKiD

## iOS

https://github.com/TrungNguyen1909/qemu-t8030/issues/52

## MITM Proxy

First download [1], then start `mitmproxy.exe`. Address and port should be in
the bottom right corner. Default should be:

~~~
*:8080
~~~

Assuming the above, go to Android Emulator and set proxy:

~~~
127.0.0.1:8080
~~~

Then open Google Chrome on Virtual Device, and browse to:

~~~
http://example.com
~~~

To exit, press `q`, then `y`. To capture HTTPS, open Google Chrome on Virtual
Device, and browse to <http://mitm.it>. Click on the Android certificate. Under
"Certificate name" enter "MITM", then click "OK". Then browse to:

~~~
https://example.com
~~~

Disable compression:

~~~
set anticomp true
~~~

1. https://mitmproxy.org/downloads

## Python

If you want to MITM a Python program, you might need to set one or more of
these:

~~~ps1
$env:HTTPS_PROXY = 'http://127.0.0.1:8080'
$env:REQUESTS_CA_BUNDLE = 'C:\Users\Steven\.mitmproxy\mitmproxy-ca.pem'
$env:SSL_CERT_FILE = 'C:\Users\Steven\.mitmproxy\mitmproxy-ca.pem'
~~~

## Geo block

~~~
set modify_headers '/~u vod.stream/X-Forwarded-For/25.0.0.0'
~~~

or for all:

~~~
set modify_headers '/~q/X-Forwarded-For/25.0.0.0'
~~~

https://docs.mitmproxy.org/stable/overview-features

## APK to Java

~~~
jadx-gui.bat com.pinterest-10098030.apk
jadx.bat com.google.android.youtube-1528288704.apk
~~~

https://github.com/skylot/jadx

## APK

To download, you can use my tool:

https://github.com/89z/googleplay

To install, drag file to emulator home screen. To uninstall, long press on the
app, and drag to "Uninstall". To force stop, long press on the app, and drag
to "App info".

## Deep linking

<https://wikipedia.org/wiki/Mobile_deep_linking>

Click a link in Android Chrome. In some cases, the target needs to be a
different origin from the source. A prompt should come up that says "Open
with". Click the option for the app, then "JUST ONCE". The link should open in
the app, and if you are monitoring, you should see the request. Also, you can
check the `Androidmanifest.xml` file:

~~~xml
<intent-filter android:autoVerify="true">
   <action android:name="android.nfc.action.NDEF_DISCOVERED"/>
   <action android:name="android.intent.action.VIEW"/>
   <category android:name="android.intent.category.DEFAULT"/>
   <category android:name="android.intent.category.BROWSABLE"/>
   <data android:scheme="https"/>
   <data android:scheme="http"/>
   <data android:host="www.pinterest.com"/>
   <data android:host="post.pinterest.com"/>
   <data android:host="pin.it"/>
   <!-- ... -->
</intent-filter>
~~~

So only link with those host will get noticed by the app. Finally, if you have
`adb`, you can use it like this:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://abc.com/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you
~~~

Note, in some cases you need to start the app at least once before trying a
deep link.

## Android Studio

First download the package [1]. Start the program, and click "More Actions",
"AVD Manager", "Create Virtual Device". On the "Select Hardware" screen, click
"Next". On the "System Image" screen, click "x86 Images". If the APK you are
using supports "x86", then you can use versions down to API 23. Once you have
chosen, click "Download". Then click "Next". On the "Android Virtual Device"
screen, click "Finish". On the "Your Virtual Devices" screen, click "Launch
this AVD in the emulator".

1. https://developer.android.com/studio#downloads

If you need to configure a proxy, in the emulator click "More". On the
"Extended Controls" screen, click "Settings", "Proxy". Uncheck "Use Android
Studio HTTP proxy settings". Click "Manual proxy configuration". Enter "Host
name" and "Port number" as determined by the proxy program you are using. Click
"Apply", and you should see "Proxy status Success".

## System certificate

https://github.com/89z/piccolo/tree/master/cmd/mitmproxy-cert

## Frida

https://github.com/89z/piccolo/tree/master/cmd/frida-script
