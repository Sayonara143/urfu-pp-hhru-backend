package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/pagination"
)

type Option func(s *Service)

type Service struct {
	store Storage

	pagination Pagination
}

func New(store Storage, opts ...Option) *Service {
	s := &Service{
		store: store,
	}

	for _, applyOpt := range opts {
		applyOpt(s)
	}

	if s.pagination == nil {
		s.pagination = new(paginationImpl)
	}

	return s
}

type paginationImpl struct{}

func (p *paginationImpl) Get(ctx context.Context) Paginator {
	return pagination.Get(ctx)
}

func WithPaginationMock(pagination Pagination) Option {
	return func(s *Service) {
		s.pagination = pagination
	}
}
