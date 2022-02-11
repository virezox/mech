package instagram

import (
   "fmt"
   "testing"
)

const address = "https://www.instagram.com/p/CT-cnxGhvvO/"

func TestBufio(t *testing.T) {
   shortcode := bufioSplit(address)
   fmt.Printf("%q\n", shortcode)
}

func TestScanner(t *testing.T) {
   shortcode := textScanner(address)
   fmt.Printf("%q\n", shortcode)
}
