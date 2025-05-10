package mail_services

import (
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/mail_services/v1"
)

func (c *ControllerV1) EditMailForward(ctx context.Context, req *v1.EditMailForwardReq) (res *v1.EditMailForwardRes, err error) {
	res = &v1.EditMailForwardRes{}

	count, err := g.DB().Model("alias").Where("address=?", req.Address).Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "check address failed: {}", err.Error())))
		return res, nil
	}

	if count == 0 {
		res.SetError(gerror.New(public.LangCtx(ctx, "forward address not found")))
		return res, nil
	}

	update := g.Map{
		"update_time": time.Now().Unix(),
		"active":      req.Active,
	}

	// if provided ForwardUsers, update target address
	if req.Goto != "" {
		// process forward target addresses
		forwardUsers := processForwardUsers(req.Goto)
		if len(forwardUsers) == 0 {
			res.SetError(gerror.New(public.LangCtx(ctx, "forward target address cannot be empty")))
			return res, nil
		}

		// add forward target to update parameters
		update["goto"] = strings.Join(forwardUsers, ",")
	}

	_, err = g.DB().Model("alias").Where("address=?", req.Address).Update(update)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "modify mail forward failed: {}", err.Error())))
		return res, nil
	}

	res.SetSuccess(public.LangCtx(ctx, "modify mail forward success"))
	return res, nil
}
