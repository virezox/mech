package instagram

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "os"
)

const Origin = "https://www.instagram.com"

var Verbose bool

type Sidecar struct {
   GraphQL struct {
      Shortcode_Media struct {
         Edge_Sidecar_To_Children struct {
            Edges []struct {
               Node struct {
                  Display_URL string
               }
            }
         }
      }
   }
}

func NewSidecar(id string) (*Sidecar, error) {
   req, err := http.NewRequest("GET", Origin + "/p/" + id + "/", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("__a", "1")
   req.URL.RawQuery = q.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   if Verbose {
      d, err := httputil.DumpRequest(req, false)
      if err != nil {
         return nil, err
      }
      os.Stdout.Write(d)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   car := new(Sidecar)
   if err := json.NewDecoder(res.Body).Decode(car); err != nil {
      return nil, err
   }
   return car, nil
}
