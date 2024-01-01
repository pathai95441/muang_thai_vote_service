package authorization

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pathai95441/muang_thai_vote_service/src/config"
)

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// func main() {
// 	secretKey := []byte("your_secret_key")

// 	token := createToken("john_doe", secretKey)

// 	// Validate and parse the token
// 	claims, err := validateToken(token, secretKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Access the custom claims
// 	fmt.Println("Username:", claims.Username)
// 	fmt.Println("Token is valid!")
// }

func CreateToken(username string) (*string, error) {
	secretKey := []byte(config.CurrentConfig.SecretKey)
	claims := Token{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ValidateToken(tokenString string, secretKey []byte) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Token is invalid")
}
