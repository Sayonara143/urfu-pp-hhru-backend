package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) ReviewByID(ctx context.Context, id uuid.UUID) (*models.Review, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	review, err := s.store.ReviewByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return review, s.store.TxCommit(ctx)
}

func (s *Service) Reviews(ctx context.Context, limit, offset int) (int, []*models.Review, error) {
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

	total, reviews, err := s.store.Reviews(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, reviews, s.store.TxCommit(ctx)
}

func (s *Service) ReviewInsert(ctx context.Context, review *models.Review) error {
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

	if err = s.store.ReviewInsert(ctx, review); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) ReviewUpdate(ctx context.Context, review *models.Review) error {
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

	if err = s.store.ReviewUpdate(ctx, review); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) ReviewDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.ReviewDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
