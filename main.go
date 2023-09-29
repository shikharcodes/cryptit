package main

import (
	"fmt"

	"github.com/shikharcodes/cryptit/decrypt"
	"github.com/shikharcodes/cryptit/encrypt"
)

func main() {
	encryptedString := encrypt.EncryptStr("Practice")
	fmt.Println("Encrypted string:", encryptedString)
	fmt.Println("Decrypted string:", decrypt.EncryptStr(encryptedString))
}
