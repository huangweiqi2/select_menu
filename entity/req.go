package entity

import "errors"

type RandomReq struct {
	HotNum  int `form:"hot_num"`
	ColdNum int `form:"cold_num"`
	SoupNum int `form:"soup_num"`
}

func (r RandomReq) Valid() error {
	if r.HotNum == 0 && r.ColdNum == 0 && r.SoupNum == 0 {
		return errors.New("参数错误")
	}
	return nil
}
