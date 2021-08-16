package decode

// Move to the next element with the given attribute. Set "Attr" to the element
// attributes.
func (d *Decoder) NextAttr(key, val string) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         d.Attr = make(map[text]text)
      }
      if t == html.AttributeToken {
         d.Attr[d.Text()] = d.AttrVal()
      }
      if t == html.StartTagCloseToken {
         if v, ok := d.Attr[key]; ok && v == val {
            return true
         }
      }
   }
   return false
}

// Move to the next element with the given tag. Set "Data" to the tag name, and
// set "Attr" to nil.
func (d *Decoder) NextTag(name Text) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         if d.Data = d.Text(); d.Data == name {
            return true
         }
      }
   }
   return false
}
