package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopos.com/m/delivery/common_resp"
	"gopos.com/m/delivery/logger"
	"gopos.com/m/utility"
	"io"
	"net/http"
	"os"
)

type testingApi struct {
}

func (t *testingApi) UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file") // key dari form file adalah nama param dari file
		if err != nil {
			logger.SendLogToDiscord("Get file", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusBadRequest, common_resp.FailedMessage(err.Error()))
			return
		}

		fileName := utility.CreateNameFile(header.Filename) // membuat nama file baru
		out, errOut := os.Create("files/" + fileName)       //`membuat file dengan nama yg sudah ditentukan
		if errOut != nil {
			logger.SendLogToDiscord("Create upload file", errOut)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(errOut.Error()))
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			logger.SendLogToDiscord("Copy file", err)
			common_resp.NewCommonResp(c).FailedResp(http.StatusInternalServerError, common_resp.FailedMessage(err.Error()))
			return
		}
		filePath := fmt.Sprintf("http://%s/files/", c.Request.Host)
		common_resp.NewCommonResp(c).SuccessResp(http.StatusOK, common_resp.SuccessMessage("testing", gin.H{
			"filepath": filePath + fileName,
		}))
	}
}

func NewTestingApi(routerGroup *gin.RouterGroup) {
	api := &testingApi{}

	routerGroup.POST("/uploadfile", api.UploadFile())
}
