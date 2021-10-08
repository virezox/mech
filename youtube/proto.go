package youtube

import (
   "encoding/base64"
   "github.com/philpearl/plenc"
)

func pUint(v uint) *uint {
   return &v
}

type Param struct {
   SortBy *uint `plenc:"1"`
   Filter struct {
      UploadDate *uint `plenc:"1"`
      Type *uint `plenc:"2"`
      Duration *uint `plenc:"3"`
      HD *uint `plenc:"4"`
      Subtitles *uint `plenc:"5"`
      CreativeCommons *uint `plenc:"6"`
      ThreeD *uint `plenc:"7"`
      Live *uint `plenc:"8"`
      Purchased *uint `plenc:"9"`
      FourK *uint `plenc:"14"`
      ThreeSixty *uint `plenc:"15"`
      Location *uint `plenc:"23"`
      HDR *uint `plenc:"25"`
      VR180 *uint `plenc:"26"`
   } `plenc:"2"`
}

func (p Param) Encode() (string, error) {
   b, err := plenc.Marshal(nil, p)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(b), nil
}

// CAASAA==
func (p *Param) Relevance() {
   p.SortBy = pUint(0)
}

// CAESAA==
func (p *Param) Rating() {
   p.SortBy = pUint(1)
}

// CAISAA==
func (p *Param) UploadDate() {
   p.SortBy = pUint(2)
}

// CAMSAA==
func (p *Param) ViewCount() {
   p.SortBy = pUint(3)
}

// EgIIAQ==
func (p *Param) LastHour() {
   p.Filter.UploadDate = pUint(1)
}

// EgIIAg==
func (p *Param) Today() {
   p.Filter.UploadDate = pUint(2)
}

// EgIIAw==
func (p *Param) ThisWeek() {
   p.Filter.UploadDate = pUint(3)
}

// EgIIBA==
func (p *Param) ThisMonth() {
   p.Filter.UploadDate = pUint(4)
}

// EgIIBQ==
func (p *Param) ThisYear() {
   p.Filter.UploadDate = pUint(5)
}

// EgIQAQ==
func (p *Param) Video() {
   p.Filter.Type = pUint(1)
}

// EgIQAg==
func (p *Param) Channel() {
   p.Filter.Type = pUint(2)
}

// EgIQAw==
func (p *Param) Playlist() {
   p.Filter.Type = pUint(3)
}

// EgIQBA==
func (p *Param) Movie() {
   p.Filter.Type = pUint(4)
}

// EgIYAQ==
func (p *Param) UnderFourMinutes() {
   p.Filter.Duration = pUint(1)
}

// EgIYAg==
func (p *Param) OverTwentyMinutes() {
   p.Filter.Duration = pUint(2)
}

// EgIYAw==
func (p *Param) FourToTwentyMinutes() {
   p.Filter.Duration = pUint(3)
}

// EgIgAQ==
func (p *Param) HD() {
   p.Filter.HD = pUint(1)
}

// EgIoAQ==
func (p *Param) Subtitles() {
   p.Filter.Subtitles = pUint(1)
}

// EgIwAQ==
func (p *Param) CreativeCommons() {
   p.Filter.CreativeCommons = pUint(1)
}

// EgI4AQ==
func (p *Param) ThreeD() {
   p.Filter.ThreeD = pUint(1)
}

// EgJAAQ==
func (p *Param) Live() {
   p.Filter.Live = pUint(1)
}

// EgJIAQ==
func (p *Param) Purchased() {
   p.Filter.Purchased = pUint(1)
}

// EgJwAQ==
func (p *Param) FourK() {
   p.Filter.FourK = pUint(1)
}

// EgJ4AQ==
func (p *Param) ThreeSixty() {
   p.Filter.ThreeSixty = pUint(1)
}

// EgO4AQE=
func (p *Param) Location() {
   p.Filter.Location = pUint(1)
}

// EgPIAQE=
func (p *Param) HDR() {
   p.Filter.HDR = pUint(1)
}

// EgPQAQE=
func (p *Param) VR180() {
   p.Filter.VR180 = pUint(1)
}
