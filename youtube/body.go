package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const goog_API = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

const android_version = "17.24.34"

func Mobile_Web() Request {
   var r Request
   r.body.Context.Client.Name = "MWEB"
   r.body.Context.Client.Version = "2.20220322.05.00"
   return r
}

func Android() Request {
   var r Request
   r.body.Context.Client.Name = "ANDROID"
   r.body.Context.Client.Version = android_version
   return r
}

func Android_Embed() Request {
   var r Request
   r.body.Context.Client.Name = "ANDROID_EMBEDDED_PLAYER"
   r.body.Context.Client.Version = android_version
   return r
}

func Android_Racy() Request {
   var r Request
   r.body.Context.Client.Name = "ANDROID"
   r.body.Context.Client.Version = android_version
   r.body.Racy_Check_OK = true
   return r
}

func Android_Content() Request {
   var r Request
   r.body.Content_Check_OK = true
   r.body.Context.Client.Name = "ANDROID"
   r.body.Context.Client.Version = android_version
   r.body.Racy_Check_OK = true
   return r
}

func (r Request) Search(query string) (*Search, error) {
   filter := New_Filter()
   filter.Type(Type["Video"])
   param := New_Params()
   param.Filter(filter)
   r.body.Params = param.Marshal()
   r.body.Query = query
   buf, err := json.Marshal(r.body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/youtubei/v1/search", bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Goog-API-Key", goog_API)
   res, err := HTTP_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   search := new(Search)
   if err := json.NewDecoder(res.Body).Decode(search); err != nil {
      return nil, err
   }
   return search, nil
}

type Request struct {
   Header *Header
   body struct {
      Context struct {
         Client struct {
            Name string `json:"clientName"`
            Version string `json:"clientVersion"`
         } `json:"client"`
      } `json:"context"`
      Content_Check_OK bool `json:"contentCheckOk,omitempty"`
      Params []byte `json:"params,omitempty"`
      Query string `json:"query,omitempty"`
      Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
      Video_ID string `json:"videoId,omitempty"`
   }
}

func (r Request) Player(id string) (*Player, error) {
   r.body.Video_ID = id
   buf, err := json.Marshal(r.body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/youtubei/v1/player", bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   if r.Header != nil {
      req.Header.Set("Authorization", "Bearer " + r.Header.Access_Token)
   } else {
      req.Header.Set("X-Goog-API-Key", goog_API)
   }
   res, err := HTTP_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}
