# August 28 2021

## JavaScript

https://github.com/GoneToneStudio/node-google-play-api/issues/5

## C#

https://github.com/kagasu/GooglePlayStoreApi/issues/13

## Android API

~~~
Google Service Framework
38B5418D8683ADBB
~~~

Yeah, the API was the issue. Using API 24 fails, but API 25 or higher works. It
applies to all devices, not just Virtual Devices.
