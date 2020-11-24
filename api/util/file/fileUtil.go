package file

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/xuese-go/goStudy01/api/util/path"
	"log"
	"path/filepath"
)

func Up(context *gin.Context) string {
	if file, err := context.FormFile("file"); err != nil {
		log.Panicln(err.Error())
		return ""
	} else {
		fin := filepath.Ext(file.Filename)
		fileName := uuid.NewV4().String() + fin
		if err := context.SaveUploadedFile(file, path.PATH+fileName); err != nil {
			log.Panicln(err.Error())
			return ""
		} else {
			return fileName
		}
	}
}