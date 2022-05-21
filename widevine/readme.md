# Widevine

> Theatricality and deception, powerful agents to the uninitiated.
>
> But we are initiated, aren’t we, Bruce?
>
> The Dark Knight Rises (2012)

## L1

Amzn needs AndroidCDM L1 for 720P/1080P/4K, Disney+ needs Android L1 for
1080P+, and NetFlix I think needs L1 for Main Profile & UHD.

## RSA PRIVATE KEY

i tried with playstore one, it worked for me, and i even tried with out
playstore, it didnt worked

Try sticking with Android 7-9, start Frida-Server, start Chrome with Bitmovin
or some similar DRM video, then run the dumper. For some reason if you try run
the script first before Chrome, it won't attach to Chrome properly and won't
dump

- <https://github.com/Avalonswanderer/widevinel3_Android_PoC/issues/1>
- https://android.stackexchange.com/questions/218850/android-studio-emulator
- https://github.com/Avalonswanderer/wideXtractor/issues/1
- https://github.com/wvdumper/dumper/issues/27
- https://youtube.com/watch?v=JR4gDRYzY2c

## Where did proto files come from?

- <https://github.com/cryptonek/widevine-l3-decryptor/blob/master/license_protocol.proto>
- https://github.com/TDenisM/widevinedump/tree/main/pywidevine/cdm/formats

## Where to get CDM?

<https://github.com/Jnzzi/4464_L3-CDM>
