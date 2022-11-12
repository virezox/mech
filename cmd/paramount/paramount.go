package main

import (
   "github.com/89z/mech/paramount"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/hls"
   "strings"
   "strconv"
)

func (f flags) DASH(preview *paramount.Preview) error {
   var err error
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   f.Name = preview.Name()
   reps, err := f.Stream.DASH(paramount.DASH(f.guid))
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MimeType != "audio/mp4" {
         return false
      }
      if r.Role() == "description" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.HasPrefix(b.Adaptation.Lang, f.lang) {
         return false
      }
      if !strings.HasPrefix(b.Codecs, f.codecs) {
         return false
      }
      return true
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   index = video.Index(func(a, b dash.Representation) bool {
      return b.Height == f.height
   })
   return f.DASH_Get(video, index)
}

func (f flags) HLS(preview *paramount.Preview) error {
   f.Name = preview.Name()
   master, err := f.Stream.HLS(paramount.HLS(f.guid))
   if err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   index := streams.Index(func(a, b hls.Stream) bool {
      return strings.HasSuffix(b.Resolution, strconv.FormatInt(f.height, 10))
   })
   return f.HLS_Streams(streams, index)
}
