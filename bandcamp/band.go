package bandcamp
import "net/http"

type Band struct {
   Band_ID int
   URL string
}

func BandID(id string) (*Band, error) {
   req, err := http.NewRequest("GET", Origin + "/api/band/3/info", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("band_id", id)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   ban := new(Band)
   if err := roundTrip(req, ban); err != nil {
      return nil, err
   }
   return ban, nil
}

// band_url fails with API 3
func BandURL(addr string) (*Band, error) {
   req, err := http.NewRequest("GET", Origin + "/api/band/2/info", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("band_url", addr)
   q.Set("key", key)
   req.URL.RawQuery = q.Encode()
   ban := new(Band)
   if err := roundTrip(req, ban); err != nil {
      return nil, err
   }
   return ban, nil
}
