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

var (
   relevance = param{SortBy: 0}
   rating = param{SortBy: 1}
   uploadDate = param{SortBy: 2}
   viewCount = param{SortBy: 3}
)
