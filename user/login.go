package user

import(
	"github.com/dgrijalva/jwt-go"
	"../db"
	_ "github.com/lib/pq"
	"errors"
	"time"
	"net/http"
	"fmt"
	"math"
)

type AuthCreds struct{
	User string
	Password string
}

type UserVerifyStatus int

const (
	TokenOK UserVerifyStatus = iota + 1
	NeedToRefresh
	TokenTooOld 
	TokenMalformed
	NoSuchUser
	InternalError
	UserLoggedOut
	TokenSignErr
)
var PrivateKey []byte

var maxTimeDelay, TokenLifeTime time.Duration

func init() {
	PrivateKey = []byte("hehe")
	maxTimeDelay = time.Hour * 40
	TokenLifeTime = time.Minute * 95
}

func HandleToken(statusToken UserVerifyStatus, username string, w http.ResponseWriter) bool{
	switch statusToken{
	case NoSuchUser:	
		fallthrough
	case TokenMalformed:
		http.Error(w, "", http.StatusBadRequest)
		return false
	case InternalError:
		http.Error(w, "", http.StatusInternalServerError)
		return false
	case UserLoggedOut:
		fallthrough
	case TokenSignErr:
		fallthrough
	case TokenTooOld:
		http.Error(w, "", http.StatusUnauthorized)
		return false
	case NeedToRefresh:
		accessTocken, err := GenerateJwtTocken(time.Now().Add(TokenLifeTime), username)
		if err != nil{
			http.Error(w, "", http.StatusInternalServerError)
			return false
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   accessTocken,
			HttpOnly : true,
		})
		
	}
	return true
}

func VerifyToken(w http.ResponseWriter, r *http.Request) (UserVerifyStatus, string){
	tokenString, err := r.Cookie("token")
	if err != nil{
		return TokenMalformed, ""
	}
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		return PrivateKey, nil
	})
	if err != nil && err.(*jwt.ValidationError).Errors != jwt.ValidationErrorExpired || token == nil{
		return TokenSignErr, ""
	}
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]
	expTimeRaw := claims["exp"].(float64)
	sec, dec := math.Modf(expTimeRaw)
	expTime := time.Unix(int64(sec), int64(dec*(1e9)))
	timeNow := time.Now()
	if expTime.Before(timeNow) && expTime.Before(timeNow.Add(maxTimeDelay)){
		return NeedToRefresh, username.(string)
	}else if expTime.Before(timeNow){
		return TokenTooOld, ""
	}
	//Is it bad? yes, but I dont want to make one more database request per HTTP request
	rows := db.DbConn.QueryRow("SELECT signed_in FROM public.users WHERE username = $1", username)
	var signedIn bool
	if err := rows.Scan(&signedIn); err != nil {
		return NoSuchUser, ""
	}
	if !signedIn{
		return UserLoggedOut, username.(string)
	}
	return TokenOK, username.(string)
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
	fmt.Println(token)
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

func RegisterUser(username, password string) bool{
	_, err:= db.DbConn.Query("INSERT INTO public.users (username, passwd) VALUES($1, $2)", username, password)
	if err!= nil{
		return false
	}
	return true
}