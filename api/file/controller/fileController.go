package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/util/file"
)

/**
上传文件
*/
func Up(context *gin.Context) {
	if str := file.Up(context); str != "" {
		resp.Respone(context, resp.ResponeStruct{Success: true, Msg: "上传成功", Data: str})
	} else {
		resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "上传失败"})
	}
	//if file, err := context.FormFile("file"); err != nil {
	//	resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "上传失败"})
	//	log.SugarLogger.Errorf(err.Error())
	//} else {
	//	fin := filepath.Ext(file.Filename)
	//	fileName := uuid.NewV4().String() + fin
	//	if err := context.SaveUploadedFile(file, path.PATH+fileName); err != nil {
	//		resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "上传失败"})
	//		log.SugarLogger.Errorf(err.Error())
	//	} else {
	//		resp.Respone(context, resp.ResponeStruct{Success: true, Msg: "上传成功", Data: fileName})
	//	}
	//}
}

/**
下载文件
*/
func Dow(context *gin.Context) {

}
