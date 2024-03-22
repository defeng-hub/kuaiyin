package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/app/ssoSDK"
	"go-admin/common"
	"gorm.io/gorm"
	"io"
	"net/http"

	"go-admin/common/global"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/mssola/user_agent"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		r, _ := v["role"].(SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey:  u.UserId,
			jwt.RoleIdKey:    r.RoleId,
			jwt.RoleKey:      r.RoleKey,
			jwt.NiceKey:      u.NickName,
			jwt.DataScopeKey: r.DataScope,
			jwt.RoleNameKey:  r.RoleName,
			jwt.DeptId:       u.DeptId,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
		"DataScope":   claims["datascope"],
		"DeptId":      claims["deptid"],
	}
}

// Authenticator 获取token
// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Tags 登陆
// @Accept  application/json
// @Product application/json
// @Param account body Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /api/v1/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var err error
	// 添加登录方式
	//1. 账户密码登录
	//2. 通过sso,获取的警官号登录，
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	var loginType LoginTypeDto

	err = json.Unmarshal(body, &loginType)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	switch loginType.LoginType {
	case SSO:
		return SSOLogin(c, loginType.Overload)
	case Normal:
		return NormalLogin(c)
	default:
		return NormalLogin(c)
	}
}

func NormalLogin(c *gin.Context) (interface{}, error) {
	var err error

	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	var loginVals Login
	if err = c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		status = "1"
		return nil, jwt.ErrFailedAuthentication
	}

	if config.ApplicationConfig.Mode != "dev" {
		if !captcha.Verify(loginVals.UUID, loginVals.Code, true) {
			username = loginVals.Username
			msg = "验证码错误"
			status = "1"

			return nil, jwt.ErrFailedAuthentication
		}
	}

	userT, role, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username
		return map[string]interface{}{"user": userT, "role": role}, nil
	} else {
		msg = "登录失败"
		status = "1"
		log.Warnf("%s login failed!", loginVals.Username)
	}
	return nil, jwt.ErrFailedAuthentication
}

func RegisterUser(info *ssoSDK.UserInfo, db *gorm.DB) error {
	s := service.SysUser{}
	req := dto.SysUserRegisterReq{}
	s.Orm = db
	req.SetCreateBy(0)
	req.NickName = info.UserName
	req.Username = info.UserName
	// 设置创建人
	req.PcNumber = info.PcNumber
	req.RoleId = 2   //role 民警管理员
	req.DeptId = 11  //民警机构
	req.Status = "2" //default is enable

	err := s.Register(&req)
	if err != nil {
		return err
	}
	return nil
}

func SSOLogin(c *gin.Context, loginType bool) (interface{}, error) {
	var ticket string
	var err error
	loginErr := errors.New("sso login fail")
	//test
	ssoClient := ssoSDK.NewSsoClient(common.GetClientIP(c))
	ticket, err = ssoClient.FindTicket()
	if loginType && ticket == "" { //慢登录仍然拿不到ticket，提示登录失败
		return nil, errors.New("overload ticket")
	}
	if err != nil || ticket == "" {
		domain := "http://10.48.105.118:88/ssologin?overload=true" //TODO: 待改为线上实际的地址
		url := ssoClient.RedirectUrl(domain)
		return nil, errors.New(fmt.Sprintf("get ticket failed:[%s]", url))
	}
	ssoUserInfo, err := ssoClient.UserInfo()
	if err != nil {
		return nil, loginErr
	}
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, err
	}

	var status = "2"
	var msg = "sso登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	var loginVals LoginPcNumberDto
	loginVals.PcNumber = ssoUserInfo.PcNumber

	userT, role, err := loginVals.GetUserByPcNumber(db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建逻辑
			err1 := RegisterUser(ssoUserInfo, db)
			if err1 != nil {
				return nil, errors.Join(loginErr, err1)
			} else {
				//	注册成功了
				userT1, role1, err2 := loginVals.GetUserByPcNumber(db)
				if err2 != nil {
					username = loginVals.PcNumber
					msg = "登录失败"
					status = "1"
					return nil, errors.Join(loginErr, err1)
				}
				return map[string]interface{}{"user": userT1, "role": role1}, nil
			}
		} else {
			username = loginVals.PcNumber
			msg = "sso登录失败:该警号不存在"
			status = "1"
			return nil, loginErr
		}
	} else {
		return map[string]interface{}{"user": userT, "role": role}, nil
	}
}

// LoginLogToDB Write log to database
func LoginLogToDB(c *gin.Context, status string, msg string, username string) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = common.GetClientIP(c)
	l["loginLocation"] = "" // pkg.GetLocation(common.GetClientIP(c),gaConfig.ExtConfig.AMap.Key)
	l["loginTime"] = pkg.GetCurrentTime()
	l["status"] = status
	l["remark"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["username"] = username
	l["msg"] = msg

	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", global.LoginLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}

// LogOut
// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security Bearer
func LogOut(c *gin.Context) {
	LoginLogToDB(c, "2", "退出成功", user.GetUserName(c))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(models.SysUser)
		r, _ := v["role"].(models.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("deptId", u.DeptId)
		c.Set("userName", u.Username)
		c.Set("dataScope", r.DataScope)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
