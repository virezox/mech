package pbs

import (
   "fmt"
   "testing"
)

const widgetTest = "https://player.pbs.org/widget/partnerplayer/3016754074/"

func TestWidget(t *testing.T) {
   test, err := NewWidgeter(widgetTest)
   if err != nil {
      t.Fatal(err)
   }
   widget, err := test.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", widget)
}
