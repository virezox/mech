package main

import (
   "crypto/rand"
   "crypto/rsa"
   "crypto/sha1"
   "encoding/base64"
   "fmt"
   "math/big"
   "os"
)

const androidKeyBase64 = "AAAAgMom/1a/v0lblO2Ubrt60J2gcuXSljGFQXgcyZWveWLEwo6prwgi3iJIZdodyhKZQrNWp5nKJ3srRXcUW+F1BD3baEVGcmEgqaLZUNBjm057pKRI16kB0YppeGx5qIQ5QjKzsR8ETQbKLNWgRY0QRNVz34kMJR3P/LgHax/6rmf5AAAAAwEAAQ=="

func bytesToLong(b []byte) *big.Int {
	return new(big.Int).SetBytes(b)
}

func signature(email, password string) (string, error) {
	androidKeyBytes, err := base64.StdEncoding.DecodeString(androidKeyBase64)
	if err != nil {
		return "", err
	}
	i := bytesToLong(androidKeyBytes[:4]).Int64()
	j := bytesToLong(androidKeyBytes[i+4 : i+8]).Int64()
	androidKey := &rsa.PublicKey{
		N: bytesToLong(androidKeyBytes[4 : 4+i]),
		E: int(bytesToLong(androidKeyBytes[i+8 : i+8+j]).Int64()),
	}
	hash := sha1.Sum(androidKeyBytes)
	msg := append([]byte(email), 0)
	msg = append(msg, []byte(password)...)
	encryptedLogin, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, androidKey, msg, nil)
	if err != nil {
		return "", err
	}
	sig := append([]byte{0}, hash[:4]...)
	sig = append(sig, encryptedLogin...)
	return base64.URLEncoding.EncodeToString(sig), nil
}

func main() {
   sig, err := signature("srpen6@gmail.com", os.Args[1])
   if err != nil {
      panic(err)
   }
   fmt.Println(sig)
}
