package main

import (
   "crypto/hmac"
   "crypto/sha256"
   "encoding/base64"
   "os"
   "io"
)

/*
{
  "@meta": {
    "operation": "GetTemporaryCredentials",
    "requestId": "8983621e-aa23-4d22-9d39-f19030ae4c0d",
    "serviceTimeMs": 8.390552
  },
  "resource": {
    "@type": "imdb.api.auth.credentials.temporary",
    "accessKeyId": "ASIAYOLDPPJ66PEWUDFX",
    "expirationTimeStamp": "2022-02-14T07:50:49Z",
    "sessionToken": "IQoJb3JpZ2luX2VjEBwaCXVzLWVhc3QtMSJGMEQCIEwdxMG+9D3leVkRGNyQidgitf7J52tqZnDdwKmHEwXmAiBRxEChX+HMSsC7lG/tSG87Jt8pDvyooX4S5AA5yvlFdyqWAghlEAAaDDU4MDU2NjA4ODMxNyIMVtAHM8cPLYASfcp2KvMBT2HmUPdV6smmjH9Fewq20CqNhQAvqFC8W2eAcM1C33q7wVUyFf9ecNya72jZOqCqpLUOu3XGEEDQRu80haqWqHbRV8UI2u7Vks2dJBGjGb/mxfKdZK2TznjB844tyijN+7Gwp8NfWnltUDCAC4aK9LNTlML46Uk5jqgNFST9Ft1Y1gX3V49B5uF83qC/jxysMJHOXhcbqwslkjbIKw4Uwl6fQsw1+i389HJ2SEH7oyiZGdinEXfwOgPU7GvMmnjgya+K0C4JMVeA55pihijLHnYUxbXDzIqk3craYRYqeSiQARrMlkSTiQbRU2Uuhoh0vtYTMJmfoJAGOpoBy+xjWQxXVjnqDiPThlpAq17xcFJqNhArrw/vNuYupOTWm9oCrT+G/4Z2Oji3Q2YBjTaYo8p3+bLBOIah4AW82tC4ex966M1wOekC86nTYkvicYHBTtjc/4767iK3kRm6tU/PuX336W7CszO1LAU49f+58Evq7ICKd9JhHRAvTmRcA0zL3EEOukf+8/ws9GB6mt29qX216RvUJw=="
    "secretAccessKey": "vOdRmrc3QqCxw0bnpE09ef1Z6kGpvzqEr1bNt5UX",
  }
}

GET https://api.imdbws.com/template/imdb-android-writable/8.5.runway-gallery-images.jstl/render?rgconst=rg2774637312&offset=0&limit=300 HTTP/2.0
x-amz-date: Sat, 12 Feb 2022 19:50:50 GMT+00:00
x-amz-security-token: IQoJb3JpZ2luX2VjEBwaCXVzLWVhc3QtMSJGMEQCIEwdxMG+9D3leVkRGNyQidgitf7J52tqZnDdwKmHEwXmAiBRxEChX+HMSsC7lG/tSG87Jt8pDvyooX4S5AA5yvlFdyqWAghlEAAaDDU4MDU2NjA4ODMxNyIMVtAHM8cPLYASfcp2KvMBT2HmUPdV6smmjH9Fewq20CqNhQAvqFC8W2eAcM1C33q7wVUyFf9ecNya72jZOqCqpLUOu3XGEEDQRu80haqWqHbRV8UI2u7Vks2dJBGjGb/mxfKdZK2TznjB844tyijN+7Gwp8NfWnltUDCAC4aK9LNTlML46Uk5jqgNFST9Ft1Y1gX3V49B5uF83qC/jxysMJHOXhcbqwslkjbIKw4Uwl6fQsw1+i389HJ2SEH7oyiZGdinEXfwOgPU7GvMmnjgya+K0C4JMVeA55pihijLHnYUxbXDzIqk3craYRYqeSiQARrMlkSTiQbRU2Uuhoh0vtYTMJmfoJAGOpoBy+xjWQxXVjnqDiPThlpAq17xcFJqNhArrw/vNuYupOTWm9oCrT+G/4Z2Oji3Q2YBjTaYo8p3+bLBOIah4AW82tC4ex966M1wOekC86nTYkvicYHBTtjc/4767iK3kRm6tU/PuX336W7CszO1LAU49f+58Evq7ICKd9JhHRAvTmRcA0zL3EEOukf+8/ws9GB6mt29qX216RvUJw==
x-amzn-authorization: AWS3 AWSAccessKeyId=ASIAYOLDPPJ66PEWUDFX,Algorithm=HmacSHA256,Signature=C8vM5+Dh/Q1jozoUpSlzsqB/enPPzCgY34gqHmS764Y=,SignedHeaders=host;x-amz-date;x-amz-security-token;x-amzn-sessionid
x-amzn-sessionid: 726-7519652-9073110

*/
func old() {
   string_to_sign := `GET
/template/imdb-android-writable/8.5.runway-gallery-images.jstl/render

host:api.imdbws.com
x-amz-date:Sat, 12 Feb 2022 19:50:50 GMT+00:00
x-amz-security-token:IQoJb3JpZ2luX2VjEBwaCXVzLWVhc3QtMSJGMEQCIEwdxMG+9D3leVkRGNyQidgitf7J52tqZnDdwKmHEwXmAiBRxEChX+HMSsC7lG/tSG87Jt8pDvyooX4S5AA5yvlFdyqWAghlEAAaDDU4MDU2NjA4ODMxNyIMVtAHM8cPLYASfcp2KvMBT2HmUPdV6smmjH9Fewq20CqNhQAvqFC8W2eAcM1C33q7wVUyFf9ecNya72jZOqCqpLUOu3XGEEDQRu80haqWqHbRV8UI2u7Vks2dJBGjGb/mxfKdZK2TznjB844tyijN+7Gwp8NfWnltUDCAC4aK9LNTlML46Uk5jqgNFST9Ft1Y1gX3V49B5uF83qC/jxysMJHOXhcbqwslkjbIKw4Uwl6fQsw1+i389HJ2SEH7oyiZGdinEXfwOgPU7GvMmnjgya+K0C4JMVeA55pihijLHnYUxbXDzIqk3craYRYqeSiQARrMlkSTiQbRU2Uuhoh0vtYTMJmfoJAGOpoBy+xjWQxXVjnqDiPThlpAq17xcFJqNhArrw/vNuYupOTWm9oCrT+G/4Z2Oji3Q2YBjTaYo8p3+bLBOIah4AW82tC4ex966M1wOekC86nTYkvicYHBTtjc/4767iK3kRm6tU/PuX336W7CszO1LAU49f+58Evq7ICKd9JhHRAvTmRcA0zL3EEOukf+8/ws9GB6mt29qX216RvUJw==
x-amzn-sessionid:726-7519652-9073110

`
   mac := hmac.New(sha256.New, []byte("vOdRmrc3QqCxw0bnpE09ef1Z6kGpvzqEr1bNt5UX"))
   io.WriteString(mac, string_to_sign)
   base64.NewEncoder(base64.StdEncoding, os.Stdout).Write(mac.Sum(nil))
}
