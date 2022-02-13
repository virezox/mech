package main

import (
   "flag"
   "github.com/89z/mech/tumblr"
)

func main() {
   var link tumblr.Permalink
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   flag.StringVar(&link.BlogName, "b", "", "blog name")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // p
   flag.Int64Var(&link.PostID, "p", 0, "post ID")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      tumblr.LogLevel = 1
   }
   if link.BlogName != "" && link.PostID >= 1 {
      err := blogPost(&link, info)
      if err != nil {
         panic(err)
      }
   } else if address != "" {
      link, err := tumblr.NewPermalink(address)
      if err != nil {
         panic(err)
      }
      if err := blogPost(link, info); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
