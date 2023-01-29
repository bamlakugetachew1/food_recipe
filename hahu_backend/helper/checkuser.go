package helper

import (
	"context"
	"log"
	"github.com/machinebox/graphql"
)

func Checkuser(name string, password string) (length int, id int) {
	type ResponseStruct struct {
		Users []struct {
			Name     string `json:"name"`
			Password string `json:"password"`
			Id       int    `json:"id"`
		}
	}
	client := graphql.NewClient("http://localhost:8081/v1/graphql")
	req := graphql.NewRequest(`
	query myquery($name:String!,$password:String!){
		users( where: {
		   _and: [
			 { name: {_eq: $name}},
			 { password: {_eq: $password}}
		   ]
		 }){
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
	if len(respData.Users) == 0 {
		return len(respData.Users), 0
	}
	return len(respData.Users), respData.Users[0].Id
}
