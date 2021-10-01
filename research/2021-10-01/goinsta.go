package goinsta

import (
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "strconv"
   "time"
)

func defaultHandler(args ...interface{}) {
   fmt.Println(args...)
}

// Instagram represent the main API handler
//
// We recommend to use Export and Import functions after first Login.
//
// Also you can use SetProxy and UnsetProxy to set and unset proxy.
// Golang also provides the option to set a proxy using HTTP_PROXY env var.
type Instagram struct {
	user string
	pass string

	// device id: android-1923fjnma8123
	dID string
	// family device id, v4 uuid: 8b13e7b3-28f7-4e05-9474-358c6602e3f8
	fID string
	// uuid: 8493-1233-4312312-5123
	uuid string
	// rankToken
	rankToken string
	// token -- I think this is depricated, as I don't see any csrf tokens being used anymore, but not 100% sure
	token string
	// phone id v4 uuid: fbf767a4-260a-490d-bcbb-ee7c9ed7c576
	pid string
	// ads id: 5b23a92b-3228-4cff-b6ab-3199f531f05b
	adid string
	// pigeonSessionId
	psID string
	// contains header options set by Instagram
	headerOptions map[string]string
	// expiry of X-Mid cookie
	xmidExpiry int64
	// Public Key
	pubKey string
	// Public Key ID
	pubKeyID int
	// Device Settings
	device Device
	// User-Agent
	userAgent string
	// Account stores all personal data of the user and his/her options.
	Account *Account
	c *http.Client
	// Set to true to debug reponses
	Debug bool
	// Non-error message handlers.
	// By default they will be printed out, alternatively you can e.g. pass them to a logger
	infoHandler  func(...interface{})
	warnHandler  func(...interface{})
	debugHandler func(...interface{})
}

// New creates Instagram structure
func New(username, password string) *Instagram {
	insta := &Instagram{
		user: username,
		pass: password,
		dID: generateDeviceID(
			generateMD5Hash(username + password),
		),
		uuid:          generateUUID(),
		pid:           generateUUID(),
		fID:           generateUUID(),
		psID:          "UFS-" + generateUUID() + "-0",
		headerOptions: map[string]string{},
		xmidExpiry:    -1,
		device:        GalaxyS10,
		userAgent:     createUserAgent(GalaxyS10),
		c: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
		infoHandler:  defaultHandler,
		warnHandler:  defaultHandler,
		debugHandler: defaultHandler,
	}
         
	for k, v := range defaultHeaderOptions {
		insta.headerOptions[k] = v
	}

	return insta
}

// Export exports selected *Instagram object options to an io.Writer
func (insta *Instagram) ExportIO(writer io.Writer) error {
	config := ConfigFile{
		ID:            insta.Account.ID,
		User:          insta.user,
		DeviceID:      insta.dID,
		FamilyID:      insta.fID,
		UUID:          insta.uuid,
		RankToken:     insta.rankToken,
		Token:         insta.token,
		PhoneID:       insta.pid,
		XmidExpiry:    insta.xmidExpiry,
		HeaderOptions: map[string]string{},
		Account:       insta.Account,
		Device:        insta.device,
	}
         for key, value := range insta.headerOptions {
            config.HeaderOptions[key] = value
         }
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}

// Login performs instagram login sequence in close resemblance to the android
// apk. Password will be deleted after login.
func (insta *Instagram) Login() (err error) {
   err = insta.sync()
   if err != nil {
      return
   }
   if insta.pubKey == "" || insta.pubKeyID == 0 {
      return errors.New("Sync returned empty public key and/or public key id")
   }
   return insta.login()
}

func (insta *Instagram) login() error {
   timestamp := strconv.Itoa(int(time.Now().Unix()))
   if insta.pubKey == "" || insta.pubKeyID == 0 {
      return errors.New(
         "No public key or public key ID set. Please call Instagram.Sync() " +
         "and verify that it works correctly",
      )
   }
   encrypted, err := EncryptPassword(insta.pass, insta.pubKey, insta.pubKeyID, timestamp)
   if err != nil {
      return err
   }
   result, err := json.Marshal(
      map[string]interface{}{
         "adid":                insta.adid,
         "country_code":        "[{\"country_code\":\"44\",\"source\":[\"default\"]}]",
         "device_id":           insta.dID,
         "enc_password":        encrypted,
         "google_tokens":       "[]",
         "guid":                insta.uuid,
         "jazoest":             jazoest(insta.dID),
         "login_attempt_count": 0,
         "phone_id":            insta.fID,
         "username":            insta.user,
      },
   )
   if err != nil {
      return err
   }
   body, _, err := insta.sendRequest(
      &reqOptions{
         Endpoint: urlLogin,
         IsPost:   true,
         Query:    map[string]string{"signed_body": "SIGNATURE." + string(result)},
      },
   )
   if err != nil {
      return err
   }
   return insta.verifyLogin(body)
}

func (insta *Instagram) sync(args ...map[string]string) error {
   query := map[string]string{
      "id":                      insta.uuid,
      "server_config_retrieval": "1",
   }
   data, err := json.Marshal(query)
   if err != nil {
      return err
   }
   _, h, err := insta.sendRequest(
      &reqOptions{
         Endpoint: urlSync,
         Query:    generateSignature(data),
         IsPost:   true,
         IgnoreHeaders: []string{"Authorization"},
      },
   )
   if err != nil {
      return err
   }
   hkey := h["Ig-Set-Password-Encryption-Pub-Key"]
   hkeyID := h["Ig-Set-Password-Encryption-Key-Id"]
   var key string
   var keyID string
   if len(hkey) > 0 && len(hkeyID) > 0 && hkey[0] != "" && hkeyID[0] != "" {
      key = hkey[0]
      keyID = hkeyID[0]
   }
   id, err := strconv.Atoi(keyID)
   if err != nil {
      insta.warnHandler(fmt.Errorf("Failed to parse public key id: %s", err))
   }
   insta.pubKey = key
   insta.pubKeyID = id
   return nil
}

func (insta *Instagram) verifyLogin(body []byte) error {
   res := accountResp{}
   err := json.Unmarshal(body, &res)
   if err != nil {
      return fmt.Errorf("failed to parse json from login response %q", err)
   }
   if res.Status != "ok" {
      switch res.ErrorType {
      case "bad_password":
         return ErrBadPassword
      }
      return fmt.Errorf("Failed to login: %v, %v", res.ErrorType, res.Message)
   }
   insta.Account = &res.Account
   insta.Account.insta = insta
   insta.rankToken = strconv.FormatInt(insta.Account.ID, 10) + "_" + insta.uuid
   return nil
}