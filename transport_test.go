package mech

func main() {
   for _, gz := range []bool{false, true} {
      var c client
      c.DisableCompression = gz
      res, err := c.get("https://github.com/manifest.json")
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      fmt.Println(res.ContentLength)
      os.Stdout.ReadFrom(res.Body)
      fmt.Println()
   }
}
