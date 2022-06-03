package apple

import (
   "bytes"
   "encoding/base64"
   "fmt"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

func TestAuth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   sign, err := OpenSignin(home, "mech/apple.json")
   if err != nil {
      t.Fatal(err)
   }
   res, err := sign.Auth()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
}

func TestSignin(t *testing.T) {
   con, err := NewConfig()
   if err != nil {
      t.Fatal(err)
   }
   sign, err := con.Signin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := sign.Create(home, "mech/apple.json"); err != nil {
      t.Fatal(err)
   }
}

const license = "CAISqgIKNAoQTVK8tm/4wLj5nPLnqI8QyBIQTVK8tm/4wLj5nPLnqI8QyBoAIAEoADgAQABIr4fmlAYSGAgBEAAYASCEBzCEBziA54QPSIwGUDxgARpmEhB38VPE8Q4ficxx5lFvXLLKGlBPmUkeH/2y0a9tIHrlIqYYxS6GgYtPGymDfUSuUa7BYiyxWlC4YMOsA45Ao7K6JAZaJBo/epmzl88WfZsjSs76nBy3aPKRtj52f0EiNY+KrSABGmIKEAAAAAAWgu8rYzAgICAgICASEAvvRO04+/Iq3xFBFTxMUk0aIIf8zsDKeG7fItLwp+D3QSJDc5qSEr0GohWAsMdNc7oIIAIoAkISChBrYzE2AAADhG3gkxCEAAAIYgJTRCCvh+aUBjjzxombBlACGiD9zGasCQwIBmr3KQkk7hIT3uOxK21s/tBVbAlAPyjCuyKAAWQTOA4NuPwFGA71AtokxQ+oLA82dDIsVYp/mka7zznc7p0ZsnNUu/t7BXssX985oJRWtG4xWymluBczNit65eQm9HNnWsgQHI70pSVI+2t7RwVX0dhkosL1yFjOlZLCPfzHtanCnXBeW1NEOYNm49m8M14pXUVC9L7qD3lcU+MIOjMKMTE2LjUuMCBCdWlsdCBvbiBEZWMgMTUgMjAyMSAxMDoyNzoxNyAoMTYzOTU5MjgxNSlAAUqwAQAAAAIAAACwAAUAEG3gkxBYbrfuAAAAVAAAABAAAABmAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAAAAAADhAAAAAAAAAAAAAAAAAHhNowAAAABAAAAvAAAABAAAADOAAAAEAAAAOAAAAAQAAAAAAAAAAAAAAEIAAAAEHAWPAtdMsPndcDmY3v1tDc8Ogh0YSLs6qsbKdFvM34AWAE="

func TestKeys(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   privateKey, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   kID, err := base64.StdEncoding.DecodeString("AAAAABaC7ytjMCAgICAgIA==")
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
   if err != nil {
      t.Fatal(err)
   }
   body, err := base64.StdEncoding.DecodeString(license)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Keys(bytes.NewReader(body))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", keys)
}
