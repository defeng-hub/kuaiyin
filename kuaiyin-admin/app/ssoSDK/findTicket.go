package ssoSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

// eg:
// {"ticketFindResult":"true",
// "ssoTicket":"e12de7fed5f666b9b111ba2f28ac244a04e796f7bc06e28d38639b32296e7f3891b69e155d27c5580aedd50b03a10710" ,
// "code":" 07","msg":"存在对应的ticket"}
type FindTicketResponse struct {
	TicketFindResult bool   `json:"ticketFindResult,string"` //请求状态，即本次请求成功或者失败 ：成功为true，失败为false
	SsoTicket        string `json:"ssoTicket"`               //当ticketFindResult为false时，返回值为空字符串； 当ticketFindResult为true时，返回值为客户端对应的票据信息
	Code             string `json:"code"`
	Msg              string `json:"msg"`
}

var findTicketUrl = "/findTicket"

// 获取sso Ticket
func (s *SsoClient) FindTicket() (string, error) {
	postBody := "/" + s.serverIP + "/" + s.clientIP + "/" + s.appName
	reqUrl := s.baseUrl + findTicketUrl + postBody
	//@todo 测试接口
	if config.ApplicationConfig.Mode != "lan" {
		reqUrl = "http://localhost:8008/api/v1/remote/findTicket"
	}
	fmt.Println("----request :", reqUrl)
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
	fmt.Println("----response :", string(body))

	var result FindTicketResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if result.TicketFindResult {
		//return "", errors.New("自定义错误")
		s.ssoTicket = result.SsoTicket
		return result.SsoTicket, nil
	} else {
		return "", fmt.Errorf("sso server find ticket error, msg: %s", result.Msg)
	}
}
