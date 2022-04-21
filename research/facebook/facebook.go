package facebook

import (
   "github.com/89z/format"
   "github.com/89z/format/xml"
   "net/http"
   "strconv"
)

type Date string

func newDate(v int64) (Date, error) {
   req, err := http.NewRequest("GET", "https://m.facebook.com/video.php", nil)
   if err != nil {
      return "", err
   }
   req.Header.Set("User-Agent", windows.String())
   req.URL.RawQuery = "v=" + strconv.FormatInt(v, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   scan, err := xml.NewScanner(res.Body)
   if err != nil {
      return "", err
   }
   scan.Split = []byte("<abbr>")
   scan.Scan()
   var date Date
   if err := scan.Decode(&date); err != nil {
      return "", err
   }
   return date, nil
}

var LogLevel format.LogLevel

type userAgent struct {
   browser string
   browserVersion int64
   engine string
   engineVersion int64
   system string
   systemFamily string
   systemVersion int64
}

var windows = userAgent{
   browser: "Firefox",
   browserVersion: 99,
   engine: "Gecko",
   engineVersion: 9,
   system: "Windows",
   systemFamily: "NT",
   systemVersion: 9,
}

func (u userAgent) String() string {
   var buf []byte
   buf = append(buf, u.system...)
   buf = append(buf, ' ')
   buf = append(buf, u.systemFamily...)
   buf = append(buf, ' ')
   buf = strconv.AppendInt(buf, u.systemVersion, 10)
   buf = append(buf, ' ')
   buf = append(buf, u.engine...)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.engineVersion, 10)
   buf = append(buf, ' ')
   buf = append(buf, u.browser...)
   buf = append(buf, '/')
   buf = strconv.AppendInt(buf, u.browserVersion, 10)
   return string(buf)
}
