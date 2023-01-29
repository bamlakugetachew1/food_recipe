package helper

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

func Saveuser(name string, password string, cont *gin.Context) {
	type ResponseStruct struct {
		Insert_users_one struct {
			Name     string `json:"name"`
			Password string `json:"password"`
			Id       int    `json:"id"`
		}
	}
	client := graphql.NewClient("http://localhost:8081/v1/graphql")
	req := graphql.NewRequest(`
	mutation insert_users_one($name:String!,$password:String){
		insert_users_one(object:{name:$name,password:$password}){
		  name
		  password
		  id
	  }
	}
    `)
	req.Var("name", name)
	req.Var("password", password)
	ctx := context.Background()
	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	cont.JSON(http.StatusCreated, gin.H{"accessToken": Generategwt(respData.Insert_users_one.Id), "success": "You have Successfully registred please Login"})
	return
}
