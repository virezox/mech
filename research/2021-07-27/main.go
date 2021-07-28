package main
import "github.com/89z/mech/youtube"

var ids = []string{
   "1UztCDH2xuQ",
   "6iKPkxfljBY",
   "F1YdyaJeb1E",
   "GlhV-OKHecI",
   "MYr5MypHAhQ",
   "R7XcAaVumgc",
   "VKvn_YxuJQc",
   "WA8oNVFPppw",
   "Wk_AOIwGeOs",
   "XbUOX4lr9Bw",
   "eud9OOVM4to",
   "mjnAE5go9dI",
   "qMQJF-7Y2h0",
   "qmlJveN9IkI",
   "svTiG5vZ0_A",
   "uKna8o35UsU",
   "uhcnxH9zTEo",
   "unN7QvSWSTo",
   "w5azY0dH67U",
   "yGsCzZuK9GI",
}

const mb =
   "https://ia600709.us.archive.org/34/items" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4" +
   "/mbid-10cc746f-786c-4307-b8de-92a687489cb4-4958564206.jpg"

var mqDef = youtube.Image{320, 180, 180, "mqdefault", youtube.JPG}

func main() {
   err := corona10_main(mqDef)
   if err != nil {
      panic(err)
   }
}
