package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) SavedSearchByID(ctx context.Context, id uuid.UUID) (*models.SavedSearch, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var search models.SavedSearch
	if err := tx.NewSelect().
		Model(&search).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &search, nil
}

func (s *Storage) SavedSearches(ctx context.Context, limit, offset int) (int, []*models.SavedSearch, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var searches []*models.SavedSearch
	total, err := tx.NewSelect().
		Model(&searches).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, searches, nil
}

func (s *Storage) SavedSearchInsert(ctx context.Context, search *models.SavedSearch) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(search).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) SavedSearchDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.SavedSearch)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
