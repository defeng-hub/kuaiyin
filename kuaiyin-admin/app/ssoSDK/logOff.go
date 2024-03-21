package ssoSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

type LogOffResponse struct {
	LogOffResult bool   `json:"logOffResult,string"`
	Code         string `json:"code"`
	Msg          string `json:"msg"`
}

var logOffUrl = "/logOff"

// 退出
func (s *SsoClient) LogOff(pcNumber string) (bool, error) {
	postBody := "/" + s.serverIP + "/" + s.clientIP + "/" + s.appName + "/" + pcNumber
	reqUrl := s.baseUrl + logOffUrl + postBody
	//@todo 测试接口
	if config.ApplicationConfig.Mode != "lan" {
		reqUrl = "http://localhost:8008/api/v1/remote/logOff"
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return false, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	fmt.Println("----request :", reqUrl)
	fmt.Println("----response :", string(body))

	var result LogOffResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}

	if result.LogOffResult {
		return true, nil
	} else {
		return false, fmt.Errorf("sso server error, msg: %s", result.Msg)
	}
}

func (s *SsoClient) RedirectUrl(url string) string {
	if config.ApplicationConfig.Mode != "lan" {
		//return "http://renxian.linuxdev.cn/ssologin"
		return "/"
	} else {
		return "http://11.36.19.209:8085/jzsso/login.jsp?appServerIP=" + s.serverIP + "&appName=" + s.appName + "&casRedirectUrl=" + url
	}
}
