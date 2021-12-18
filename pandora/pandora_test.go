package pandora

import (
   "encoding/hex"
   "fmt"
   "testing"
)

const exchange = "0b5b3f806abef32879a802a0749e65e95532213d5d6f3e13bf1af53914ae78b0126c0066335814d5dd9855ad55afe1013a4397fbbf18824ad73ea9310d11ae0eded999de98861f981fda8c79ce45ce9db4831b8c3d3aebc7d42c60acc484f600a356f4ceaf5bb53af1b5238bd72e51715e2a094733a01a0234057f785c279b0732a13d073647dd8fb0c8427e19a8b331252e93e50298048fb9aabcdd014155feaf121737ea31a493a136d13e9b0096c1310b76d03f4dffa0c1197668a530fd7dd4194bfb777abd81589c69c9a59ba2b9d8909d37673375399b7a8811cea418dd97346a7aa7f4b71d085ca8b15dc64526a90cb83a985947feeeac6e210b13f341cd0eb6df0faa50c1d0065d35829a24ed3ab716ff46d861b307275f718116fc2432a70e78eec88b821a897f0d7491367a057c408be2220be330bc0ed32e4c6380142c2a3d009ce958e223cd7d6705334ac6f214714ee0b05600c29789f14b8cb900192ab8faebe40f0db6ecc2473e1d844dc38b2302358693d3d45adaa2bdc22296b9ee3443ac80dd728e60162c23063c4fe2acb18dced62052c9434b54aea0f37b5a2cbde2e4c22523f38a56f4ecdc07f12a8a741b8b3bdbee67579bbd98f710906fd6e385ffd5106e11f50a76f8ceb57af7be4ad0502f8e00460d613a2c5c1c957bf4350bc7dcc3775433b134599e747498babbd90d56460fbad0e585dbe4d04110fb6365945c2aa49605b64ffb8a049abf26ac6336b152c3d74d2189d9ea59ad9887a1daa62fb477b4cb3e96e2b3facc3b6106bc4ebfd9cde2f60c636b9cd66ef22dd64d2605ef30696df89c16a646dc7601d43bf9541007df4b618e2caeeed5ae216753f7a77d30696df89c16a646692ce2f3aa7147f8969a72ccd0c845f26d6214f3a077932a32023fdbfe809e51fd16858e9d50fd42ff5fff100763e45283dfc96db7d7b4b160ef8fde5f854556cacba9ed94bcee2cdc8962cf0b36263e56f7f07ae9cdead9d0587b83ede2f0d07b901c3b49e9a2ff60ef8fde5f854556cee824a38ffa662e1ff730ee532489329ef2638c1c296a8dc8a4c72283158fb492c7e9f601c05ff647ade59284801a9c798c6a883cb57f998b1d33a3a1926f6b70e0e06bb389027fc94fa31251eb858ffec867daaffff04e928848169e34546646eab4a411d8d9c76b7c990941627c1eaffaaaa7fbac8ea2e025e3efba4d887afed7f46751a5a21534f02b863fab8f6b0e5fcc3380d96b2d187a6da907c842a9c0a07486a61e7804948f45fc03e61db2f36521d8ccb8e775f274b19a1d1d2901370eba71167671e3a08420ba579e7b414a844fb63bb4fa2f937829db60f047a33b02ceb6030ad8a28e44201d95a3c6cc791344d7abf11589d7602a508744cb3ff049e8e7bb13923d1fa019222afd3ad709988d56c80b370d5362673acfb5e92cec06b9f8868a109fad4746bb2493042026c90d2232c3cc1f7e356b26e10fcda4"

func TestLogin(t *testing.T) {
   part, err := NewPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(part.Result.PartnerAuthToken)
   if tLen != 34 {
      t.Fatal(tLen)
   }
   user, err := part.UserLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   if tLen := len(user.Result.UserAuthToken); tLen != 58 {
      panic(tLen)
   }
   if err := user.ValueExchange(); err != nil {
      t.Fatal(err)
   }
   info, err := user.PlaybackInfo()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", info)
}

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(exchange)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := Decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", dec)
}
