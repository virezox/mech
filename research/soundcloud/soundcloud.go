package soundcloud

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
)

var LogLevel format.LogLevel

type Track struct {
   Display_Date string // 2021-04-12T07:00:01Z
   ID int
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
   Title string
   Artwork_URL string
   User struct {
      Avatar_URL string
      Username string
   }
}

type UserStream struct {
   Collection []struct {
      Track Track
   }
}

func NewUserStream(user int64) (*UserStream, error) {
   buf := []byte("https://api-v2.soundcloud.com/stream/users/")
   buf = strconv.AppendInt(buf, user, 10)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stream := new(UserStream)
   if err := json.NewDecoder(res.Body).Decode(stream); err != nil {
      return nil, err
   }
   return stream, nil
}
