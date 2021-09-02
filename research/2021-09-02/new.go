package youtube

type param struct {
   SortBy uint32 `plenc:"1"`
   Filter struct {
      UploadDate uint32 `plenc:"1"`
      Type uint32 `plenc:"2"`
      Duration uint32 `plenc:"3"`
      HD uint32 `plenc:"4"`
      Subtitles uint32 `plenc:"5"`
      CreativeCommons uint32 `plenc:"6"`
      ThreeD uint32 `plenc:"7"`
      Live uint32 `plenc:"8"`
      Purchased uint32 `plenc:"9"`
      FourK uint32 `plenc:"14"`
      ThreeSixty uint32 `plenc:"15"`
      Location uint32 `plenc:"23"`
      HDR uint32 `plenc:"25"`
      VR180 uint32 `plenc:"26"`
   } `plenc:"2"`
}

func (p *param) underFourMinutes() {
   p.Filter.Duration = 1
}

func (p *param) overTwentyMinutes() {
   p.Filter.Duration = 2
}

func (p *param) fourToTwentyMinutes() {
   p.Filter.Duration = 3
}

func (p *param) video() {
   p.Filter.Type = 1
}

func (p *param) channel() {
   p.Filter.Type = 2
}

func (p *param) playlist() {
   p.Filter.Type = 3
}

func (p *param) movie() {
   p.Filter.Type = 4
}

func (p *param) lastHour() {
   p.Filter.UploadDate = 1
}

func (p *param) today() {
   p.Filter.UploadDate = 2
}

func (p *param) thisWeek() {
   p.Filter.UploadDate = 3
}

func (p *param) thisMonth() {
   p.Filter.UploadDate = 4
}

func (p *param) thisYear() {
   p.Filter.UploadDate = 5
}

func (p *param) relevance() {
   p.SortBy = 0
}

func (p *param) rating() {
   p.SortBy = 1
}

func (p *param) uploadDate() {
   p.SortBy = 2
}

func (p *param) viewCount() {
   p.SortBy = 3
}
