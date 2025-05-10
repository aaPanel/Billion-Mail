package entity

// BmBcc
type BmBcc struct {
	Id         int    `json:"id"          description:"ID"`
	Type       string `json:"type"        description:"类型: sender或recipient"`
	Address    string `json:"address"     description:"地址"`
	Goto       string `json:"goto"        description:"转发目标"`
	Domain     string `json:"domain"      description:"域名"`
	CreateTime int64  `json:"create_time" description:"创建时间"`
	UpdateTime int64  `json:"update_time" description:"更新时间"`
	Active     int    `json:"active"      description:"状态: 1-启用, 0-禁用"`
}
