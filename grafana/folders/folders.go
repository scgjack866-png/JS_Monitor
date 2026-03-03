package folders

import (
	"OperationAndMonitoring/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"strings"
)

var (
	request = *gorequest.New()
)

func ListFolders(c *gin.Context) {

	limit := c.Query("limit")
	_, body, _ := request.Get(initialize.Grafana.ApiUrl+"/api/folders?limit="+limit).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		End()
	c.String(200, body)
}

func CreateFolder(title string) (string, bool) {

	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/folders").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(`{
 "uid": null,
 "title": "` + title + `"
}`).
		End()
	if res.StatusCode == 200 {
		return body, true
	}

	if res.StatusCode == 409 {

		res1, body1, _ := request.
			Get(initialize.Grafana.ApiUrl+"/api/folders").
			Set("Authorization", initialize.Grafana.Authorization).
			Set("Content-Type", "application/json").
			Set("Accept", "application/json").
			End()

		if res1.StatusCode == 200 {
			index := strings.Index(body1, title)
			startIndex := strings.LastIndex(body1[:index], "{")
			endIndex := strings.LastIndex(body1[index+len(title):], "}")

			return body1[startIndex : index+len(title)+endIndex+1], true
		}

		return body1, false
	}
	return body, false
}

func UpdateFolder(folderId, title string) bool {
	// 结束
	res, body, _ := request.
		Put(initialize.Grafana.ApiUrl+"/api/folders/"+folderId).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(`{
	"title": "` + title + `",
	"overwrite": true
}`).
		End()
	fmt.Println(body, folderId, title)
	return res.StatusCode == 200

}

func DeleteFolder(folders string) bool {

	res, _, _ := request.
		Delete(initialize.Grafana.ApiUrl+"/api/folders/"+folders).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		End()
	return res.StatusCode == 200
}
