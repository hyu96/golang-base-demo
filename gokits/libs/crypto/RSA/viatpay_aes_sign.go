package RSA

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func Sign(keyPrivate string, data string) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(keyPrivate)
	if err != nil {
		return "", err
	}

	privateKey, err := parsePrivateKeyFromDER(privateKeyBytes)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func Verify(key_public, data, sign string) bool {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(key_public)
	if err != nil {
		fmt.Println(err)
	}

	publicKey, err := parsePublicKeyFromDER(publicKeyBytes)
	if err != nil {
		fmt.Println(err)
	}

	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println(err)
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, signByte)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func parsePublicKeyFromDER(der []byte) (*rsa.PublicKey, error) {
	var (
		pub *rsa.PublicKey
		err error
	)
	parsedKey, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKIXPublicKey() failed: %s, der:%s", err, string(der))
	}

	var ok bool
	if pub, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("parsed key is not an RSA public key")
	}

	return pub, nil
}

func parsePrivateKeyFromDER(der []byte) (*rsa.PrivateKey, error) {
	var (
		priv *rsa.PrivateKey
		err  error
	)
	parsedKey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS8PrivateKey() failed: %s, der:%s", err, string(der))
	}

	var ok bool
	if priv, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, fmt.Errorf("parsed key is not an RSA private key")
	}

	return priv, nil
}
