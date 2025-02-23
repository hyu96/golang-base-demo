package RSA

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	// Generate key pair
	rsaPrivate := readfile("keyRSA/private_key.pem")
	rsaPublic := readfile("keyRSA/public_key.pem")

	// Test signing and verifying
	data := "paytechvnpaytechvnPAY73fee5e0757b4ec9addeebf189f1859a20230309214037"
	signature, err := Sign(rsaPrivate, data)
	if err != nil {
		t.Fatalf("Failed to sign data: %s", err)
	}
	if !Verify(rsaPublic, data, signature) {
		t.Fatalf("Signature verification failed")
	}
}

func readfile(path string) string {
	var xau string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xau += scanner.Text() + "\n"
	}
	xau = strings.TrimSpace(xau)
	return xau
}
