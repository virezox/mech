# Android Widevine

## How to dump L3 CDM?

Install [Android Studio][1]. Then create Android virtual device:

API Level | ABI | Target
----------|-----|--------------------------
24        | x86 | Android 7.0 (Google APIs)

Then download [Widevine Dumper][2]. Then install:

~~~
pip install -r requirements.txt
~~~

Then download [Frida server][3], example file:

~~~
frida-server-15.1.22-android-x86.xz
~~~

Then start Frida server:

~~~
adb root
adb push frida-server-15.1.17-android-x86 /data/frida-server
adb shell chmod +x /data/frida-server
adb shell /data/frida-server
~~~

Then start Android Chrome and visit [Shaka Player][4]. Then start dumper:

~~~
python dump_keys.py
~~~

Once you see "Hooks completed", go back to Chrome, scroll down and click LOAD.
Result:

~~~
2022-05-21 02:10:52 PM - Helpers.Scanner - 49 - INFO - Key pairs saved at
key_dumps\Android Emulator 5554/private_keys/4464/2770936375
~~~

[1]://developer.android.com/studio
[2]://github.com/wvdumper/dumper
[3]://github.com/frida/frida/releases
[4]://integration.widevine.com/player
