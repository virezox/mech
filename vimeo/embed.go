package vimeo

import (
   "github.com/89z/rosso/json"
   "io"
   "strconv"
   "time"
)

func (c Config) String() string {
   b := []byte("Date: ")
   b = append(b, c.SEO.Upload_Date...)
   b = append(b, "\nDuration: "...)
   b = append(b, c.Duration().String()...)
   b = append(b, "\nID: "...)
   b = strconv.AppendInt(b, c.Video.ID, 10)
   b = append(b, "\nTitle: "...)
   b = append(b, c.Video.Title...)
   for _, p := range c.Request.Files.Progressive {
      b = append(b, "\nWidth:"...)
      b = strconv.AppendInt(b, p.Width, 10)
      b = append(b, " Height:"...)
      b = strconv.AppendInt(b, p.Height, 10)
      b = append(b, " FPS:"...)
      b = strconv.AppendInt(b, p.FPS, 10)
   }
   return string(b)
}

func (c Config) Duration() time.Duration {
   return time.Duration(c.Video.Duration) * time.Second
}

type Config struct {
   SEO struct {
      Thumbnail string
      Upload_Date string
   }
   Video struct {
      Duration int64
      ID int64
      Title string
   }
   Request struct {
      Files struct {
         Progressive []struct {
            Width int64
            Height int64
            FPS int64
            URL string
         }
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
