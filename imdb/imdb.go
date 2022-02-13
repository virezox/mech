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
   "strconv"
   "strings"
   "time"
)

const (
   appKey = "4f833099-e4fe-4912-80f3-b1b169097914"
   origin = "https://api.imdbws.com"
)

var LogLevel format.LogLevel

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

type Credential struct {
   Resource struct {
      AccessKeyID string
      SecretAccessKey string
      SessionToken string
   }
}

func NewCredential() (*Credential, error) {
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   cred := new(Credential)
   if err := json.NewDecoder(res.Body).Decode(cred); err != nil {
      return nil, err
   }
   return cred, nil
}

func (c Credential) Gallery(rgconst string) (*Gallery, error) {
   var buf strings.Builder
   buf.WriteString(origin)
   buf.WriteString("/template/imdb-android-writable")
   buf.WriteString("/8.5.runway-gallery-images.jstl/render")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "rgconst=" + url.QueryEscape(rgconst)
   // Since we arent using `Set`, the case on these needs to be correct:
   req.Header = http.Header{
      "X-Amz-Date": {newDate()},
      "X-Amz-Security-Token": {c.Resource.SessionToken},
   }
   c.sign(req)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   gal := new(Gallery)
   if err := json.NewDecoder(res.Body).Decode(gal); err != nil {
      return nil, err
   }
   return gal, nil
}

func (c Credential) sign(req *http.Request) {
   buf := new(bytes.Buffer)
   buf.WriteString(req.Method)
   buf.WriteByte('\n')
   buf.WriteString(req.URL.Path)
   buf.WriteByte('\n')
   buf.WriteString(req.URL.RawQuery)
   buf.WriteString("\nhost:")
   buf.WriteString(req.URL.Host)
   buf.WriteString("\nx-amz-date:")
   buf.WriteString(req.Header.Get("x-amz-date"))
   buf.WriteString("\nx-amz-security-token:")
   buf.WriteString(req.Header.Get("x-amz-security-token"))
   buf.WriteString("\n\n")
   // Yes, it is stupid to SHA256 twice, but that is what the server requires.
   sha := newSha(buf)
   hmacSha := newHmacSha(sha, []byte(c.Resource.SecretAccessKey))
   // All of this is needed:
   buf.WriteString("AWS3 AWSAccessKeyId=")
   buf.WriteString(c.Resource.AccessKeyID)
   buf.WriteString(",Algorithm=HmacSHA256,Signature=")
   buf.WriteString(base64.StdEncoding.EncodeToString(hmacSha))
   buf.WriteString(",SignedHeaders=host;x-amz-date;x-amz-security-token")
   req.Header.Set("x-amzn-authorization", buf.String())
}

type Image struct {
   URL string
}

func (i Image) Format(width int64) string {
   n := strings.LastIndexByte(i.URL, '.')
   if n == -1 {
      return i.URL
   }
   buf := []byte(i.URL[:n])
   buf = append(buf, "UX"...)
   buf = strconv.AppendInt(buf, width, 10)
   buf = append(buf, i.URL[n:]...)
   return string(buf)
}

type Gallery struct {
   Images []Image
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
