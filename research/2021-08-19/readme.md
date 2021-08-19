# August 19 2021

https://developer.vimeo.com/api/reference/videos

## Request 1

~~~
GET /api/reference/videos HTTP/1.1
Host: developer.vimeo.com
~~~

then in response:

~~~
Set-Cookie: XSRF-TOKEN=WaY1qXeM4NeWsOD24S701Md79pyJQp9qBQOcooJZ
Set-Cookie: session=nTUmzZZ11A8tTRZdObpI0huKpcFk0zEoqJjdqutW

x-playground-token	"9bd75d9905527e8e298a7602…186a78bec931.1628973628"
~~~

## Request 2

~~~
POST /api/playground/callable HTTP/1.1
Host: developer.vimeo.com
Content-Type: application/json;charset=utf-8
Cookie: session=xdQ3gT5ZpwNJ721eHt1UI9AAOHoYP2SaTiu07jzL
X-CSRF-TOKEN: QALMpMDZmFgmoLbab1bUK6nn5rxW5XHIZtTvqaDb

{
   "app":"2221",
   "group":"videos",
   "operation_id":"get_video",
   "payload_params":"{}",
   "ptoken":"326c3d16c93ab0da858467e26947eeecded3423c1a8628e747b01c786d65c388.1628971412",
   "query_params":"{}",
   "segments":"{\"video_id\":\"66531465\"}"
}
~~~

then in response:

~~~
"Authorization":"jwt eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2Mjg5NjU..."
~~~

## Request 3
