package config

import (
	"io/ioutil"
	"os"
	"crypto/rsa"

	jwt "github.com/dgrijalva/jwt-go"
)

//Cfg store config
var (
	Cfg *Config
)

//Config ...
type Config struct {
	ConnectionString string
	VerifyKey *rsa.PublicKey
	SignKey *rsa.PrivateKey
}

func init() { 	
	Cfg = &Config{
		ConnectionString: os.Getenv("DB_CONNECTION"),
	}

	//Load rsa file to signed and verify
	signBytes, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_FILE"))
	if err != nil {
		panic(err)
	}

	Cfg.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_FILE"))
	if err != nil {
		panic(err)
	}

	Cfg.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}