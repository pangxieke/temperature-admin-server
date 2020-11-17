package controllers

import (
	"temp-admin/models"
	"temp-admin/util"
)

type TempControllers struct {
	BaseController
}
type tempListResp struct {
	ListReq
	Data []models.Record `json:"data"`
}

func (this TempControllers) Lists() {
	q := make(map[string]interface{}, 0)
	q["limit"], _ = this.GetInt("limit", 15)
	q["offset"], _ = this.GetInt("offset", 0)
	q["start_time"], _ = this.GetInt("start_time", 0)
	q["end_time"], _ = this.GetInt("end_time", 0)
	q["num"], _ = this.GetFloat("num", 0)
	q["name"] = this.GetString("name", "")
	q["id_num"] = this.GetString("id_num", "")
	q["user_no"] = this.GetString("user_no", "")

	var resp tempListResp
	storeId := this.session.StoreID
	if this.session.Store.Role == 0 {
		storeId = 0
	}
	dao := models.RecordDao{}
	dev, total, err := dao.GetList(storeId, q)
	if err != nil {
		this.ReturnJson(401, util.ServerError, err.Error(), util.ServerErrorMsg)
		return
	}
	resp.Total = total
	resp.Offset = q["offset"].(int)
	resp.Data = make([]models.Record, 0)
	if len(dev) > 0 {
		for _, v := range dev {
			resp.Data = append(resp.Data, v)
		}
	}
	this.ReturnJson(200, util.SUCCESS, resp, util.SUCCESSMsg)
	return
}
