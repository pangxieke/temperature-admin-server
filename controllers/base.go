package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"runtime"
	"strconv"
	"strings"
	"temp-admin/logs"
	"temp-admin/models"
	"temp-admin/util"
	"time"
)

var TokenOverdueSeconds int

//Init  初始化_controllers
func Init() error {
	return nil

}

type BaseController struct {
	beego.Controller
	VisitId string
	session *models.Session
}

func (this *BaseController) Prepare() {
	this.getAuthorization()
}

func (this *BaseController) Finish() {
	this.StopRun()
}

type ReturnResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type PageSizeReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListReq struct {
	Total  int64 `json:"total"`
	Offset int   `json:"offset"`
}

//鉴权
func (this *BaseController) getAuthorization() {
	this.Authorize()
	return
}

func (this *BaseController) Authorize() {
	authorization := this.Ctx.Input.Header("Authorization")
	if authorization == "" {
		this.ReturnJson(401, util.LackParameter, "Authorization is empty", util.LackParameterMsg)
		return
	}
	s := strings.Split(authorization, " ")
	if len(s) < 2 {
		this.ReturnJson(401, util.InvaildParameter, "invalid Authorization format", util.InvaildParameterMsg)
		return
	}
	typ, token := s[0], s[1]
	if typ != "Bearer" {
		this.ReturnJson(401, util.InvaildParameter, fmt.Sprintf("invalid Authorization type: `%s`, only support Bearer", s[0]), util.InvaildParameterMsg)
		return
	}
	session, err := models.LoadSession(token)
	if err != nil || session == nil {
		this.ReturnJson(401, util.UserPermissionExpiry, err, util.UserPermissionExpiryMsg)
		return
	}
	if session.StoreID == 0 {
		this.ReturnJson(401, util.UserPermissionExpiry, "invalid token", util.UserPermissionExpiryMsg)
		return
	}
	this.session = session
}

func (this *BaseController) LogOut() {
	if this.session != nil {
		this.session.Token.Claims["exp"] = time.Now()
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}
func (this *BaseController) ReturnJson(httpCode, code int, response interface{}, msg string) {
	var data ReturnResult
	data.Code = code
	data.Msg = msg
	data.Data = ""
	if code == 0 {
		data.Data = response
	}

	this.Ctx.Output.Status = httpCode
	_ = this.Ctx.Output.JSON(data, false, false)

	//日志记录
	_, file, line, _ := runtime.Caller(1)
	FileLine := file + " Line " + strconv.Itoa(line)
	logMap := make(map[string]interface{}, 0)
	logMap["Uri"] = this.Ctx.Input.URI()
	logMap["Method"] = this.Ctx.Input.Method()
	if "POST" == logMap["Method"] {
		logMap["Request"] = string(this.Ctx.Input.RequestBody)
	}
	logMap["FileLen"] = FileLine
	logMap["HttpCode"] = httpCode
	logMap["Code"] = code
	logMap["Msg"] = msg
	logMap["Response"] = response
	logMap["Header"] = struct {
		Authorization string
	}{
		Authorization: this.Ctx.Input.Header("Authorization"),
	}
	logByte, _ := json.Marshal(logMap)
	logs.Info(logs.FILE|logs.CONSO, "%s", logByte)
	this.StopRun()
	return
}
