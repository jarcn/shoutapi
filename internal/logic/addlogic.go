package logic

import (
	"context"

	"github.com/jarcn/bookadd/adder"
	"github.com/jarcn/shoutapi/internal/svc"
	"github.com/jarcn/shoutapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	// 手动代码开始
	r, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: r.Ok,
	}, nil
	// 手动代码结束
}
