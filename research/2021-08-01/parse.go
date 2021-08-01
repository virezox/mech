package parse
import "time"

var dates = []string{
   "2020",
   "2020-12",
   "2020-12-31",
}

var times = []string{
   "23:59:59",
   "9:59:59",
   "59:59",
   "9:59",
}

//  403.9 ns/op
func datePad(l string) (time.Time, error) {
   r := "1970-01-01"[len(l):]
   return time.Parse("2006-01-02", l + r)
}

// 532.7 ns/op
func dateErr(s string) (time.Time, error) {
   t, err := time.Parse("2006-01-02", s)
   if err == nil {
      return t, nil
   }
   t, err = time.Parse("2006-01", s)
   if err == nil {
      return t, nil
   }
   return time.Parse("2006", s)
}

//  332.6 ns/op
func timeCrop(s string) (time.Time, error) {
   f := "15:04:05"[8-len(s):]
   return time.Parse(f, s)
}

//  510.3 ns/op
func timePad(r string) (time.Time, error) {
   l := "00:00:00"[:8-len(r)]
   return time.Parse("15:04:05", l + r)
}

//  694.0 ns/op
func timeErr(s string) (time.Time, error) {
   t, err := time.Parse("4:05", s)
   if err == nil {
      return t, nil
   }
   return time.Parse("15:04:05", s)
}
