package instagram

type one struct {
   GraphQL struct {
      Shortcode_Media struct {
         Edge_Media_To_Caption three
         Edge_Media_To_Parent_Comment three
         Display_URL string
         Video_URL string
         Edge_Sidecar_To_Children four
      }
   }
}

type two struct {
   GraphQL struct {
      User struct {
         Edge_Owner_To_Timeline_Media four
      }
   }
}

type three struct {
   Edges []struct {
      Node struct {
         Text string
      }
   }
}

type four struct {
   Edges []struct {
      Node struct {
         Display_URL string
         Video_URL string
      }
   }
}
