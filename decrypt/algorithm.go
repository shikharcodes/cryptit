// decrypt package consists of all the decryption algorithms
package decrypt

// decrypts by reducing the ascii code for each character by 3
func EncryptStr(str string) string {
	decryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedString += character
	}
	return decryptedString
}
