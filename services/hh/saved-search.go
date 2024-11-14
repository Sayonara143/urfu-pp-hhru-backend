package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) SavedSearchByID(ctx context.Context, id uuid.UUID) (*models.SavedSearch, error) {
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

	search, err := s.store.SavedSearchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return search, s.store.TxCommit(ctx)
}

func (s *Service) SavedSearches(ctx context.Context, limit, offset int) (int, []*models.SavedSearch, error) {
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

	total, searches, err := s.store.SavedSearches(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, searches, s.store.TxCommit(ctx)
}

func (s *Service) SavedSearchInsert(ctx context.Context, search *models.SavedSearch) error {
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

	if err = s.store.SavedSearchInsert(ctx, search); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) SavedSearchDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.SavedSearchDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
