package mail_services

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/mail_services/v1"
)

func (c *ControllerV1) DeleteMailForward(ctx context.Context, req *v1.DeleteMailForwardReq) (res *v1.DeleteMailForwardRes, err error) {
	res = &v1.DeleteMailForwardRes{}

	count, err := g.DB().Model("alias").Where("address=?", req.Address).Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "check address failed: {}", err.Error())))
		return res, nil
	}

	if count == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "forward address not found")))
		return res, nil
	}

	_, err = g.DB().Model("alias").Where("address=?", req.Address).Delete()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "delete mail forward failed: {}", err.Error())))
		return res, nil
	}

	res.SetSuccess(public.LangCtx(ctx, "delete mail forward success"))
	return res, nil
}
