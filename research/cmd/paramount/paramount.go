package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/format/hls"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func (d *download) setKey(privateKey, clientID, keyID []byte) error {
   sess, err := paramount.NewSession(d.mediaID)
   if err != nil {
      return err
   }
   d.key, err = sess.Key(privateKey, clientID, kID)
   if err != nil {
      return err
   }
   return nil
}

func (d *download) setBase() error {
   preview, err := paramount.NewMedia(d.mediaID).Preview()
   if err != nil {
      return err
   }
   d.base = preview.Base()
}

func (d download) hls_url() (*url.URL, error) {
   return paramount.NewMedia(d.mediaID).HLS()
}

func setVerbose() {
   paramount.LogLevel = 1
}
