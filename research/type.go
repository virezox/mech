package mech

import (
   "github.com/89z/mech/widevine"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/http"
)

var client = http.Default_Client

type Streamer interface {
   Address() string
   Client_ID() string
   Info() bool
   Name() string
   Private_Key() string
   widevine.Poster
}

type dash_filter func([]dash.Representation) ([]dash.Representation, int)
