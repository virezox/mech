package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
)

func (d downloader) DASH(video, audio int64) error {
   if d.info {
      fmt.Println(d.Content)
   }
   video_dash := d.Content.DASH()
   res, err := roku.Client.Redirect(nil).Get(video_dash.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = res.Request.URL
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   reps := d.media.Representations().Codecs("mp4a")
   if err := d.download(audio, reps); err != nil {
      return err
   }
   reps = d.media.Representations().Codecs("avc1")
   return d.download(video, reps)
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   play, err := site.Playback(d.Meta.ID)
   if err != nil {
      return err
   }
}
