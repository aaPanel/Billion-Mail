package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type ServiceType string

const (
	ServiceTypePostfix ServiceType = "postfix"
	ServiceTypeDovecot ServiceType = "dovecot"
	ServiceTypeRspamd  ServiceType = "rspamd"
)

const (
	ServiceType_Postfix = "/opt/Billion-Mail/conf/postfix/main.cf"
	ServiceType_Dovecot = "/opt/Billion-Mail/conf/dovecot/dovecot.conf"
	ServiceType_Rspamd  = "/opt/Billion-Mail/conf/rspamd/rspamd.conf"
)

type GetConfigFileReq struct {
	g.Meta      `path:"/services/get_config" method:"post" summary:"get config file"`
	ServiceType ServiceType `json:"service_type" v:"required" desc:"service type"`
}

type GetConfigFileRes struct {
	api_v1.StandardRes
	Data struct {
		ConfigPath   string `json:"config_path"`   // config file path
		Content      string `json:"content"`       // file content
		Encoding     string `json:"encoding"`      // file encoding
		LastModified int64  `json:"last_modified"` // last modified time
		Size         int64  `json:"size"`          // file size
		OnlyRead     bool   `json:"only_read"`     // only read (file too large)
	} `json:"data"`
}

type SaveConfigFileReq struct {
	g.Meta      `path:"/services/save_config" method:"post" summary:"save config file"`
	ServiceType ServiceType `json:"service_type" v:"required" desc:"service type"`
	Content     string      `json:"content" v:"required" desc:"file content"`
}

type SaveConfigFileRes struct {
	api_v1.StandardRes
}
