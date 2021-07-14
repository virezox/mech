package main

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

const origin = "https://www.youtube.com"

var clients = []youtube.Client{                    
   {"ANDROID", "16.07.34"},                        
   {"ANDROID_CREATOR", "21.06.103"},               
   {"ANDROID_EMBEDDED_PLAYER", "16.20"},           
   {"ANDROID_KIDS", "6.02.3"},                     
   {"ANDROID_MUSIC", "4.32"},                      
   {"IOS", "16.05.7"},                             
   {"IOS_CREATOR", "20.47.100"},                   
   {"IOS_KIDS", "5.42.2"},                         
   {"IOS_MUSIC", "4.16.1"},                        
   {"MWEB", "2.19700101"},                         
   {"TVHTML5", "7.20210224.00.00"},                
   {"WEB", "2.20210223.09.00"},                    
   {"WEB_CREATOR", "1.20210223.01.00"},            
   {"WEB_EMBEDDED_PLAYER", "1.20210620.0.1"},      
   {"WEB_KIDS", "2.1.3"},                          
   {"WEB_REMIX", "0.1"},                           
}

func main() {
   req, err := http.NewRequest("GET", origin + "/get_video_info", nil)
   if err != nil {
      panic(err)
   }
   q := req.URL.Query()
   q.Set("c", "ANDROID")
   q.Set("cver", "16.05")
   q.Set("eurl", origin)
   q.Set("html5", "1")
   // pass
   //q.Set("video_id", "yFWKXXeOH_s")
   // fail
   q.Set("video_id", "Cr381pDsSsA")
   //fail
   //q.Set("el", "detailpage")
   //q.Set("el", "embedded")
   req.URL.RawQuery = q.Encode()
   fmt.Println(req.Method, req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   req.URL.RawQuery = string(body)
   player := req.URL.Query().Get("player_response")
   var play struct {
      PlayabilityStatus struct {
         Status string
      }
   }
   json.Unmarshal([]byte(player), &play)
   if play.PlayabilityStatus.Status == "OK" {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
