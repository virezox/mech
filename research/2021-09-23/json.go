package instagram

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

const body = `
{
   "query_hash": "8c2a529969ee035a5063f2fc8602a0fd",
   "variables": {"id":"294582047","first":1}
}
`

/*
curl -v -A 'Instagram 1.1.1' https://i.instagram.com/api/v1/users/181306552/info/
*/

// severe rate limit
func jsonGraphQL(id string) error {
   req, err := http.NewRequest(
      "POST", origin + "/graphql/query/", strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0)"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

// severe rate limit
func jsonP(id string) error {
   req, err := http.NewRequest("GET", origin + "/p/" + id + "/", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("__a", "1")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

// severe rate limit
func jsonTV(id string) error {
   req, err := http.NewRequest("GET", origin + "/tv/" + id + "/", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("__a", "1")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}
