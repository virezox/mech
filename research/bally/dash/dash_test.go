package dash

import (
   "encoding/hex"
   "github.com/89z/rosso/mp4"
   "github.com/89z/rosso/http"
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

var client = http.Default_Client

func Test_DASH(t *testing.T) {
   file, err := os.Create(".mp4")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   dec := mp4.New_Decrypt(file)
   key, err := hex.DecodeString("f4a5865c0f0f4523e0f055b66fb3b285")
   if err != nil {
      t.Fatal(err)
   }
   for i, ref := range refs {
      req, err := http.NewRequest("GET", ref, nil)
      if err != nil {
         t.Fatal(err)
      }
      res, err := client.Redirect(nil).Do(req)
      if err != nil {
         t.Fatal(err)
      }
      if i == 0 {
         err = dec.Init(res.Body)
      } else {
         err = dec.Segment(res.Body, key)
      }
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
   }
}
