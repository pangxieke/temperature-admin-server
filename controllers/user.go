package controllers

import (
	"encoding/json"
	"fmt"
	"temp-admin/models"
	"temp-admin/util"
)

type UserControllers struct {
	BaseController
}

type userListResp struct {
	ListReq
	Data []models.User `json:"data"`
}

func (this *UserControllers) Lists() {
	q := make(map[string]interface{}, 0)
	q["limit"], _ = this.GetInt("limit", 15)
	q["offset"], _ = this.GetInt("offset", 0)
	q["name"] = this.GetString("name", "")
	//
	//size, _ := this.GetInt("limit", 15)
	//num, _ := this.GetInt("offset", 0)
	var resp userListResp
	storeId := this.session.StoreID
	if this.session.Store.Role == 0 {
		storeId = 0
	}
	dao := models.UserDao{}
	dev, total, err := dao.GetList(storeId, q)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	resp.Total = total
	resp.Offset = q["offset"].(int)
	resp.Data = make([]models.User, 0)
	if len(dev) > 0 {
		for _, v := range dev {
			resp.Data = append(resp.Data, v)
		}
	}
	this.ReturnJson(200, util.SUCCESS, resp, util.SUCCESSMsg)
	return
}

func (this *UserControllers) Update() {
	var req models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.ReturnJson(401, util.ServerError, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
		return
	}
	store := this.session.Store
	if req.Name == "" || req.Id == 0 || req.Tel == "" {
		this.ReturnJson(401, util.LackParameter, string(this.Ctx.Input.RequestBody), util.LackParameterMsg)
		return
	}
	dao := models.UserDao{}
	info, err := dao.FindById(req.Id)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if info.Id == 0 {
		this.ReturnJson(401, util.NotFound, this.Ctx.Input.URI(), util.NotFoundMsg)
		return
	}
	if store.Role == 1 && store.Id != info.CompanyId {
		this.ReturnJson(401, util.UserPermissionDend, fmt.Sprintf("store:%+v, info:%+v", store, info), util.UserPermissionDendMsg)
		return
	}
	info.Name = req.Name
	info.Age = req.Age
	info.IdNum = req.IdNum
	info.Tel = req.Tel
	info.Department = req.Department
	err = dao.Update(info)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}

func (this *UserControllers) Del() {
	id, _ := this.GetInt(":id", 0)
	if id == 0 {
		this.ReturnJson(401, util.LackParameter, this.Ctx.Input.URI(), util.LackParameterMsg)
		return
	}
	dao := models.UserDao{}
	info, err := dao.FindById(id)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if info.Id == 0 {
		this.ReturnJson(401, util.NotFound, this.Ctx.Input.URI(), util.NotFoundMsg)
		return
	}
	store := this.session.Store
	if store.Role == 1 && store.Id != info.CompanyId {
		this.ReturnJson(401, util.UserPermissionDend, fmt.Sprintf("store:%+v, info:%+v", store, info), util.UserPermissionDendMsg)
		return
	}
	err = dao.Del(id)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}
