package nbc

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
   "testing"
)

var videos = []int{
   // nbc.com/botched/video/seeing-double/3049418
   // 2304982139 3049418 200 OK
   // "resourceID": "e",
   3049418,
   // nbc.com/pasion-de-gavilanes/video/la-valentia-de-norma/9000221348
   // 2304991196 9000221348 200 OK
   // "resourceID": "telemundo",
   9000221348,
}

func TestNBC(t *testing.T) {
}
