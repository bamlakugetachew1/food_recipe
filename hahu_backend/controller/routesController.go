package controller

import (
	"fmt"
	"hahu_jobs_projects/helper"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
)

func handlerequestdata(cont *gin.Context) (name string, password string) {
	body, err := ioutil.ReadAll(cont.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}
	name = jsonParsed.Path("input.name").String()
	password = jsonParsed.Path("input.password").String()
	unname, err := strconv.Unquote(name)
	unpassword, err := strconv.Unquote(password)
	return unname, unpassword

}

func Registfromhasura(cont *gin.Context) {
	unname, unpassword := handlerequestdata(cont)
	len, id := helper.Checkuser(unname, unpassword)
	fmt.Println(len, id)
	if len != 0 && id != 0 {
		cont.JSON(http.StatusOK, gin.H{"accessToken": "notokengenerated", "success": "you have already registred"})

	} else {
		helper.Saveuser(unname, unpassword, cont)
	}
}

func Logfromhasura(cont *gin.Context) {
	unname, unpassword := handlerequestdata(cont)
	len, id := helper.Checkuser(unname, unpassword)
	if len != 0 {
		cont.JSON(http.StatusOK, gin.H{"accessToken": helper.Generategwt(id), "success": "true", "id": id, "name": unname})

	} else {
		cont.JSON(http.StatusOK, gin.H{"accessToken": "notokengenerated", "success": "you have not  already registred", "id": 0, "name": "noname"})
	}
}

func Uploadfile(c *gin.Context) {
	helper.Uploadfile(c)
}

func Handleuser(c *gin.Context) {
	from := "abuyegetachew2@gmail.com"
	password := "your16charactergoogleapppassword"
	to := []string{
		"abuyeget@gmail.com",
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("A new user is signup to Your site")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
