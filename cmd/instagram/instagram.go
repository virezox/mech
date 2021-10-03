package main                                                                   
                                                                               
import (                                                                       
   "flag"                                                                      
   "fmt"                                                                       
   "github.com/89z/mech/instagram"                                             
   "net/http"                                                                  
   "net/url"                                                                   
   "os"                                                                        
   "path"                                                                      
)                                                                              
                                                                               
func main() {                                                                  
   var info bool                                                               
   flag.BoolVar(&info, "i", false, "info only")                                
   flag.Parse()                                                                
   if len(os.Args) == 1 {                                                      
      fmt.Println("instagram [-i] [ID]")                                       
      flag.PrintDefaults()                                                     
      return                                                                   
   }                                                                           
   id := flag.Arg(0)                                                           
   err := instagram.ValidID(id)                                                
   if err != nil {                                                             
      panic(err)                                                               
   }                                                                           
   instagram.Verbose = true                                                    
   med, err := instagram.NewMedia(id)                                          
   if err != nil {                                                             
      panic(err)                                                               
   }                                                                           
   if info {                                                                   
      fmt.Printf("%+v\n", med.Shortcode_Media)                                 
      return                                                                   
   }                                                                           
   // images                                                                   
   for _, edge := range med.Edges() {                                          
      err := download(edge.Node.Display_URL)                                   
      if err != nil {                                                          
         panic(err)                                                            
      }                                                                        
   }                                                                           
   // video                                                                    
   if med.Shortcode_Media.Video_URL != "" {                                    
      err := download(med.Shortcode_Media.Video_URL)                           
      if err != nil {                                                          
         panic(err)                                                            
      }                                                                        
   }                                                                           
}                                                                              
                                                                               
func download(addr string) error {                                             
   fmt.Println("GET", addr)                                                    
   res, err := http.Get(addr)                                                  
   if err != nil {                                                             
      return err                                                               
   }                                                                           
   defer res.Body.Close()                                                      
   par, err := url.Parse(addr)                                                 
   if err != nil {                                                             
      return err                                                               
   }                                                                           
   file, err := os.Create(path.Base(par.Path))                                 
   if err != nil {                                                             
      return err                                                               
   }                                                                           
   defer file.Close()                                                          
   if _, err := file.ReadFrom(res.Body); err != nil {                          
      return err                                                               
   }                                                                           
   return nil                                                                  
}                                                                              
