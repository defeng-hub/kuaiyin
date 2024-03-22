package ssoSDK

import (
	"net/http"
)

type SsoClient struct {
	client    *http.Client
	baseUrl   string
	ssoTicket string
	serverIP  string
	clientIP  string
	appName   string
	ticketKey string
}

func NewSsoClient(clientIP string) *SsoClient {
	if clientIP == "::1" || clientIP == "127.0.0.1" {
		clientIP = "10.48.105.118"
	}
	return &SsoClient{
		client:    &http.Client{},
		baseUrl:   "http://11.36.19.209:8085/jzsso/ticket",
		ticketKey: "",
		appName:   "ywshxjffx",
		serverIP:  "10.48.105.118",
		clientIP:  clientIP,
	}
}
