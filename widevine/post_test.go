package widevine

import (
   "encoding/base64"
   "os"
   "testing"
)

const (
   raw_response = "CAIShgQKNwoQUjsj+XzX0+TUCd6LFEElRhIQYU43Z1dXUU1SNTNkdjJXaxoAIAEoADgAQICangFIztPolgYSGggBEAAYASCAmp4BKICangEwgJqeATiA54QPGmYSEAl3zTn1/jZyWvKUS0YxGO0aUNIuGzZLC5m15lO0N40vhkD9I1ISiHh2HDSVBlLNt8VVWFAZScZoAzfAqzzaQvWDpc6CyCseiNeIuqpIOkLpUgQE/9ik5UgoaymhdmXJCIvuIAEaaAoQBRRhar6OXE+Y2Rv0snUmOxIQUT4SzCGrMsIuPs/PlqwU3BogRDSGvLI+djIimUieIAAWGgZK9jv1q/mVCi5VPubTdRIgAigBOgQIABADQhIKEGtjMTYAJ40AMbjeY4AAAAhiAkhEGmgKEIf93v7Vu1OasRpHjXVv/GgSED6BX/bjLlMp9qxYWvB8vhMaIKFFDvOnXssaa+fIsCRur47hjlJ9OGFZ/W2F2S6+ZZHWIAIoAjoECAAQA0ISChBrYzE2ACeNADG43mOEAAAIYgJTRBprChBIgc8uESFXQK0299dyaSJTEhCE49TviAVwXb3J3q5BwK3FGiBRXuRJJKMIs9SnVqSyCg3vpPDWZYPwClYrt54BlrYe8yACKAE6BAgAEANCEgoQa2MxNgAnjQAxuN5jgAAACGIFQVVESU8gztPolgZQBRogH42FT1M/eb/wbQcQGmkfoRNPU4BvQ5uSk8Fzgbaw/0gigAGGMhO95ZeA73FOq4GKlV28DgSwD8GFh1/MW79/+B2abCNRJeyRyXkgTK8X2P9tGSPO/zXK4yH7aN9pxSoQ0Jn+t9REP0jCJgDAz0iwqGQchIdZVUvb8PT4qDVKLy7kS2DeIR0Ba041qd8trEwvIP9gaeDuzcMHeABnu+EutNRx1ToyCjAxNy4xLjAgQnVpbHQgb24gSnVuIDEgMjAyMiAyMTozMjozMSAoMTY1NDE0NDMzMylAAUqAAgAAAAIAAAEAAAUAEDG43mNcoYRLAAAAWQAAABAAAABrAAAAUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAAAACeNAAAAAAAAJ40AAAAAAAHhM4AAAAADAAAAwQAAABAAAADTAAAAEAAAAOUAAAAQAAAAAAAAAAAAAAETAAAAEAAAASsAAAAQAAABPQAAABAAAAFPAAAAEAAAAAAAAAAAAAABfQAAABAAAAGVAAAAEAAAAacAAAAQAAABuQAAABAAAAAAAAAAAAAAARMAAAAQYvEVsHbtWlsJw9LQaLRhe52V0B4d3O0tcXq2yQtc+5VYAQ=="
   raw_key_ID = "59545F4D454449413A33353531353839636234313138643164"
)

func Test_Response(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   key_ID, err := Key_ID(raw_key_ID)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := New_Module(private_key, client_ID, key_ID)
   if err != nil {
      t.Fatal(err)
   }
   response, err := base64.StdEncoding.DecodeString(raw_response)
   if err != nil {
      t.Fatal(err)
   }
   if _, err := mod.signed_response(response); err != nil {
      t.Fatal(err)
   }
}
