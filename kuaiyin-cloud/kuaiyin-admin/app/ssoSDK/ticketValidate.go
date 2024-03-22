package ssoSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

type TicketValidateResponse struct {
	TicketValidateResult bool   `json:"ticketValidateResult,string"`
	Code                 string `json:"code"`
	Msg                  string `json:"msg"`
}

var ticketValidateUrl = "/ticketValidate"

// 获取sso Ticket
func (s *SsoClient) TicketValidate(ticket string) (bool, error) {
	postBody := "/" + ticket + "/" + s.ticketKey + "/" + s.serverIP + "/" + s.clientIP + "/" + s.appName
	reqUrl := s.baseUrl + getTicketUrl + postBody
	//@todo 测试接口
	if config.ApplicationConfig.Mode != "lan" {
		reqUrl = "http://localhost:8008/api/v1/remote/ticketValidate"
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

	var result TicketValidateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}
	if result.TicketValidateResult {
		return true, nil
	} else {
		return false, fmt.Errorf("sso server validate ticket error, msg: %s", result.Msg)
	}
}
