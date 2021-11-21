package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthenticated(endpoint func(w http.ResponseWriter, r *http.Request, userID int)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(os.Getenv("SECRET")), nil
			})
			if err != nil {
				fmt.Println(w, err.Error())
			}
			if token.Valid {
				Claims := token.Claims
				userIDFloat64 := Claims.(jwt.MapClaims)["userID"]
				//convert float64 to int

				var userID int = int(userIDFloat64.(float64))
				fmt.Printf("%T\n", userID)
				fmt.Println(userID)
				endpoint(w, r, userID)
			} else {
				fmt.Fprintf(w, "Not Authorized")
			}
		}
	})
}
