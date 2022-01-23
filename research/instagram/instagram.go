package instagram

type old struct {
   Display_URL string
   Edge_Media_Preview_Like struct { // Likes
      Count int64
   }
   Edge_Media_To_Parent_Comment struct { // Comments
      Edges []struct {
         Node struct {
            Text string
         }
      }
   }
   Edge_Sidecar_To_Children *struct { // Sidecar
      Edges []struct {
         Node struct {
            Display_URL string
            Video_URL string
         }
      }
   }
   Video_URL string
}
