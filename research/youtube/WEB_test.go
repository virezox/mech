package youtube

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "os"
   "testing"
)

func TestWebUnplugged(t *testing.T) {
   const name = "WEB_UNPLUGGED"
   version, err := newVersion("https://www.youtube.com/embed/MIchMEqVwvg", "")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestWebEmbeddedPlayer(t *testing.T) {
   const name = "WEB_EMBEDDED_PLAYER"
   version, err := newVersion("https://www.youtube.com/embed/MIchMEqVwvg", "")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestWebRemix(t *testing.T) {
   const name = "WEB_REMIX"
   version, err := newVersion("https://music.youtube.com", "Firefox/99")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestWeb(t *testing.T) {
   const name = "WEB"
   version, err := newVersion("https://www.youtube.com", "")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func TestWebKids(t *testing.T) {
   const name = "WEB_KIDS"
   version, err := newVersion("https://www.youtubekids.com", "Firefox/99")
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}

func webCreator() (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   tok, err := format.Open[token](cache, "mech/youtube.json")
   if err != nil {
      return "", err
   }
   req, err := http.NewRequest("GET", "https://studio.youtube.com", nil)
   if err != nil {
      return "", err
   }
   req.URL.RawQuery = "approve_browser_access=true"
   req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   sep := []byte(`"client":`)
   var client struct {
      ClientVersion string
   }
   if err := json.Decode(res.Body, sep, &client); err != nil {
      return "", err
   }
   return client.ClientVersion, nil
}

func TestWebCreator(t *testing.T) {
   const name = "WEB_CREATOR"
   version, err := webCreator()
   if err != nil {
      t.Fatal(err)
   }
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
