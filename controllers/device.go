package controllers

import (
	"encoding/json"
	"temp-admin/models"
	"temp-admin/util"
)

type DeviceControllers struct {
	BaseController
}

type devlistResp struct {
	ListReq
	Data []models.DeviceStore `json:"data"`
}

func (this *DeviceControllers) Lists() {
	size, _ := this.GetInt("limit", 15)
	num, _ := this.GetInt("offset", 0)
	var resp devlistResp
	//err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	//if err != nil {
	//	this.ReturnJson(401, util.InvaildParameter, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
	//	return
	//}
	storeId := this.session.StoreID
	if this.session.Store.Role == 0 {
		storeId = 0
	}
	dao := models.DeviceDao{}
	resp.Data = make([]models.DeviceStore, 0)
	dev, total, err := dao.GetList(size, num, storeId)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	resp.Total = total
	resp.Offset = num
	if total > 0 {
		resp.Data = dev
	}
	//if len(dev) > 0 {
	//	for _, v := range dev {
	//		item := devitemResp{
	//			Id:        v.Id,
	//			Sn:        v.Sn,
	//			Remake:    v.Remake,
	//			Status:    v.Status,
	//			Province:  v.Province,
	//			City:      v.City,
	//			CreatedAt: v.CreatedAt,
	//			StoreId:   v.StoreId,
	//			StoreName: v.StoreName,
	//		}
	//		resp.Data = append(resp.Data, item)
	//	}
	//}
	this.ReturnJson(200, util.SUCCESS, resp, util.SUCCESSMsg)
	return
}

type devbindReq struct {
	StoreId  int    `json:"store_id"`
	Sn       string `json:"sn"`
	Province int    `json:"province"`
	City     int    `json:"city"`
	Remake   string `json:"remake"`
}

func (this *DeviceControllers) Bind() {
	var req devbindReq
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.ReturnJson(401, util.InvaildParameter, string(this.Ctx.Input.RequestBody), util.InvaildParameterFormatMsg)
		return
	}
	if req.Sn == "" {
		this.ReturnJson(401, util.LackParameter, string(this.Ctx.Input.RequestBody), util.LackParameterMsg)
		return
	}
	sess := this.session.Store
	storeId := req.StoreId
	if sess.Role == 1 && sess.Id != storeId {
		this.ReturnJson(401, util.UserPermissionDend, this.session.Store, util.UserPermissionDendMsg)
		return
	}
	dao := models.DeviceDao{}
	info, err := dao.FindBySn(req.Sn)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if info.Id == 0 {
		this.ReturnJson(401, util.NotFound, req, util.NotFoundMsg)
		return
	}

	if sess.Role == 1 && info.StoreId > 0 && info.StoreId != storeId {
		this.ReturnJson(401, util.DeviceBinded, req, util.DeviceBindedMsg)
		return
	}

	info.StoreId = storeId
	info.Remake = req.Remake
	info.City = req.City
	info.Province = req.Province
	err = dao.Bind(info)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}

func (this *DeviceControllers) UnBind() {
	sn := this.GetString(":sn", "")
	if sn == "" {
		this.ReturnJson(401, util.LackParameter, this.Ctx.Input.URI(), util.LackParameterMsg)
		return
	}
	dao := models.DeviceDao{}
	info, err := dao.FindBySn(sn)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	if info.Id == 0 {
		this.ReturnJson(401, util.NotFound, sn, util.NotFoundMsg)
		return
	}
	storeId := this.session.StoreID
	if this.session.Store.Role == 1 && info.StoreId != storeId {
		this.ReturnJson(401, util.UserPermissionDend, this.session.Store, util.UserPermissionDendMsg)
		return
	}

	info.StoreId = 0
	info.City = 0
	info.Province = 0
	err = dao.Bind(info)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	this.ReturnJson(200, util.SUCCESS, "", util.SUCCESSMsg)
	return
}
