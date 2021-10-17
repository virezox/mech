# Certificate pinning

## Stack Exchange

Check questions, and ask question if need be:

https://android.stackexchange.com/search?q=certificate+pinning

## MITM Proxy

Fails with version 3:

https://github.com/mitmproxy/mitmproxy/issues/4838

but might be able to get it working:

> adding client certificates to the system-wide trust store, which is by default
> trusted by all apps

~~~
C:\Users\Steven\.mitmproxy
~~~

- https://blog.nviso.eu/2017/12/22/intercepting-https-traffic-from-apps-on-android-7-using-magisk-burp
- https://docs.mitmproxy.org/stable/concepts-certificates
- https://stackoverflow.com/questions/44942851/install-user-certificate-via-adb

## EdXposed

- https://github.com/ElderDrivers/EdXposed
- https://github.com/ViRb3/TrustMeAlready

First need to get Magisk:

https://github.com/topjohnwu/Magisk

## HTTP Toolkit

Works with version 3 fails with version 4.1.

https://httptoolkit.tech/blog/frida-certificate-pinning

## APK MITM

CLI application that automatically removes certificate pinning from Android APK
files:

https://github.com/shroudedcode/apk-mitm/issues/72

## Burp Suite

Fails with version 3.

## JustTrustMe

https://github.com/Fuzion24/JustTrustMe/issues/61
