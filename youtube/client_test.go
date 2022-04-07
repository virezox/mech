package youtube

import (
   "fmt"
   "testing"
   "time"
)

/*
fail 63RmMXCd_bQ
fail 6SJNVb0GnPI
fail Cr381pDsSsA
fail CsmdDsKjzN8
fail DJztXj2GPfl
fail HsUATh_Nc2U
fail HtVdAasjOgU
fail Ms7iBXnlUO8
fail Q39EVAstoRM
fail SZJvDhaSDnc
fail Tq92D6wQ1mg
fail V36LpHqtcDY
fail WaOKSUlf4TM
fail i1Ko8UG-Tdo
fail nGC3D_FkCmg
fail qEJwOuvDf7I
fail s7_qI6_mIXc
fail sJL6WA-aGkQ
fail yYr8q0y5Jfg
fail yZIXLfi8CZQ
*/
var videoIDs = []string{
   "1t24XAntNCY",
   "2NUZ8W2llS4",
   "3b8nCWDgZ6Q",
   "5KLPxDtMqe8",
   "63RmMXCd_bQ",
   "6SJNVb0GnPI",
   "9lWxNJF-ufM",
   "BGQWPY4IigY",
   "BaW_jenozKc",
   "CHqg6qOn4no",
   "Cr381pDsSsA",
   "CsmdDsKjzN8",
   "DJztXj2GPfl",
   "FIl7x6_3R5Y",
   "FRhJzUSJbGI",
   "FlRa-iH7PGw",
   "HsUATh_Nc2U",
   "HtVdAasjOgU",
   "IB3lcPjvWLA",
   "M4gD1WSo5mA",
   "MeJVWBSsPAY",
   "MgNrAu2pzNs",
   "Ms7iBXnlUO8",
   "OtqTfy26tG0",
   "Q39EVAstoRM",
   "SZJvDhaSDnc",
   "Tq92D6wQ1mg",
   "V36LpHqtcDY",
   "WaOKSUlf4TM",
   "XclachpHxis",
   "YOelRv7fMxY",
   "Yh0AhrY9GjA",
   "Z4Vy8R84T1U",
   "__2ABJjxzNo",
   "_b-2C3KPAM0",
   "a9LDPn-MO4I",
   "cBvYw8_A0vQ",
   "eQcmzGIKrzg",
   "gVfLd0zydlo",
   "gVfgbahppCY",
   "i1Ko8UG-Tdo",
   "iqKdEhx-dD4",
   "jvGDaLqkpTg",
   "kgx4WGK0oNU",
   "lqQg6PlCWgI",
   "lsguqyKfVQg",
   "mzZzzBU6lrM",
   "nGC3D_FkCmg",
   "qEJwOuvDf7I",
   "s7_qI6_mIXc",
   "sJL6WA-aGkQ",
   "wsQiKKfKxug",
   "x41yOUIvK2k",
   "yYr8q0y5Jfg",
   "yZIXLfi8CZQ",
}

func TestPlayer(t *testing.T) {
   for _, videoID := range videoIDs {
      play, err := Android.Player(videoID)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status == "OK" {
         fmt.Println("pass", videoID)
      } else {
         fmt.Println("fail", videoID)
      }
      time.Sleep(time.Second)
   }
}

func TestSearch(t *testing.T) {
   search, err := Mweb.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.CompactVideoRenderer)
   }
}
