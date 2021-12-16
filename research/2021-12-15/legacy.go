package api

import (
   "golang.org/x/crypto/blowfish"
)

const (
	blowfishBlockSize = 8

	legacyAPIEndpoint = "https://tuner.pandora.com/services/json/"

	legacyPartnerUsername        = "android"
	legacyPartnerPassword        = "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7"
	legacyPartnerDeviceID        = "android-generic"
	legacyPartnerEncryptPassword = `6#26FRL$ZWD`
	legacyPartnerDecryptPassword = `R=U!LH$O2B#`
	legacyPartnerAPIVersion      = "5"
)

func mustCipher(key string) *blowfish.Cipher {
	c, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	return c
}

var (
	legacyEncryptCipher = mustCipher(legacyPartnerEncryptPassword)
	legacyDecryptCipher = mustCipher(legacyPartnerDecryptPassword)
)

type LegacyRequest struct {
	SyncTime int64 `json:"syncTime,omitempty"`
}

type LegacyResponse struct {
	Stat string `json:"stat"`
}

type LegacyPartnerLoginRequest struct {
	LegacyRequest

	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`

	IncludeURLs                bool `json:"includeUrls"`
	ReturnDeviceType           bool `json:"returnDeviceType"`
	ReturnUpdatePromptVersions bool `json:"returnUpdatePromptVersions"`
}

type LegacyPartnerLoginResponseResult struct {
	EncryptedSyncTime string `json:"syncTime"`
	PartnerID         string `json:"partnerId"`
	PartnerAuthToken  string `json:"partnerAuthToken"`
}

type LegacyPartnerLoginResponse struct {
	LegacyResponse
	Result LegacyPartnerLoginResponseResult `json:"result"`
}

type LegacyUserLoginRequest struct {
	LegacyRequest

	LoginType        string `json:"loginType"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	PartnerAuthToken string `json:"partnerAuthToken"`
}

type LegacyUserLoginResponseResult struct {
	UserAuthToken string `json:"userAuthToken"`
}

type LegacyUserLoginResponse struct {
	LegacyResponse
	Result LegacyUserLoginResponseResult `json:"result"`
}

func blowfishPad(b []byte) []byte {
	return append(b, make([]byte, blowfishBlockSize-(len(b)%blowfishBlockSize))...)
}

func legacyEncrypt(dst []byte, src []byte) {
	for bs, be := 0, blowfishBlockSize; bs < len(src); bs, be = bs+blowfishBlockSize, be+blowfishBlockSize {
		legacyEncryptCipher.Encrypt(dst[bs:be], src[bs:be])
	}
}

func legacyDecrypt(dst []byte, src []byte) {
	for bs, be := 0, blowfishBlockSize; bs < len(src); bs, be = bs+blowfishBlockSize, be+blowfishBlockSize {
		legacyDecryptCipher.Decrypt(dst[bs:be], src[bs:be])
	}
}
