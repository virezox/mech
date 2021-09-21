package instagram

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
)

const Origin = "https://www.instagram.com"

var Verbose bool

// instagram.com/p/CT-cnxGhvvO
func ValidID(id string) error {
   if len(id) == 11 {
      return nil
   }
   return fmt.Errorf("%q invalid as ID", id)
}

type Edge struct {
   Node struct {
      Display_URL string
   }
}

type Sidecar struct {
   GraphQL struct {
      Shortcode_Media struct {
         Edge_Sidecar_To_Children struct {
            Edges []Edge
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
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   car := new(Sidecar)
   if err := json.NewDecoder(res.Body).Decode(car); err != nil {
      return nil, err
   }
   return car, nil
}

func (s Sidecar) Edges() []Edge {
   return s.GraphQL.Shortcode_Media.Edge_Sidecar_To_Children.Edges
}
