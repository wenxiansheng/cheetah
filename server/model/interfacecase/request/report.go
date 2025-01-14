package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ReportSearch struct {
	interfacecase.ApiReport
	request.PageInfo
	ApiType int `json:"type" form:"type"`
}

//type InterfaceTemplateApi struct {
//	Name   string
//	Method string
//}
//
//type InterfaceTemplateList struct {
//	ID   uint
//	Name string
//	//Request InterfaceTemplateApi `json:"request"`
//}
