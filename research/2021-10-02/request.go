package goinsta

import (
   "bytes"
   "compress/gzip"
   "crypto/md5"
   "crypto/rand"
   "encoding/hex"
   "errors"
   "fmt"
   "io"
   "io/ioutil"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
   mathRand "math/rand"
)

type reqOptions struct {
	// Connection is connection header. Default is "close".
	Connection string

	// Endpoint is the request path of instagram api
	Endpoint string

	// Omit API omit the /api/v1/ part of the url
	OmitAPI bool

	// IsPost set to true will send request with POST method.
	//
	// By default this option is false.
	IsPost bool

	// Compress post form data with gzip
	Gzip bool

	// UseV2 is set when API endpoint uses v2 url.
	UseV2 bool

	// Use b.i.instagram.com
	Useb bool

	// Query is the parameters of the request
	//
	// This parameters are independents of the request method (POST|GET)
	Query map[string]string

	// DataBytes can be used to pass raw data to a request, instead of a
	//   form using the Query param. This is used for e.g. photo and vieo uploads.
	DataBytes *bytes.Buffer

	// List of headers to ignore
	IgnoreHeaders []string

	// Extra headers to add
	ExtraHeaders map[string]string

	// Timestamp
	Timestamp string
}


func (insta *Instagram) sendRequest(o *reqOptions) (body []byte, h http.Header, err error) {
   if insta == nil {
      return nil, nil, errors.New(
         "insta has not been defined, this is most likely a bug in the code. " +
         "Please backtrack which call this error came from, and open an issue " +
         "detailing exactly how you got to this error",
      )
   }
   method := "GET"
   if o.IsPost {
      method = "POST"
   }
   if o.Connection == "" {
      o.Connection = "close"
   }
   if o.Timestamp == "" {
      o.Timestamp = strconv.Itoa(int(time.Now().Unix()))
   }
   var nu string
   if o.Useb {
      nu = instaAPIUrlb
   } else {
      nu = instaAPIUrl
   }
   if o.UseV2 && !o.Useb {
      nu = instaAPIUrlv2
   } else if o.UseV2 && o.Useb {
      nu = instaAPIUrlv2b
   }
   if o.OmitAPI {
      nu = baseUrl
      o.IgnoreHeaders = append(o.IgnoreHeaders, omitAPIHeadersExclude...)
   }
   u, err := url.Parse(nu + o.Endpoint)
   if err != nil {
      return nil, nil, err
   }
   vs := url.Values{}
   bf := bytes.NewBuffer([]byte{})
   reqData := bytes.NewBuffer([]byte{})
   for k, v := range o.Query {
      vs.Add(k, v)
   }
   // If DataBytes has been passed, use that as data, else use Query
   if o.DataBytes != nil {
      reqData = o.DataBytes
   } else {
      reqData.WriteString(vs.Encode())
   }
   var contentEncoding string
   if o.IsPost && o.Gzip {
      // If gzip encoding needs to be applied
      zw := gzip.NewWriter(bf)
      defer zw.Close()
      if _, err := zw.Write(reqData.Bytes()); err != nil {
         return nil, nil, err
      }
      if err := zw.Close(); err != nil {
         return nil, nil, err
      }
      contentEncoding = "gzip"
   } else if o.IsPost {
      // use post form if POST request
      bf = reqData
   } else {
      // append query to url if GET request
      for k, v := range u.Query() {
         vs.Add(k, strings.Join(v, " "))
      }
      u.RawQuery = vs.Encode()
   }
   var req *http.Request
   req, err = http.NewRequest(method, u.String(), bf)
   if err != nil {
      return
   }
   ignoreHeader := func(h string) bool {
      for _, k := range o.IgnoreHeaders {
         if k == h {
            return true
         }
      }
      return false
   }
   setHeaders := func(h map[string]string) {
      for k, v := range h {
         if v != "" && !ignoreHeader(k) {
            req.Header.Set(k, v)
         }
      }
   }
   headers := map[string]string{
      "Accept-Encoding":             "gzip,deflate",
      "Accept-Language":             locale,
      "Connection":                  o.Connection,
      "Content-Type":                "application/x-www-form-urlencoded; charset=UTF-8",
      "User-Agent":                  insta.userAgent,
      "X-Bloks-Is-Layout-Rtl":       "false",
      "X-Bloks-Is-Panorama-Enabled": "true",
      "X-Bloks-Version-Id":          bloksVerID,
      "X-Fb-Client-Ip":              "True",
      "X-Fb-Http-Engine":            "Liger",
      "X-Fb-Server-Cluster":         "True",
      "X-Ig-Android-Id":             insta.dID,
      "X-Ig-App-Id":                 fbAnalytics,
      "X-Ig-App-Locale":             locale,
      "X-Ig-App-Startup-Country":    "unkown",
      "X-Ig-Bandwidth-Speed-KBPS":   fmt.Sprintf("%d.000", random(1000, 9000)),
      "X-Ig-Bandwidth-TotalBytes-B": strconv.Itoa(random(1000000, 5000000)),
      "X-Ig-Bandwidth-Totaltime-Ms": strconv.Itoa(random(200, 800)),
      "X-Ig-Capabilities":           igCapabilities,
      "X-Ig-Connection-Type":        connType,
      "X-Ig-Device-Id":              insta.uuid,
      "X-Ig-Device-Locale":          locale,
      "X-Ig-Family-Device-Id":       insta.fID,
      "X-Ig-Mapped-Locale":          locale,
      "X-Ig-Timezone-Offset":        timeOffset,
      "X-Pigeon-Rawclienttime":      fmt.Sprintf("%s.%d", o.Timestamp, random(100, 900)),
      "X-Pigeon-Session-Id":         insta.psID,
   }
   if insta.Account != nil {
      req.Header.Set("Ig-Intended-User-Id", strconv.Itoa(int(insta.Account.ID)))
   } else {
      req.Header.Set("Ig-Intended-User-Id", "0")
   }
   if contentEncoding != "" {
      headers["Content-Encoding"] = contentEncoding
   }
   setHeaders(headers)
   setHeaders(o.ExtraHeaders)
   for key, value := range insta.headerOptions {
      if value != "" && !ignoreHeader(key) {
         req.Header.Set(key, value)
      }
   }
   dum, err := httputil.DumpRequest(req, true)
   if err != nil {
      return nil, nil, err
   }
   os.Stdout.Write(append(dum, '\n'))
   resp, err := insta.c.Do(req)
   if err != nil {
      return nil, nil, err
   }
   defer resp.Body.Close()
   body, err = ioutil.ReadAll(resp.Body)
   if err != nil {
      return nil, nil, err
   }
   insta.extractHeaders(resp.Header)
   // Decode gzip encoded responses
   encoding := resp.Header.Get("Content-Encoding")
   if encoding != "" && encoding == "gzip" {
      buf := bytes.NewBuffer(body)
      zr, err := gzip.NewReader(buf)
      if err != nil {
         return nil, nil, err
      }
      body, err = ioutil.ReadAll(zr)
      if err != nil {
         return nil, nil, err
      }
      if err := zr.Close(); err != nil {
         return nil, nil, err
      }
   }
   return body, resp.Header.Clone(), err
}

