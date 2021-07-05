package youtube

const (
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

func (r Search) Videos() []CompactVideoRenderer {
   var vids []CompactVideoRenderer
   for _, sect := range r.Contents.SectionListRenderer.Contents {
      for _, item := range sect.ItemSectionRenderer.Contents {
         vids = append(vids, item.CompactVideoRenderer)
      }
   }
   return vids
}
