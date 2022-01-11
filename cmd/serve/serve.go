package main
 
import (
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
)
 
func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      id := r.URL.Query().Get("v")
      play, err := youtube.NewPlayer(id, youtube.Key, youtube.Android)
      if err != nil {
         fmt.Fprint(w, err)
      } else {
         for _, form := range play.StreamingData.AdaptiveFormats {
            if form.Itag == 251 {
               fmt.Fprint(w, form.URL)
            }
         }
      }
   })
   addr := ":8080"
   fmt.Println(addr)
   http.ListenAndServe(addr, nil)
}
