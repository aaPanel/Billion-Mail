package mail_services

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"billionmail-core/api/mail_services/v1"
	"billionmail-core/internal/service/public"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) DeleteBcc(ctx context.Context, req *v1.DeleteBccReq) (res *v1.DeleteBccRes, err error) {
	res = &v1.DeleteBccRes{}

	info, err := g.DB().Model("bm_bcc").Where("id", req.ID).One()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "check record failed: {}", err.Error())))
		return res, nil
	}
	if info.IsEmpty() {
		res.SetError(gerror.New(public.LangCtx(ctx, "record not found")))
		return res, nil
	}

	_, err = g.DB().Model("bm_bcc").Where("id", req.ID).Delete()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "delete bcc rule failed: {}", err.Error())))
		return res, nil
	}

	if err := SyncBccToPostfix(ctx); err != nil {
		g.Log().Error(ctx, "sync bcc config failed: {}", err)
		res.SetError(gerror.New(public.LangCtx(ctx, "delete success but sync config failed: {}", err.Error())))
		return res, nil
	}

	res.SetSuccess(public.LangCtx(ctx, "delete bcc rule success"))
	return res, nil
}
