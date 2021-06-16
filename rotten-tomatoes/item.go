package tomato

type Item struct {
   Name string
   ReleaseYear string
   TomatoMeterScore struct {
      Score string
   }
   URL string
}

func NewItem(title string) {
}
