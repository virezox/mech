package bandcamp

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "io"
   "net/http"
)

var Log_Level format.Log_Level

type Params struct {
   A_ID int
   I_ID int
   I_Type string
}

func New_Params(addr string) (*Params, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   Log_Level.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan xml.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte(`<p id="report-account-vm"`)
   scan.Scan()
   var p struct {
      DataTouReportParams []byte `xml:"data-tou-report-params,attr"`
   }
   if err := scan.Decode(&p); err != nil {
      return nil, err
   }
   param := new(Params)
   if err := json.Unmarshal(p.DataTouReportParams, param); err != nil {
      return nil, err
   }
   return param, nil
}

func (p Params) Band() (*Band, error) {
   return new_band(p.A_ID)
}

func (p Params) Tralbum() (*Tralbum, error) {
   switch p.I_Type {
   case "a":
      return newTralbum('a', p.I_ID)
   case "t":
      return newTralbum('t', p.I_ID)
   }
   return nil, invalidType{p.I_Type}
}
