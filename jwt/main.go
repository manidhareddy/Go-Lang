package main

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	//"github.com/golang-jwt/jwt/v5/signingmethod"
)
func main() {
	mySigningKey := []byte("AllYourBase")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	name string `json:"name"`
	
	//jwt.RegisteredClaims
}

// Create claims with multiple fields populated
claims := MyCustomClaims{
	"bar",
	"manidhar"
}
	// jwt.RegisteredClaims{
	// 	// A usual scenario is to set the expiration time relative to the current time
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 	IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	NotBefore: jwt.NewNumericDate(time.Now()),
	// 	Issuer:    "test",
	// 	Subject:   "somebody",
	// 	ID:        "1",
	// 	Audience:  []string{"somebody_else"},
	// },
//}

fmt.Printf("foo: %v\n", claims.Foo)

// Create claims while leaving out some of the optional fields
// claims = MyCustomClaims{
// 	"bar",
// 	jwt.RegisteredClaims{
// 		// Also fixed dates can be used for the NumericDate
// 		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
// 		Issuer:    "test",
// 	},
// }

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
ss, err := token.SignedString(mySigningKey)
fmt.Println(ss, err)

}