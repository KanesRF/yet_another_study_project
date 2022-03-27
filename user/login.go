package user

import(
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func GenerateSalt(saltSize int)[]byte{
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil{
		return nil
	}
	return salt
}

func GenerateTocken(password []byte)string{
	salt := GenerateSalt(16)
	if salt == nil{
		return ""
	}
	var sha512 = sha512.New()
	var passwordWithSalt = make([]byte, 64)
	passwordWithSalt = append(passwordWithSalt, password...)
	passwordWithSalt = append(passwordWithSalt, salt...)
	sha512.Write(passwordWithSalt)
   	hashPassword := sha512.Sum(nil)
	var encodedHash = base64.URLEncoding.EncodeToString(hashPassword)
	return encodedHash
}