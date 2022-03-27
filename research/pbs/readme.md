# PBS

- https://github.com/yt-dlp/yt-dlp/issues/3092
- https://www.pbs.org/wnet/nature/about-american-horses/26867/

This is it:

~~~
GET /v3/android/screens/video-assets/american-horses-i5v309/?station_id=b3291387-78a4-41e1-beb0-da2f61a96a3e HTTP/1.1
X-PBS-PlatformVersion: 5.5.5
Authorization: Basic YW5kcm9pZDpiYVhFN2h1bXVWYXQ=
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Host: content.services.pbs.org
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

fuck yeah:

~~~xml
<intent-filter android:label="@string/app_name" android:autoVerify="true">
    <action android:name="android.intent.action.VIEW"/>
    <action android:name="android.intent.action.CAST"/>
    <category android:name="android.intent.category.DEFAULT"/>
    <category android:name="android.intent.category.BROWSABLE"/>
    <data android:scheme="http" android:host="www.pbs.org"/>
    <data android:scheme="https" android:host="www.pbs.org"/>
    <data android:pathPattern="/video/.*/"/>
</intent-filter>
<intent-filter android:label="@string/app_name" android:autoVerify="true">
    <action android:name="android.intent.action.VIEW"/>
    <action android:name="android.intent.action.CAST"/>
    <category android:name="android.intent.category.DEFAULT"/>
    <category android:name="android.intent.category.BROWSABLE"/>
    <data android:scheme="http" android:host="www.pbs.org"/>
    <data android:scheme="https" android:host="www.pbs.org"/>
    <data android:pathPattern="/show/.*/"/>
    <data android:pathPattern="/show/.*/episodes/season/.*/"/>
</intent-filter>
~~~

https://stackoverflow.com/questions/28802115/i-am-trying-to-test-android-deep

~~~
adb shell am start -d https://www.pbs.org/video/american-horses-i5v309/
https://www.pbs.org/video/nova-universe-revealed-milky-way-4io957/
https://www.pbs.org/video/pandora-papersmassacre-in-el-salvador-v2wjzz/

https://www.pbs.org/nova/video/nova-universe-revealed-milky-way/
https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/
~~~
