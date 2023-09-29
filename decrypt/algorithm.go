package decrypt

func EncryptStr(str string) string {
	decryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedString += character
	}
	return decryptedString
}
