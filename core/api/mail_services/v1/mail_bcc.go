package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type BmBcc struct {
	ID         int    `json:"id"         dc:"ID"`
	Type       string `json:"type"       dc:"type: sender/recipient"`
	Address    string `json:"address"    dc:"address"`
	Goto       string `json:"goto"       dc:"Forwarding target"`
	Domain     string `json:"domain"     dc:"domain"`
	CreateTime int64  `json:"create_time" dc:"create time"`
	UpdateTime int64  `json:"update_time" dc:"update time"`
	Active     int    `json:"active"     dc:"Status: 1-enabled, 0-disabled"`
}

type GetBccListReq struct {
	g.Meta    `path:"/mail_bcc/list" method:"get" summary:"get mail bcc list"`
	Type      string `json:"type" desc:"bcc type: sender/recipient, not set means get all"`
	Domain    string `json:"domain" desc:"domain filter"`
	SearchKey string `json:"search_key" desc:"search key"`
	PageNum   int    `json:"page_num" d:"1" desc:"page num"`
	PageSize  int    `json:"page_size" d:"20" desc:"page size"`
}

type GetBccListRes struct {
	api_v1.StandardRes
	Data struct {
		List  []BmBcc `json:"list" dc:"BCC list"`
		Total int     `json:"total" dc:"total"`
	} `json:"data"`
}

type AddBccReq struct {
	g.Meta  `path:"/mail_bcc/add" method:"post" summary:"add mail bcc"`
	Type    string `json:"type" v:"required|in:sender,recipient" dc:"bcc type: sender or recipient"`
	Address string `json:"address" v:"required" dc:"email address or domain"`
	Goto    string `json:"goto" v:"required" dc:"forward target address"`
	Domain  string `json:"domain" dc:"domain"`
	Active  int    `json:"active" d:"1" dc:"status: 1-enabled, 0-disabled"`
}

type AddBccRes struct {
	api_v1.StandardRes
}

type EditBccReq struct {
	g.Meta  `path:"/mail_bcc/edit" method:"post" summary:"edit mail bcc"`
	ID      int    `json:"id" v:"required" dc:"record id"`
	Type    string `json:"type" v:"in:sender,recipient" dc:"bcc typeype: sen or er or recipient"`
	Address string `json:"address"  dc:"email address or domainddress or domain"`
	Goto    string `json:"goto"  dc:"forward target addressd target address"`
	Domain  string `json:"domain" dc:"domain"`
	Active  int    `json:"active" v:"required" d:"1" dc:"status: 1-enabled, 0-disabled"`
}

type EditBccRes struct {
	api_v1.StandardRes
}

type DeleteBccReq struct {
	g.Meta `path:"/mail_bcc/delete" method:"post" summary:"delete mail bcc"`
	ID     int `json:"id" v:"required" dc:"ID"`
}

type DeleteBccRes struct {
	api_v1.StandardRes
}
