# Widevine

## How to dump L3 CDM?

Download [Widevine Dumper][2]. Then install:

~~~
pip install -r requirements.txt
~~~

Then download [Frida server][3], example file:

~~~
frida-server-15.1.27-windows-x86_64.exe.xz
~~~

Then start Frida server. Then start browser and visit [Shaka Player][4]. Then
start dumper:

~~~
python dump_keys.py
~~~

Once you see "Hooks completed", go back to Chrome, scroll down and click LOAD.
Result:

~~~
2022-05-21 02:10:52 PM - Helpers.Scanner - 49 - INFO - Key pairs saved at
key_dumps\Android Emulator 5554/private_keys/4464/2770936375
~~~

[2]://github.com/wvdumper/dumper
[3]://github.com/frida/frida/releases
[4]://integration.widevine.com/player
