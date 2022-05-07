# Channel 4

- <https://dashif.org/identifiers/content_protection>
- https://bento4.com/downloads
- https://tools.axinom.com/generators/PsshBox

## How to watch from US

Download MITM Proxy:

https://mitmproxy.org/downloads

Start `mitmproxy`. Enter the following:

~~~
:set modify_headers '/~u vod.stream/X-Forwarded-For/25.0.0.0'
~~~

The following instruction is for Firefox, but Chrome should be similar. Open
Application Menu and click "Options". Then click "Privacy & Security". Then
scroll to "Security", and click "View Certificates". Then click "Authorities".
Then click "Import". Find this file:

~~~
C:\Users\Steven\.mitmproxy\mitmproxy-ca.pem
~~~

and click "Open". Then click "Trust this CA to identify websites". Then click
"OK". Then click "OK". Then click "General". Scroll to "Network Settings", and
click "Settings". Click "Manual proxy configuration". After "HTTP Proxy", enter:

~~~
127.0.0.1
~~~

after "Port", enter:

~~~
8080
~~~

then click "OK". Now browse to a page:

https://channel4.com/programmes/frasier/on-demand/18926-001
