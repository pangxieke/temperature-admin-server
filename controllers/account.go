package controllers

import (
	"encoding/json"
	"fmt"
	"temp-admin/models"
	"temp-admin/util"
	"time"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Prepare() {

}

type loginUserReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// 请求 && 返回
type loginUserResp struct {
	Token string `json:"token"`
	//UserId   int    `json:"userId"`
	Id       int    `json:"id"`
	Username string `json:"userName" `
	Role     int    `json:"role"`
}

func (this *AccountController) Login() {
	var req loginUserReq
	var resp loginUserResp
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.ReturnJson(401, util.InvaildParameter, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
		return
	}
	if req.UserName == "" || len(req.Password) < 5 {
		this.ReturnJson(401, util.UsernameOrPwdEmpty, string(this.Ctx.Input.RequestBody), util.UsernameOrPwdEmptyMsg)
		return
	}
	storeDao := models.StoreDao{}
	storeInfo, err := storeDao.FindByUserName(req.UserName)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if storeInfo.Id == 0 || storeInfo.Pwd != util.EncryptPwd(req.Password) {
		this.ReturnJson(401, util.UsernameOrPwdEmpty, req, util.UsernameOrPwdErrorMsg)
		return
	}
	if storeInfo.Role > 0 && storeInfo.Id == 0 {
		this.ReturnJson(401, util.NoPower, req, util.NoPowerMsg)
		return
	}
	session, err := models.NewSession(&storeInfo)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	resp.Id = storeInfo.Id
	resp.Username = storeInfo.UserName
	resp.Role = storeInfo.Role
	resp.Token = fmt.Sprintf("%s", session.Token)
	this.ReturnJson(200, util.SUCCESS, resp, util.SUCCESSMsg)
}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Tel      string `json:"tel"`
	Role     int    `json:"role"`
}

func (this *AccountController) Register() {
	var req registerReq
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.ReturnJson(401, util.ServerError, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
		return
	}
	if req.Username == "" || req.Password == "" || req.Name == "" || req.Address == "" || req.Tel == "" {
		this.ReturnJson(401, util.LackParameter, string(this.Ctx.Input.RequestBody), util.LackParameterMsg)
		return
	}
	storeDao := models.StoreDao{}
	store, err := storeDao.FindByUserName(req.Username)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if store.Id > 0 {
		this.ReturnJson(401, util.UserNameExists, "", util.UserNameExistsMsg)
		return
	}

	now := time.Now().Unix()
	info := models.Store{
		UserName:  req.Username,
		Pwd:       util.EncryptPwd(req.Password),
		Name:      req.Name,
		Address:   req.Address,
		Tel:       req.Tel,
		Role:      1,
		Status:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = storeDao.Add(info)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(201, util.SUCCESS, info, util.SUCCESSMsg)
	return
}
