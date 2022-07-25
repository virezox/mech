package dash

import (
   "encoding/xml"
   "fmt"
   "os"
   "testing"
)

var refs = []string{
   // video 0
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_init.m4i",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101732240460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101772280460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101812320460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101852360460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101892400460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101932440460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-101972480460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102012520460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102052560460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102092600460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102132640460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102172680460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102212720460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102252760460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102292800460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102332840460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102372880460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102412920460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102452960460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102493000460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102533040460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102573080460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102613120460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102653160460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102693200460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102733240460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102773280460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102813320460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102853360460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102893400460.m4v",
   // video 1
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102933440460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-102973480460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-103013520460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-103053560460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-103093600460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-103133640460.m4v",
   "https://live-d-02-ballysports.akamaized.net/Content/DASH_DASH/Live/channel(sinc-fs-prime-ticket-2042)/1658775954314item-01item_Segment-103173680460.m4v",
}

func Test_DASH(t *testing.T) {
   file, err := os.Create(s.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", item.Initialization(), nil)
   if err != nil {
      return err
   }
   req.URL = s.base.ResolveReference(req.URL)
   res, err := client.Redirect(nil).Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   media := item.Media()
   pro := os.Progress_Chunks(file, len(media))
   dec := mp4.New_Decrypt(pro)
   var key []byte
   if item.ContentProtection != nil {
      private_key, err := os.ReadFile(s.Private_Key)
      if err != nil {
         return err
      }
      client_ID, err := os.ReadFile(s.Client_ID)
      if err != nil {
         return err
      }
      key_ID, err := widevine.Key_ID(item.ContentProtection.Default_KID)
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
      if err := dec.Init(res.Body); err != nil {
         return err
      }
   } else {
      _, err := io.Copy(pro, res.Body)
      if err != nil {
         return err
      }
   }
   for _, ref := range media {
      req, err := http.NewRequest("GET", ref, nil)
      if err != nil {
         return err
      }
      req.URL = s.base.ResolveReference(req.URL)
      res, err := client.Redirect(nil).Level(0).Do(req)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if item.ContentProtection != nil {
         err = dec.Segment(res.Body, key)
      } else {
         _, err = io.Copy(pro, res.Body)
      }
      if err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
}
