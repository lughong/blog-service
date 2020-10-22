package service

import (
	"context"

	"github.com/lughong/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context, dao *dao.Dao) Service {
	return Service{
		ctx: ctx,
		dao: dao,
	}
}
