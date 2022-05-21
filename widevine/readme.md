# Widevine

> Theatricality and deception, powerful agents to the uninitiated.
>
> But we are initiated, aren’t we, Bruce?
>
> The Dark Knight Rises (2012)

## L1

Amzn needs AndroidCDM L1 for 720P/1080P/4K, Disney+ needs Android L1 for
1080P+, and NetFlix I think needs L1 for Main Profile & UHD.

## Where did proto files come from?

- <https://github.com/cryptonek/widevine-l3-decryptor/blob/master/license_protocol.proto>
- https://github.com/TDenisM/widevinedump/tree/main/pywidevine/cdm/formats

## Where to download L3 CDM?

<https://github.com/Jnzzi/4464_L3-CDM>

## How to dump L3 CDM?

First create Android virtual device:

API Level | ABI | Target
----------|-----|--------------------------
24        | x86 | Android 7.0 (Google APIs)

Then download tool:

https://github.com/wvdumper/dumper

Then install:

~~~
pip install -r requirements.txt
~~~

Download Frida server:

https://github.com/frida/frida/releases

Example file:

~~~
frida-server-15.1.22-android-x86.xz
~~~

Start Frida server:

~~~
adb root
adb push frida-server-15.1.17-android-x86 /data/frida-server
adb shell chmod +x /data/frida-server
adb shell /data/frida-server
~~~

Then start Android Chrome and visit this page:

https://bitmovin.com/demos/drm

If you receive this prompt:

> bitmovin.com wants to play protected content. Your device’s identity will be
> verified by Google.

Click ALLOW. Then start dumper:

~~~
python dump_keys.py
~~~

Once you see "Hooks completed", go back to Chrome, scroll down and click LOAD.
Result:

~~~
2022-05-21 02:10:52 PM - Helpers.Scanner - 49 - INFO - Key pairs saved at
key_dumps\Android Emulator 5554/private_keys/4464/2770936375
~~~
