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

func setVerbose() {
   paramount.LogLevel = 1
}

func (d *downloader) setBase() error {
   preview, err := paramount.NewMedia(d.mediaID).Preview()
   if err != nil {
      return err
   }
   d.base = preview.Base()
}
