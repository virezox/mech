# August 14 2021

First make request 0:

~~~
GET /api/reference/videos HTTP/1.1
Host: developer.vimeo.com
~~~

then in response:

~~~
"x-playground-token":"81f5441838f52d92c8b5de970abf819fbfc843fb38c6337a5119958..."
~~~

Then make request 1:

~~~
POST /api/playground/callable HTTP/1.1
Host: developer.vimeo.com
Content-Type: application/json;charset=utf-8
Cookie: session=xdQ3gT5ZpwNJ721eHt1UI9AAOHoYP2SaTiu07jzL
X-CSRF-TOKEN: QALMpMDZmFgmoLbab1bUK6nn5rxW5XHIZtTvqaDb

{
   "ptoken":"326c3d16c93ab0da858467e26947eeecded3423c1a8628e747b01c786d65c388.1628971412",
   "group":"videos",
   "operation_id":"get_video",
   "app":"2221",
   "query_params":"{}",
   "payload_params":"{}",
   "segments":"{\"video_id\":\"66531465\"}"
}
~~~

then in response:

~~~
"Authorization":"jwt eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2Mjg5NjU..."
~~~

Then make request 2.

https://developer.vimeo.com/api/reference/videos
