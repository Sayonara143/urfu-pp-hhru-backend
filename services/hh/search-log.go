package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
)

func (s *Service) SearchLogInsert(ctx context.Context, log *models.SearchLog) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err = s.store.SearchLogInsert(ctx, log); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) SearchLogs(ctx context.Context, limit, offset int) (int, []*models.SearchLog, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	total, logs, err := s.store.SearchLogs(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, logs, s.store.TxCommit(ctx)
}
