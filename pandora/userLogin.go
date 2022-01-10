package pandora

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/net"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "strconv"
   "strings"
)

func ID(addr string) (string, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return "", err
   }
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   for _, node := range net.ReadHTML(res.Body, "meta") {
      if node.Attr["property"] == "al:android:url" {
         con := node.Attr["content"]
         addr, err := url.Parse(con)
         if err != nil {
            return "", err
         }
         return addr.Query().Get("pandoraId"), nil
      }
   }
   return "", notFound{"al:android:url"}
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
func (p PlaybackInfo) Base() string {
   if p.Result == nil {
      return ""
   }
   addr, err := url.Parse(p.Result.AudioUrlMap.HighQuality.AudioUrl)
   if err != nil {
      return ""
   }
   return filepath.Base(addr.Path)
}

type UserLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

func OpenUserLogin(name string) (*UserLogin, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   user := new(UserLogin)
   if err := json.NewDecoder(file).Decode(user); err != nil {
      return nil, err
   }
   return user, nil
}

func (u UserLogin) Create(name string) error {
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   return enc.Encode(u)
}

func (u UserLogin) PlaybackInfo(id string) (*PlaybackInfo, error) {
   rInfo := playbackInfoRequest{
      IncludeAudioToken: true,
      PandoraID: id,
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   }
   buf, err := json.Marshal(rInfo)
   if err != nil {
      return nil, err
   }
   body := Cipher(buf).Pad().Encrypt().Encode()
   req, err := http.NewRequest(
      "POST", origin + "/services/json/", strings.NewReader(body),
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
   format.Log.Dump(req)
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
   rValue := valueExchangeRequest{
      OfferName: "premium_access",
      SyncTime: syncTime,
      UserAuthToken: u.Result.UserAuthToken,
   }
   buf, err := json.Marshal(rValue)
   if err != nil {
      return err
   }
   body := Cipher(buf).Pad().Encrypt().Encode()
   req, err := http.NewRequest(
      "POST", origin + "/services/json/", strings.NewReader(body),
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
   format.Log.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
