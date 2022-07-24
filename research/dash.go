package mech

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/os"
   "io"
)

func Presentation(stream Streamer, filters ...dash_filter) error {
   mpd, err := client.Redirect(nil).Get(stream.Address())
   if err != nil {
      return err
   }
   defer mpd.Body.Close()
   var pre dash.Presentation
   if err := xml.NewDecoder(mpd.Body).Decode(&pre); err != nil {
      return err
   }
   reps := pre.Representation()
   for _, filter := range filters {
      reps, index := filter(reps)
      if stream.Info() {
         for i, rep := range reps {
            if i == index {
               fmt.Print("!")
            }
            fmt.Println(rep)
         }
         continue
      }
      rep := reps[index]
      file, err := os.Create(stream.Name() + rep.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      req, err := http.NewRequest("GET", rep.Initialization(), nil)
      if err != nil {
         return err
      }
      req.URL = mpd.Request.URL.ResolveReference(req.URL)
      initial, err := client.Redirect(nil).Do(req)
      if err != nil {
         return err
      }
      defer initial.Body.Close()
      refs := rep.Media()
      pro := os.Progress_Chunks(file, len(refs))
      dec := mp4.New_Decrypt(pro)
      var key []byte
      if rep.ContentProtection != nil {
         private_key, err := os.ReadFile(stream.Private_Key())
         if err != nil {
            return err
         }
         client_ID, err := os.ReadFile(stream.Client_ID())
         if err != nil {
            return err
         }
         key_ID, err := widevine.Key_ID(rep.ContentProtection.Default_KID)
         if err != nil {
            return err
         }
         mod, err := widevine.New_Module(private_key, client_ID, key_ID)
         if err != nil {
            return err
         }
         keys, err := mod.Post(stream)
         if err != nil {
            return err
         }
         key = keys.Content().Key
         if err := dec.Init(initial.Body); err != nil {
            return err
         }
      } else {
         _, err := io.Copy(pro, initial.Body)
         if err != nil {
            return err
         }
      }
      for _, ref := range refs {
         req, err := http.NewRequest("GET", ref, nil)
         if err != nil {
            return err
         }
         req.URL = mpd.Request.URL.ResolveReference(req.URL)
         media, err := client.Redirect(nil).Level(0).Do(req)
         if err != nil {
            return err
         }
         pro.Add_Chunk(media.ContentLength)
         if rep.ContentProtection != nil {
            err = dec.Segment(media.Body, key)
         } else {
            _, err = io.Copy(pro, media.Body)
         }
         if err != nil {
            return err
         }
         if err := media.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}

var client = http.Default_Client

type Streamer interface {
   Address() string
   Client_ID() string
   Info() bool
   Name() string
   Private_Key() string
   widevine.Poster
}

type dash_filter func([]dash.Representation) ([]dash.Representation, int)
