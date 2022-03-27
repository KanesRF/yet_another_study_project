package user

import(
	"github.com/dgrijalva/jwt-go"
	"../db"
	_ "github.com/lib/pq"
	"errors"
	"time"
)

type AuthCreds struct{
	User string
	Password string
}

var PrivateKey []byte

func init() {
	PrivateKey = []byte("hehe")
}

func GenerateJwtTocken(exp time.Time, username string) (string, error){
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = exp.Unix()
	tmp_tocken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tmp_tocken.SignedString(PrivateKey)
	if err != nil {
	   return "", errors.New("Internal error")
	}
	return token, nil
}

func AuthByPassword(password, username string) bool{
	rows := db.DbConn.QueryRow("SELECT passwd FROM public.users WHERE username = $1", username)
	var db_passwd *string
	if err := rows.Scan(&db_passwd); err != nil {
		return false
	}
	if password != *db_passwd{
		return false
	}
	return true
}


/*
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
*/