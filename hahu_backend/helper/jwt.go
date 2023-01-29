package helper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generategwt(id int) (datas string) {
	idstr := strconv.Itoa(id)
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"eat":   time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
		"sub":   "1234567890",
		"admin": false,
		"name":  "John Doe",
		"iat":   1516239022,
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []interface{}{"public", "user", "admin"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       idstr,
			"x-hasura-org-id":        "123",
			"x-hasura-custom":        "custom-value",
		},
	})
	datas, err := token.SignedString([]byte("YOURJWT32CHARACTERSTOKENPRIVATEKEY"))
	if err != nil {
		fmt.Println(err)
	}
	return datas
}
