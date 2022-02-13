package imdb

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/base64"
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

const (
   appKey = "4f833099-e4fe-4912-80f3-b1b169097914"
   origin = "https://api.imdbws.com"
   sessionID = "726-7519652-9073110"
)

var logLevel format.LogLevel

func newDate() string {
   return time.Now().Format(time.RFC1123)
}

func newHmacSha(src, key []byte) []byte {
   dst := hmac.New(sha256.New, key)
   dst.Write(src)
   return dst.Sum(nil)
}

func newSha(src io.Reader) []byte {
   dst := sha256.New()
   io.Copy(dst, src)
   return dst.Sum(nil)
}

type credentials struct {
   Resource struct {
      AccessKeyID string
      SecretAccessKey string
      SessionToken string
   }
}

func newCredentials() (*credentials, error) {
   body := map[string]string{"appKey": appKey}
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/authentication/credentials/temporary/android850", buf,
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

func (c credentials) Gallery(rgconst string) (*http.Response, error) {
   var buf strings.Builder
   buf.WriteString(origin)
   buf.WriteString("/template/imdb-android-writable")
   buf.WriteString("/8.5.runway-gallery-images.jstl/render")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "rgconst=" + url.QueryEscape(rgconst)
   req.Header = http.Header{
      "X-Amz-Date": {newDate()},
      "X-Amz-Security-Token": {c.Resource.SessionToken},
      "X-Amzn-Sessionid": {sessionID},
   }
   c.sign(req)
   logLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func (c credentials) sign(req *http.Request) {
   buf := new(bytes.Buffer)
   buf.WriteString(req.Method)
   buf.WriteByte('\n')
   buf.WriteString(req.URL.Path)
   buf.WriteByte('\n')
   buf.WriteString(req.URL.RawQuery)
   buf.WriteString("\nhost:")
   buf.WriteString(req.URL.Host)
   buf.WriteString("\nx-amz-date:")
   buf.WriteString(req.Header.Get("X-Amz-Date"))
   buf.WriteString("\nx-amz-security-token:")
   buf.WriteString(req.Header.Get("X-Amz-Security-Token"))
   buf.WriteString("\nx-amzn-sessionid:")
   buf.WriteString(req.Header.Get("X-Amzn-Sessionid"))
   buf.WriteString("\n\n")
   sha := newSha(buf)
   hmacSha := newHmacSha(sha, []byte(c.Resource.SecretAccessKey))
   buf.WriteString("AWS3 AWSAccessKeyId=")
   buf.WriteString(c.Resource.AccessKeyID)
   buf.WriteString(",Algorithm=HmacSHA256,Signature=")
   buf.WriteString(base64.StdEncoding.EncodeToString(hmacSha))
   buf.WriteString(",SignedHeaders=")
   buf.WriteString("host;x-amz-date;x-amz-security-token;x-amzn-sessionid")
   req.Header.Set("x-amzn-authorization", buf.String())
}
