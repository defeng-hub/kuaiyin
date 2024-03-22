package ssoSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

type UserInfoResponse struct {
	UserInfoResult bool   `json:"userInfoResult,string"`
	Code           string `json:"code"`
	Msg            string `json:"msg"`
	UserInfo       *UserInfo
}

type UserInfo struct {
	IdcardNo string `json:"idcard_no"` //身份证号
	RealName string `json:"realName"`  //真实姓名
	UserId   string `json:"userId"`
	UserName string `json:"userName"` //用户名
	PcNumber string `json:"pcNumber"` //警官号
}

var userInfoUrl = "/userInfo"

// 获取sso Ticket
func (s *SsoClient) UserInfo() (*UserInfo, error) {
	postBody := "/" + s.ssoTicket + "/" + s.serverIP + "/" + s.clientIP + "/" + s.appName
	reqUrl := s.baseUrl + userInfoUrl + postBody
	//@todo 测试接口
	if config.ApplicationConfig.Mode != "lan" {
		reqUrl = "http://localhost:8008/api/v1/remote/userInfo"
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("----request :", reqUrl)
	fmt.Println("----response :", string(body))

	var result UserInfoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.UserInfoResult {
		return result.UserInfo, nil
	} else {
		return nil, fmt.Errorf("sso server error, msg: %s", result.Msg)
	}
}
