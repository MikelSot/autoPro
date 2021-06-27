package jwt

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"sync"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadFiles solo se ejecutara una vez
func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

// loadFiles carga los archivos que tienen los certificados
func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}

// parseRSA a rsa
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}

//func GenerateJWT(client model.Client) (string, error) {
//	myKey := []byte("autoPro")
//	payload := jwt.MapClaims{
//		"email": client.Email,
//		"name": client.Name,
//		"last_name":client.LastName,
//		"id":client.ID,
//		"expiration":time.Now().Add(time.Hour*24).Unix(),
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
//	tokenStr, err := token.SignedString(myKey)
//	if err != nil {
//		return tokenStr, err
//	}
//	return tokenStr, nil
//}
