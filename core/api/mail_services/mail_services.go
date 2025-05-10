// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package mail_services

import (
	"context"

	"billionmail-core/api/mail_services/v1"
)

type IMailServicesV1 interface {
	GetConfigFile(ctx context.Context, req *v1.GetConfigFileReq) (res *v1.GetConfigFileRes, err error)
	SaveConfigFile(ctx context.Context, req *v1.SaveConfigFileReq) (res *v1.SaveConfigFileRes, err error)
	GetBccList(ctx context.Context, req *v1.GetBccListReq) (res *v1.GetBccListRes, err error)
	AddBcc(ctx context.Context, req *v1.AddBccReq) (res *v1.AddBccRes, err error)
	EditBcc(ctx context.Context, req *v1.EditBccReq) (res *v1.EditBccRes, err error)
	DeleteBcc(ctx context.Context, req *v1.DeleteBccReq) (res *v1.DeleteBccRes, err error)
	GetMailForwardList(ctx context.Context, req *v1.GetMailForwardListReq) (res *v1.GetMailForwardListRes, err error)
	AddMailForward(ctx context.Context, req *v1.AddMailForwardReq) (res *v1.AddMailForwardRes, err error)
	EditMailForward(ctx context.Context, req *v1.EditMailForwardReq) (res *v1.EditMailForwardRes, err error)
	DeleteMailForward(ctx context.Context, req *v1.DeleteMailForwardReq) (res *v1.DeleteMailForwardRes, err error)
}
