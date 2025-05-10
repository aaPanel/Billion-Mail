package mail_services

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"billionmail-core/api/mail_services/v1"
	"billionmail-core/internal/model/entity"
	"billionmail-core/internal/service/public"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetBccList(ctx context.Context, req *v1.GetBccListReq) (res *v1.GetBccListRes, err error) {
	res = &v1.GetBccListRes{}

	model := g.DB().Model("bm_bcc")

	if req.Type != "" {
		model = model.Where("type", req.Type)
	}
	if req.Domain != "" {
		model = model.Where("domain", req.Domain)
	}
	if req.SearchKey != "" {
		model = model.WhereOr("address LIKE ?", "%"+req.SearchKey+"%").
			WhereOr("goto LIKE ?", "%"+req.SearchKey+"%")
	}

	total, err := model.Count()
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "get bcc total failed: {}", err.Error())))
		return res, nil
	}

	list := make([]*entity.BmBcc, 0)
	err = model.Page(req.PageNum, req.PageSize).
		Order("id DESC").
		Scan(&list)
	if err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "get bcc list failed: {}", err.Error())))
		return res, nil
	}

	res.Data.List = make([]v1.BmBcc, 0)
	if err = gconv.Structs(list, &res.Data.List); err != nil {
		res.SetError(gerror.New(public.LangCtx(ctx, "convert data failed: {}", err.Error())))
		return res, nil
	}

	res.Data.Total = total

	res.SetSuccess(public.LangCtx(ctx, "get bcc list success"))
	return res, nil
}
