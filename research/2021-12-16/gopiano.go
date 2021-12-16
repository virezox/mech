package main

import (
   "encoding/hex"
   "fmt"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "strings"
   "time"
)

const data = 
   "0b5b3f806abef32879a802a0749e65e9bea1623d9ff53d4c47e1db0a11135f61e8de2089919ef34facde3af32214db0ec6a249a41f1dee680de4a53dd649b4e447abe2f430167fbef7d0c40952bf86b" +
   "4f89edfd49c846b348affb88374876158e61e7a1897570b3c5d649871a9f7a843eaea7ad91b6568cce8a4825e79fe600a497a074213cba320157cb173b2f7160be19ebcd105667b7cd8e3c16f3cb658" +
   "8ec972859029ba17df0d643f758c918b8021c242b6fc63cb77698d9f20cd456757ae960d62daf16609f129dd0a75014a89649c7cb29abc3154a41b24843187b22c416e43635bc12a2aac171ac10dcb0" +
   "ae32a8c36352df5b63177dfa7ddc0fd577b94ab4855ef1746ccf72a766bd4bbb1a6599fc44629c0c24b90e7ae1a12a4a47f698ad33d515193ab4f3eed1001037610909df3e58804038634eeec596eed" +
   "c58f1d067b2d4a7b6e05136801425788d872e512be24936fa65e2a4ef9b796b76f3ea86242c72346da40f27a39516e3e887ca3ffbd0003be0e09735ef77a8ced7125b443bdb0ca3470fe098755d5d7d" +
   "559ccf552d87546587a1ecf9cb940b2a3417f2f080c3e3cd5dda618392fa73d82b896e12f222c09c70ed7e67df541a956a0c86f570dfed8d921a32f18e82f705579b232dc6ba4debf9d52d0bf280d5d" +
   "4b5b3f07a0b5c1965561912c2a452e255b799ba187c2c6eab13480a41275f6bddd777fddc2380d380bab700f77a58ee7b37328b66d7a7f3fbfd266"

func main() {
   client, err := NewClient(AndroidClient)
   if err != nil {
      panic(err)
   }
   dec, err := client.decrypt(data)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", dec)
}

// Describes a particular type of client to emulate.
type ClientDescription struct {
	DeviceModel string
	Username string
	Password string
	BaseURL string
	EncryptKey string
	DecryptKey string
	Version string
}

// The data for the Android client.
var AndroidClient ClientDescription = ClientDescription{
	DeviceModel: "android-generic",
	Username:    "android",
	Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "6#26FRL$ZWD",
	DecryptKey:  "R=U!LH$O2B#",
	Version:     "5",
}

// Class for a Client object.
type Client struct {
	description      ClientDescription
	http             *http.Client
	encrypter        *blowfish.Cipher
	decrypter        *blowfish.Cipher
	timeOffset       time.Duration
	partnerAuthToken string
	partnerID        string
	userAuthToken    string
	userID           string
}

// Create a new Client with specified ClientDescription
func NewClient(d ClientDescription) (*Client, error){
	client := new(http.Client)
	encrypter, err := blowfish.NewCipher([]byte(d.EncryptKey))
	if err != nil {
		return nil, err
	}
	decrypter, err := blowfish.NewCipher([]byte(d.DecryptKey))
	if err != nil {
		return nil, err
	}
	return &Client{
		description: d,
		http:        client,
		encrypter:   encrypter,
		decrypter:   decrypter,
	}, nil
}

// Blowfish encrypts a string in ECB mode.
// Many methods of the Pandora API take their JSON data as Blowfish encrypted data.
// The key for the encryption is provided by the ClientDescription.
func (c *Client) encrypt(data string) string {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 8 {
		var buf [8]byte
		var crypt [8]byte
		copy(buf[:], data[i:])
		c.encrypter.Encrypt(crypt[:], buf[:])
		encoded := hex.EncodeToString(crypt[:])
		chunks = append(chunks, encoded)
	}
	return strings.Join(chunks, "")
}

// Blowfish decrypts a string in ECB mode.
// Some data returned from the Pandora API is encrypted. This decrypts it.
// The key for the decryption is provided by the ClientDescription.
func (c *Client) decrypt(data string) (string, error) {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 16 {
		var buf [16]byte
		var decoded, decrypted [8]byte
		copy(buf[:], data[i:])
		_, err := hex.Decode(decoded[:], buf[:])
		if err != nil {
			return "", err
		}
		c.decrypter.Decrypt(decrypted[:], decoded[:])
		chunks = append(chunks, strings.Trim(string(decrypted[:]), "\x00"))
	}
	return strings.Join(chunks, ""), nil
}

// Most calls require a SyncTime int argument (Unix epoch). We store our current time offset
// but must calculate the SyncTime for each call. This method does that.
func (c *Client) GetSyncTime() int {
	return int(time.Now().Add(c.timeOffset).Unix())
}
