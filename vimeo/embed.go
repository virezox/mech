package vimeo

import (
   "github.com/89z/rosso/json"
   "io"
   "strconv"
)

func (c Config) String() string {
   b := []byte("ID: ")
   b = strconv.AppendInt(b, c.Video.ID, 10)
   b = append(b, "\nTitle: "...)
   b = append(b, c.Video.Title...)
   b = append(b, "\nDate: "...)
   b = append(b, c.SEO.Upload_Date...)
   for _, pro := range c.Request.Files.Progressive {
      b = append(b, '\n')
      b = append(b, pro.String()...)
   }
   return string(b)
}

type Config struct {
   Video struct {
      ID int64
      Title string
   }
   SEO struct {
      Upload_Date string
      Thumbnail string
   }
   Request struct {
      Files struct {
         Progressive []Progressive
      }
   }
}

func (e Embed) Config() (*Config, error) {
   res, err := Client.Get(e.Config_URL)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Config)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

type Embed struct {
   Config_URL string
}

func New_Embed(ref string) (*Embed, error) {
   res, err := Client.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte(".OTTData =")
   scan.Scan()
   scan.Sep = nil
   emb := new(Embed)
   if err := scan.Decode(emb); err != nil {
      return nil, err
   }
   return emb, nil
}
