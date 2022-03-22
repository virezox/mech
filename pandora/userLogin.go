package pandora

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "path/filepath"
   "strconv"
   "strings"
)

func OpenUserLogin(elem ...string) (*UserLogin, error) {
   return format.Open[UserLogin](elem...)
}

func (u UserLogin) Create(elem ...string) error {
   return format.Create(u, elem...)
}

var LogLevel format.LogLevel

type Details struct {
   Result struct {
      Annotations map[string]struct {
         ArtistName string
         Name string
      }
   }
}

func NewDetails(id string) (*Details, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{"pandoraId": id})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/services/json/", buf)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "method=catalog.v4.getDetails"
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   det := new(Details)
   if err := json.NewDecoder(res.Body).Decode(det); err != nil {
      return nil, err
   }
   return det, nil
}

type PlaybackInfo struct {
   Stat string
   Result *struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioUrl string
         }
      }
   }
}

// audio-dc6-t3-1-v4v6.pandora.com/access/3648302390726192234.mp3?version=5
func (p PlaybackInfo) Ext() (string, error) {
   if p.Result != nil {
      addr, err := url.Parse(p.Result.AudioUrlMap.HighQuality.AudioUrl)
      if err != nil {
         return "", err
      }
      return filepath.Ext(addr.Path), nil
   }
   return "", notPresent{"Result"}
}

type UserLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

func (u UserLogin) PlaybackInfo(id string) (*PlaybackInfo, error) {
   src, err := json.Marshal(playbackInfoRequest{
      IncludeAudioToken: true,
      PandoraID: id,
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   })
   if err != nil {
      return nil, err
   }
   dst, err := Cipher.Pad(src).Encrypt()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(dst)),
   )
   if err != nil {
      return nil, err
   }
   // auth_token and user_Id can be empty, but they must be included
   req.URL.RawQuery = url.Values{
      "auth_token": {""},
      "method": {"onDemand.getAudioPlaybackInfo"},
      "partner_id": {"42"},
      "user_id": {""},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   info := new(PlaybackInfo)
   if err := json.NewDecoder(res.Body).Decode(info); err != nil {
      return nil, err
   }
   return info, nil
}

// Token is good for 30 minutes.
func (u UserLogin) ValueExchange() error {
   src, err := json.Marshal(valueExchangeRequest{
      OfferName: "premium_access",
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   })
   if err != nil {
      return err
   }
   dst, err := Cipher.Pad(src).Encrypt()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(dst)),
   )
   if err != nil {
      return err
   }
   // auth_token and user_Id can be empty, but they must be included
   req.URL.RawQuery = url.Values{
      "auth_token": {""},
      "method": {"user.startValueExchange"},
      "partner_id": {"42"},
      "user_id": {""},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return strconv.Quote(n.value) + " is not present"
}
