package product_item_aes

import (
	"fmt"
	"testing"
)

const KEY = "Yan.(-123"

func TestEncrypt(t *testing.T) {
	expected := "a0abadc2d95e2e0efb3604134537a04e7b5d24458178cc34dc215fe0a78e6810-41f7b"
	data := "Z14C47H/iTJBZmR9QxUHtw==" // Gifpop
	pass := "22442"
	encrypted, _ := Encrypt(data, pass, KEY)
	if encrypted != expected {
		t.Errorf("AESEncrypt mismatch(%x:%x)", encrypted, expected)
	}
}

func TestDecrypt(t *testing.T) {
	expected := "Z14C47H/iTJBZmR9QxUHtw==" // Gifpop
	encryptedData := "a0abadc2d95e2e0efb3604134537a04e7b5d24458178cc34dc215fe0a78e6810-41f7b"
	pass := "22442"
	decrypted, _ := Decrypt(encryptedData, pass, KEY)
	if string(decrypted) != expected {
		t.Errorf("AESDecrypt mismatch(%x:%x)", string(decrypted), expected)
	}
}

func TestAesDecrypt(t *testing.T) {
	expected := "999131263930" // Gifpop
	encryptedData := "Z14C47H/iTJBZmR9QxUHtw=="
	pass := "SEAKOVMV1RABDS8B"

	// DECRYPT
	decryptedData, err := AesDecryptBase64(encryptedData, pass, "")
	if err != nil {
		return
	}
	fmt.Println("decryptedData1: ", string(decryptedData))
	if string(decryptedData) != expected {
		t.Errorf("AESDecrypt mismatch(%x:%x)", string(decryptedData), expected)
	}
}
