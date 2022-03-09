package main

import (
   "os"
)

type osFile struct {
   *os.File
}