func (insta *Instagram) extractHeaders(h http.Header) {
   extract := func(in string, out string) {
      x := h[in]
      if len(x) > 0 && x[0] != "" {
         // prevent from auth being set without token post login
         if in == "Ig-Set-Authorization" {
            old, ok := insta.headerOptions[out]
            if ok && len(old) != 0 {
               current := strings.Split(old, ":")
               newHeader := strings.Split(x[0], ":")
               if len(current[2]) > len(newHeader[2]) {
                  return
               }
            }
         }
         insta.headerOptions[out] = x[0]
      }
   }
   extract("Ig-Set-Authorization", "Authorization")
   extract("Ig-Set-Ig-U-Ds-User-Id", "Ig-U-Ds-User-Id")
   extract("Ig-Set-Ig-U-Ig-Direct-Region-Hint", "Ig-U-Ig-Direct-Region-Hint")
   extract("Ig-Set-Ig-U-Rur", "Ig-U-Rur")
   extract("Ig-Set-Ig-U-Shbid", "Ig-U-Shbid")
   extract("Ig-Set-Ig-U-Shbts", "Ig-U-Shbts")
   extract("Ig-Set-X-Mid", "X-Mid")
   extract("X-Ig-Set-Www-Claim", "X-Ig-Www-Claim")
}

func random(min, max int) int {
	mathRand.Seed(time.Now().UnixNano())
	return mathRand.Intn(max-min) + min
}

var timeOffset = getTimeOffset()

func getTimeOffset() string {
   _, offset := time.Now().Zone()
   return strconv.Itoa(offset)
}

const volatileSeed = "12345"

func generateMD5Hash(text string) string {
   hasher := md5.New()
   hasher.Write([]byte(text))
   return hex.EncodeToString(hasher.Sum(nil))
}

func generateDeviceID(seed string) string {
   hash := generateMD5Hash(seed + volatileSeed)
   return "android-" + hash[:16]
}

func generateSignature(d interface{}, extra ...map[string]string) map[string]string {
   var data string
   switch x := d.(type) {
   case []byte:
      data = string(x)
   case string:
      data = x
   }
   r := map[string]string{"signed_body": "SIGNATURE." + data}
   for _, e := range extra {
      for k, v := range e {
         r[k] = v
      }
   }
   return r
}

func newUUID() (string, error) {
   uuid := make([]byte, 16)
   n, err := io.ReadFull(rand.Reader, uuid)
   if n != len(uuid) || err != nil {
      return "", err
   }
   // variant bits; see section 4.1.1
   uuid[8] = uuid[8]&^0xc0 | 0x80
   // version 4 (pseudo-random); see section 4.1.3
   uuid[6] = uuid[6]&^0xf0 | 0x40
   return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func generateUUID() string {
   uuid, err := newUUID()
   if err != nil {
      // default value when error occurred
      return "cb479ee7-a50d-49e7-8b7b-60cc1a105e22"
   }
   return uuid
}
