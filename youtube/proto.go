package youtube

import (
   "encoding/base64"
   "github.com/philpearl/plenc"
)

func encode(value interface{}) (string, error) {
   b, err := plenc.Marshal(nil, value)
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(b), nil
}

func pUint32(v uint32) *uint32 {
   return &v
}

type Continuation struct {
   A struct {
      VideoID string `plenc:"2"`
   } `plenc:"2"`
   B uint32 `plenc:"3"`
   C struct {
      A struct {
         VideoID string `plenc:"4"`
      } `plenc:"4"`
   } `plenc:"6"`
}

func NewContinuation(videoID string) Continuation {
   var con Continuation
   con.A.VideoID = videoID
   con.B = 6
   con.C.A.VideoID = videoID
   return con
}

func (c Continuation) Encode() (string, error) {
   return encode(c)
}

type Param struct {
   SortBy *uint32 `plenc:"1"`
   Filter struct {
      UploadDate *uint32 `plenc:"1"`
      Type *uint32 `plenc:"2"`
      Duration *uint32 `plenc:"3"`
      HD *uint32 `plenc:"4"`
      Subtitles *uint32 `plenc:"5"`
      CreativeCommons *uint32 `plenc:"6"`
      ThreeD *uint32 `plenc:"7"`
      Live *uint32 `plenc:"8"`
      Purchased *uint32 `plenc:"9"`
      FourK *uint32 `plenc:"14"`
      ThreeSixty *uint32 `plenc:"15"`
      Location *uint32 `plenc:"23"`
      HDR *uint32 `plenc:"25"`
      VR180 *uint32 `plenc:"26"`
   } `plenc:"2"`
}

func (p Param) Encode() (string, error) {
   return encode(p)
}

// 1
func (p *Param) Relevance() {
   p.SortBy = pUint32(0)
}

// 1
func (p *Param) Rating() {
   p.SortBy = pUint32(1)
}

// 1
func (p *Param) UploadDate() {
   p.SortBy = pUint32(2)
}

// 1
func (p *Param) ViewCount() {
   p.SortBy = pUint32(3)
}

// 2 1
func (p *Param) LastHour() {
   p.Filter.UploadDate = pUint32(1)
}

// 2 1
func (p *Param) Today() {
   p.Filter.UploadDate = pUint32(2)
}

// 2 1
func (p *Param) ThisWeek() {
   p.Filter.UploadDate = pUint32(3)
}

// 2 1
func (p *Param) ThisMonth() {
   p.Filter.UploadDate = pUint32(4)
}

// 2 1
func (p *Param) ThisYear() {
   p.Filter.UploadDate = pUint32(5)
}

// 2 2
func (p *Param) Video() {
   p.Filter.Type = pUint32(1)
}

// 2 2
func (p *Param) Channel() {
   p.Filter.Type = pUint32(2)
}

// 2 2
func (p *Param) Playlist() {
   p.Filter.Type = pUint32(3)
}

// 2 2
func (p *Param) Movie() {
   p.Filter.Type = pUint32(4)
}

// 2 3
func (p *Param) UnderFourMinutes() {
   p.Filter.Duration = pUint32(1)
}

// 2 3
func (p *Param) OverTwentyMinutes() {
   p.Filter.Duration = pUint32(2)
}

// 2 3
func (p *Param) FourToTwentyMinutes() {
   p.Filter.Duration = pUint32(3)
}

// 2 4
func (p *Param) HD() {
   p.Filter.HD = pUint32(1)
}

// 2 5
func (p *Param) Subtitles() {
   p.Filter.Subtitles = pUint32(1)
}

// 2 6
func (p *Param) CreativeCommons() {
   p.Filter.CreativeCommons = pUint32(1)
}

// 2 7
func (p *Param) ThreeD() {
   p.Filter.ThreeD = pUint32(1)
}

// 2 8
func (p *Param) Live() {
   p.Filter.Live = pUint32(1)
}

// 2 9
func (p *Param) Purchased() {
   p.Filter.Purchased = pUint32(1)
}

// 2 14
func (p *Param) FourK() {
   p.Filter.FourK = pUint32(1)
}

// 2 15
func (p *Param) ThreeSixty() {
   p.Filter.ThreeSixty = pUint32(1)
}

// 2 23
func (p *Param) Location() {
   p.Filter.Location = pUint32(1)
}

// 2 25
func (p *Param) HDR() {
   p.Filter.HDR = pUint32(1)
}

// 2 26
func (p *Param) VR180() {
   p.Filter.VR180 = pUint32(1)
}
