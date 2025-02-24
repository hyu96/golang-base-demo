package aes

import (
	"crypto/aes"
	"encoding/hex"
	"strings"

	"github.com/huydq/gokits/libs/ilog"
)

func AESEncryptWithHex(srcStr, keyStr string) string {
	encoded := AESEncrypt([]byte(srcStr), []byte(keyStr))
	if encoded == nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(encoded))
}

func AESDecryptWithHex(hexEncrypted, keyStr string) string {
	encoded, err := hex.DecodeString(hexEncrypted)
	if err != nil {
		ilog.Error("AESDecryptWithHex DecodeString err ", err)
		return ""
	}
	return string(AESDecrypt(encoded, []byte(keyStr)))
}

func AESEncrypt(src []byte, key []byte) (encrypted []byte) {
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		ilog.Error("AESEncrypt NewCipher err ", err)
		return nil
	}
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
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
