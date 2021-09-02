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

type continuation struct {
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

func newContinuation(videoID string) continuation {
   var con continuation
   con.A.VideoID = videoID
   con.B = 6
   con.C.A.VideoID = videoID
   return con
}

type param struct {
   SortBy *uint32 `plenc:"1"`
   Filter *struct {
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

// 1
func (p *param) relevance() {
   p.SortBy = pUint32(0)
}

// 1
func (p *param) rating() {
   p.SortBy = pUint32(1)
}

// 1
func (p *param) uploadDate() {
   p.SortBy = pUint32(2)
}

// 1
func (p *param) viewCount() {
   p.SortBy = pUint32(3)
}

// 2 1
func (p *param) lastHour() {
   p.Filter.UploadDate = pUint32(1)
}

// 2 1
func (p *param) today() {
   p.Filter.UploadDate = pUint32(2)
}

// 2 1
func (p *param) thisWeek() {
   p.Filter.UploadDate = pUint32(3)
}

// 2 1
func (p *param) thisMonth() {
   p.Filter.UploadDate = pUint32(4)
}

// 2 1
func (p *param) thisYear() {
   p.Filter.UploadDate = pUint32(5)
}

// 2 2
func (p *param) video() {
   p.Filter.Type = pUint32(1)
}

// 2 2
func (p *param) channel() {
   p.Filter.Type = pUint32(2)
}

// 2 2
func (p *param) playlist() {
   p.Filter.Type = pUint32(3)
}

// 2 2
func (p *param) movie() {
   p.Filter.Type = pUint32(4)
}

// 2 3
func (p *param) underFourMinutes() {
   p.Filter.Duration = pUint32(1)
}

// 2 3
func (p *param) overTwentyMinutes() {
   p.Filter.Duration = pUint32(2)
}

// 2 3
func (p *param) fourToTwentyMinutes() {
   p.Filter.Duration = pUint32(3)
}

// 2 4
func (p *param) hd() {
   p.Filter.HD = pUint32(1)
}

// 2 5
func (p *param) subtitles() {
   p.Filter.Subtitles = pUint32(1)
}

// 2 6
func (p *param) creativeCommons() {
   p.Filter.CreativeCommons = pUint32(1)
}

// 2 7
func (p *param) threeD() {
   p.Filter.ThreeD = pUint32(1)
}

// 2 8
func (p *param) live() {
   p.Filter.Live = pUint32(1)
}

// 2 9
func (p *param) purchased() {
   p.Filter.Purchased = pUint32(1)
}

// 2 14
func (p *param) fourK() {
   p.Filter.FourK = pUint32(1)
}

// 2 15
func (p *param) threeSixty() {
   p.Filter.ThreeSixty = pUint32(1)
}

// 2 23
func (p *param) location() {
   p.Filter.Location = pUint32(1)
}

// 2 25
func (p *param) hdr() {
   p.Filter.HDR = pUint32(1)
}

// 2 26
func (p *param) vr180() {
   p.Filter.VR180 = pUint32(1)
}
