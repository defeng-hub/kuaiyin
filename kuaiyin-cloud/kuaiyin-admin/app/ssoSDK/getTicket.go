package ssoSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

type GetTicketResponse struct {
	TicketGenerateResult bool   `json:"ticketGenerateResult,string"`
	Ticket               string `json:"ticket"`
	Code                 string `json:"code"`
	Msg                  string `json:"msg"`
}

var getTicketUrl = "/getTicket"

// 获取sso Ticket
func (s *SsoClient) GetTicket() (string, error) {
	postBody := "/" + s.serverIP + "/" + s.clientIP + "/" + s.appName
	reqUrl := s.baseUrl + getTicketUrl + postBody
	//@todo 测试接口
	if config.ApplicationConfig.Mode != "lan" {
		reqUrl = "http://localhost:8008/api/v1/remote/getTicket"
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("----request :", reqUrl)
	fmt.Println("----response :", string(body))
	var result GetTicketResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	//成功拿到ticket
	if result.TicketGenerateResult {
		//return "", errors.New("自定义错误")
		s.ssoTicket = result.Ticket
		return result.Ticket, nil
	} else {
		return "", fmt.Errorf("sso server get ticket error, msg: %s", result.Msg)
	}
}
