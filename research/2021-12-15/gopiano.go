package gopiano

import (
   "encoding/hex"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "strings"
   "time"
)

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
