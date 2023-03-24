package logic

import (
	"context"

	"github.com/jarcn/shoutapi/internal/svc"
	"github.com/jarcn/shoutapi/internal/types"
	"github.com/jarcn/shoutrpc/transform"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (types.ExpandResp, error) {
	// 调用rpc服务中的expand服务
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return types.ExpandResp{}, err
	}
	return types.ExpandResp{
		Url: resp.Url,
	}, nil
}
