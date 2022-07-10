package main

import (
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/apple.json")
   if err != nil {
      t.Fatal(err)
   }
   env, err := New_Environment()
   if err != nil {
      t.Fatal(err)
   }
   episode, err := New_Episode(content_ID)
   if err != nil {
      t.Fatal(err)
   }
   Poster{auth, env, episode, pssh}
   master, err := f.Flag.HLS(video.URL, content.Base())
   if err != nil {
      return err
   }
   streams := master.Streams
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}
