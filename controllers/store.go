package controllers

import (
	"encoding/json"
	"temp-admin/models"
	"temp-admin/util"
	"time"
)

type StoreController struct {
	BaseController
}

type storeEditReq struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Tel      string `json:"tel"`
}

func (this *StoreController) Update() {
	var req storeEditReq
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.ReturnJson(401, util.ServerError, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
		return
	}
	store := this.session.Store
	if store.Role == 1 && (req.Name == "" || req.Id == 0 || req.Address == "" || req.Tel == "") {
		this.ReturnJson(401, util.LackParameter, string(this.Ctx.Input.RequestBody), util.LackParameterMsg)
		return
	}
	store.Id = req.Id
	store.Name = req.Name
	store.Address = req.Address
	store.Tel = req.Tel
	store.UpdatedAt = time.Now().Unix()
	if req.Password != "" {
		store.Pwd = util.EncryptPwd(req.Password)
	}
	storeDao := models.StoreDao{}
	err = storeDao.Update(store)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}

type storeListResp struct {
	ListReq
	Data []models.Store `json:"data"`
}

func (this *StoreController) Lists() {
	size, _ := this.GetInt("limit", 0)
	num, _ := this.GetInt("offset", 0)
	var resp storeListResp
	if this.session.Store.Role == 1 {
		this.ReturnJson(401, util.UserPermissionDend, this.session.Store, util.UserPermissionDendMsg)
		return
	}
	dao := models.StoreDao{}
	lists, total, err := dao.GetList(size, num)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	resp.Total = total
	resp.Offset = num
	resp.Data = make([]models.Store, 0)
	if len(lists) > 0 {
		for _, v := range lists {
			resp.Data = append(resp.Data, v)
		}
	}
	this.ReturnJson(200, util.SUCCESS, resp, util.SUCCESSMsg)
	return
}

func (this *StoreController) Info() {
	dao := models.StoreDao{}
	info, err := dao.FindById(this.session.StoreID)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, info, util.SUCCESSMsg)
	return
}
