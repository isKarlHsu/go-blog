package upload_api

import (
	"blog/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"time"
)

type UploadData struct {
	Url string `json:"url"`
	Alt string `json:"alt"`
	Href string `json:"href"`
} 

type UploadRes struct {
	Errno int `json:"errno"`
	Data UploadData `json:"data"` 
}

func (UploadApi) Image(c *gin.Context){
	// 单文件
	file, _ := c.FormFile("uploaded-image")

	extstring := path.Ext(file.Filename)
	fileName := uuid.NewString() + extstring
	folderName := time.Now().Format("2006/01/02")
	dst := "resource/upload/images/" + folderName + "/" + fileName
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		global.Logger.Println(err)
	}

	data := UploadRes{
		Errno: 0,
		Data:  UploadData{
			Url: "http://" + global.Config.System.Addr() + "/" + dst,
		},
	}
	c.JSON(http.StatusOK, data)
}
