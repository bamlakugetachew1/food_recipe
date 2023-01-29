package helper

import (
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Uploadfile(c *gin.Context) {
	var concatFileName = ""
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]
	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		concatFileName += newFileName + ","
		localTargetFilePath := "images/" + newFileName

		if err := c.SaveUploadedFile(file, localTargetFilePath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Your file has been successfully uploaded.",
		"fileslocations": concatFileName,
	})
}
