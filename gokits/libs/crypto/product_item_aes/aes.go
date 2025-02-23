package product_item_aes

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/forgoer/openssl"
)

const (
	ERR_WRONG_KEY = "wrong key"

	blockSize = aes.BlockSize
)

func Encrypt(data, pass string, mainKey string) (string, error) {
	hexEnc, err := AesEncrypt(data, pass, mainKey)
	if err != nil {
		return "", err
	}

	return hexEnc + "-" + fmt.Sprintf("%x", md5.Sum([]byte(data)))[:5], nil
}

func Decrypt(data, pass string, mainKey string) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	hexData := strings.Split(data, "-")
	if len(hexData) == 0 {
		return "", errors.New("data invalid")
	}

	clearText, err := AesDecryptHex(hexData[0], pass, mainKey)
	if err != nil {
		return "", err
	}
	// fmt.Println("Decrypted data1:", fmt.Sprintf("%x", md5.Sum([]byte(clearText))))

	if len(hexData) == 2 && fmt.Sprintf("%x", md5.Sum([]byte(clearText)))[:5] != hexData[1] {
		return "", errors.New(ERR_WRONG_KEY)
	}
	return clearText, nil
}

func AesEncrypt(data, pass string, mainKey string) (string, error) {

	// format lại key theo blocksize
	key := []byte(mainKey + pass)
	if len(key) < blockSize {
		key = generateKey(key)
	} else {
		key = key[:blockSize]
	}

	cipherText, err := openssl.AesECBEncrypt([]byte(data), key, openssl.ZEROS_PADDING)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(cipherText), nil
}

func AesDecryptHex(hexData, pass string, mainKey string) (string, error) {

	// format lại key theo blocksize
	key := []byte(mainKey + pass)
	if len(key) < blockSize {
		key = generateKey(key)
	} else {
		key = key[:blockSize]
	}

	// decode hex string
	cipherText, err := hex.DecodeString(hexData)
	if err != nil {
		return "", err
	}
	if len(cipherText)%2 != 0 {
		return "", errors.New("invalid hex string")
	}

	// DECRYPT
	decryptedData, err := openssl.AesECBDecrypt(cipherText, key, openssl.ZEROS_PADDING)
	if err != nil {
		return "", err
	}

	// cắt kết quả thừa
	decryptedData = trimPadValue(decryptedData)

	return string(decryptedData), nil
}

func AesDecryptBase64(base64Data, pass string, mainKey string) (string, error) {
	// format lại key theo blocksize
	key := []byte(mainKey + pass)
	if len(key) < blockSize {
		key = generateKey(key)
	} else {
		key = key[:blockSize]
	}

	// decode base64 string
	cipherText, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// DECRYPT
	decryptedData, err := openssl.AesECBDecrypt(cipherText, key, openssl.ZEROS_PADDING)
	if err != nil {
		return "", err
	}

	// cắt kết quả thừa
	decryptedData = trimPadValue(decryptedData)

	return string(decryptedData), nil
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func trimPadValue(decryptedData []byte) []byte {
	padValue := decryptedData[len(decryptedData)-1]
	if padValue > 0 && padValue <= blockSize {
		decryptedData = decryptedData[:len(decryptedData)-int(padValue)]
	}

	return decryptedData
}
