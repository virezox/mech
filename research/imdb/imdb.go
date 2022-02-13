package imdb

import (
   "crypto/hmac"
   "crypto/sha256"
   "encoding/base64"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

var logLevel format.LogLevel

func newCredentials() (*credentials, error) {
   req, err := http.NewRequest(
      "POST",
      "https://api.imdbws.com/authentication/credentials/temporary/android850",
      strings.NewReader(`{"appKey":"4f833099-e4fe-4912-80f3-b1b169097914"}`),
   )
   if err != nil {
      return nil, err
   }
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   cred := new(credentials)
   if err := json.NewDecoder(res.Body).Decode(cred); err != nil {
      return nil, err
   }
   return cred, nil
}

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
*/
type credentials struct {
   Resource struct {
      AccessKeyID string
      SecretAccessKey string
      SessionToken string
   }
}

////////////////////////////////////////////////////////////////////////////////

/*
GET https://api.imdbws.com/template/imdb-android-writable/8.5.runway-gallery-images.jstl/render?rgconst=rg2774637312&offset=0&limit=300 HTTP/2.0
x-amzn-sessionid: 726-7519652-9073110
x-amz-security-token: IQoJb3JpZ2luX2VjEBwaCXVzLWVhc3QtMSJGMEQCIEwdxMG+9D3leVkRGNyQidgitf7J52tqZnDdwKmHEwXmAiBRxEChX+HMSsC7lG/tSG87Jt8pDvyooX4S5AA5yvlFdyqWAghlEAAaDDU4MDU2NjA4ODMxNyIMVtAHM8cPLYASfcp2KvMBT2HmUPdV6smmjH9Fewq20CqNhQAvqFC8W2eAcM1C33q7wVUyFf9ecNya72jZOqCqpLUOu3XGEEDQRu80haqWqHbRV8UI2u7Vks2dJBGjGb/mxfKdZK2TznjB844tyijN+7Gwp8NfWnltUDCAC4aK9LNTlML46Uk5jqgNFST9Ft1Y1gX3V49B5uF83qC/jxysMJHOXhcbqwslkjbIKw4Uwl6fQsw1+i389HJ2SEH7oyiZGdinEXfwOgPU7GvMmnjgya+K0C4JMVeA55pihijLHnYUxbXDzIqk3craYRYqeSiQARrMlkSTiQbRU2Uuhoh0vtYTMJmfoJAGOpoBy+xjWQxXVjnqDiPThlpAq17xcFJqNhArrw/vNuYupOTWm9oCrT+G/4Z2Oji3Q2YBjTaYo8p3+bLBOIah4AW82tC4ex966M1wOekC86nTYkvicYHBTtjc/4767iK3kRm6tU/PuX336W7CszO1LAU49f+58Evq7ICKd9JhHRAvTmRcA0zL3EEOukf+8/ws9GB6mt29qX216RvUJw==

x-amz-date: Sat, 12 Feb 2022 19:50:50 GMT+00:00
x-amzn-authorization: AWS3 AWSAccessKeyId=ASIAYOLDPPJ66PEWUDFX,Algorithm=HmacSHA256,Signature=C8vM5+Dh/Q1jozoUpSlzsqB/enPPzCgY34gqHmS764Y=,SignedHeaders=host;x-amz-date;x-amz-security-token;x-amzn-sessionid
*/
func (c credentials) Gallery() (*http.Response, error) {
   var buf strings.Builder
   buf.WriteString("https://api.imdbws.com/template")
   buf.WriteString("/template/imdb-android-writable")
   buf.WriteString("/8.5.runway-gallery-images.jstl/render")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "limit=300&offset=0&rgconst=rg2774637312"
   req.Header = http.Header{
      "x-amzn-security-token": {c.Resource.SessionToken},
      "x-amzn-sessionid": {"726-7519652-9073110"},
   }
   logLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
   /*
   sign := []string{"GET"}
   sign = append(sign, "/template/imdb-android-writable/8.5.runway-gallery-images.jstl/render")
   sign = append(sign, "limit=300&offset=0&rgconst=rg2774637312")
   sign = append(sign, `host:api.imdbws.com
x-amz-date:Sat, 12 Feb 2022 19:50:50 GMT+00:00
x-amz-security-token:IQoJb3JpZ2luX2VjEBwaCXVzLWVhc3QtMSJGMEQCIEwdxMG+9D3leVkRGNyQidgitf7J52tqZnDdwKmHEwXmAiBRxEChX+HMSsC7lG/tSG87Jt8pDvyooX4S5AA5yvlFdyqWAghlEAAaDDU4MDU2NjA4ODMxNyIMVtAHM8cPLYASfcp2KvMBT2HmUPdV6smmjH9Fewq20CqNhQAvqFC8W2eAcM1C33q7wVUyFf9ecNya72jZOqCqpLUOu3XGEEDQRu80haqWqHbRV8UI2u7Vks2dJBGjGb/mxfKdZK2TznjB844tyijN+7Gwp8NfWnltUDCAC4aK9LNTlML46Uk5jqgNFST9Ft1Y1gX3V49B5uF83qC/jxysMJHOXhcbqwslkjbIKw4Uwl6fQsw1+i389HJ2SEH7oyiZGdinEXfwOgPU7GvMmnjgya+K0C4JMVeA55pihijLHnYUxbXDzIqk3craYRYqeSiQARrMlkSTiQbRU2Uuhoh0vtYTMJmfoJAGOpoBy+xjWQxXVjnqDiPThlpAq17xcFJqNhArrw/vNuYupOTWm9oCrT+G/4Z2Oji3Q2YBjTaYo8p3+bLBOIah4AW82tC4ex966M1wOekC86nTYkvicYHBTtjc/4767iK3kRm6tU/PuX336W7CszO1LAU49f+58Evq7ICKd9JhHRAvTmRcA0zL3EEOukf+8/ws9GB6mt29qX216RvUJw==
x-amzn-sessionid:726-7519652-9073110`)
   sign = append(sign, "")
   sign = append(sign, "")
   key := "vOdRmrc3QqCxw0bnpE09ef1Z6kGpvzqEr1bNt5UX"
   signed := signRequest(strings.Join(sign, "\n"), key)
   */
}

func SignRequest(plain, key string) string {
   h1 := sha256.New()
   h1.Write([]byte(plain))
   h2 := hmac.New(sha256.New, []byte(key))
   h2.Write(h1.Sum(nil))
   return base64.StdEncoding.EncodeToString(h2.Sum(nil))
}
