package mech

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/research/dash"
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
)

var client = http.Default_Client

type Stream struct {
   Address string
   Client_ID string
   Name string
   Poster widevine.Poster
   Private_Key string
}

type dash_filter func([]dash.Representation) ([]dash.Representation, int)

func (s Stream) Decode(filters ...dash_filter) error {
   mpd, err := client.Redirect(nil).Get(s.Address)
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
      if s.Name != "" {
         rep := reps[index]
         err := s.get(mpd.Request.URL, rep)
         if err != nil {
            return err
         }
      } else {
         for i, rep := range reps {
            if i == index {
               fmt.Print("!")
            }
            fmt.Println(rep)
         }
      }
   }
   return nil
}

func (s Stream) get(base *url.URL, rep dash.Representation) error {
   file, err := os.Create(s.Name + rep.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", rep.Initialization(), nil)
   if err != nil {
      return err
   }
   req.URL = base.ResolveReference(req.URL)
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
      private_key, err := os.ReadFile(s.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(s.Client_ID)
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
      keys, err := mod.Post(s.Poster)
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
      req.URL = base.ResolveReference(req.URL)
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
   return nil
}
