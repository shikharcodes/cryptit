// encrypt package consists of all the encryption algorithms
package encrypt

// encrypts by increasing the ascii code for each character by 3
func EncryptStr(str string) string {
	encryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedString += character
	}
	return encryptedString
}
