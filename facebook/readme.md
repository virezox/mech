# Facebook

- https://android.stackexchange.com/questions/246303/facebook-whitehat-settings
- https://play.google.com/store/apps/details?id=com.facebook.katana

HTTP Toolkit:

Install Facebook app. Go to Menu, Settings & Privacy, Whitehat Settings. Click
"allow user installed certificates". FORCE STOP app. Start HTTP Toolkit. Start
app. Notice that requests are not being captured, and Certificate rejected
errors are occurring.

https://github.com/httptoolkit/frida-android-unpinning/issues/18

MITM Proxy:

Install user certificate. Turn off proxy. Install Facebook app. Go to Menu,
Settings & Privacy, Whitehat Settings. Click "allow user installed
certificates". FORCE STOP app. Turn on proxy. Start app. Notice that requests
are not being captured. Go back to Whitehat Settings, and click "Proxy for
Platform API requests". Enter 127.0.0.1:8080. FORCE STOP app. Start app. Notice
that requests are not being captured.

https://github.com/mitmproxy/mitmproxy/discussions/5271
