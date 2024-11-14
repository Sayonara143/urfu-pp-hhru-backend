package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) ReviewByID(ctx context.Context, id uuid.UUID) (*models.Review, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var review models.Review
	if err := tx.NewSelect().
		Model(&review).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &review, nil
}

func (s *Storage) Reviews(ctx context.Context, limit, offset int) (int, []*models.Review, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var reviews []*models.Review
	total, err := tx.NewSelect().
		Model(&reviews).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, reviews, nil
}

func (s *Storage) ReviewInsert(ctx context.Context, review *models.Review) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(review).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ReviewUpdate(ctx context.Context, review *models.Review) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(review).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ReviewDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.Review)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
