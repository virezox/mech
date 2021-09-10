# GPS OAuth

## JA3

- https://github.com/CUCyber/ja3transport/issues/8
- https://github.com/D4-project/sensor-d4-tls-fingerprinting/issues/20
- https://github.com/Danny-Dasilva/CycleTLS/issues/42
- https://github.com/dreadl0ck/ja3/issues/10
- https://github.com/open-ch/ja3/issues/2
- https://github.com/salesforce/ja3
- https://ja3er.com

~~~json
4865-49196-49200-49195-49199-52393-52392-159-158-52394-49327-49325-49326-49324-49188-49192-49187-49191-49162-49172-49161-49171-49315-49311-49314-49310-107-103-57-51-157-156-49313-49309-49312-49308-61-60-53-47-255
~~~

## TLS fingerprint

https://tlsfingerprint.io/find/cipher/ccaa

## Python

~~~
pip install pycryptodomex
pip install requests
~~~

First get this package:

https://github.com/simon-weber/gpsoauth

Then make a request here:

- https://client.tlsfingerprint.io:8443
- https://www.howsmyssl.com/a/check

Then rebuild that fingerprint in your language of choice. How to get the cipher
names:

https://unix.stackexchange.com/questions/208412
