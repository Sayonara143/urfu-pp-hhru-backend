package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) InterviewByID(ctx context.Context, id uuid.UUID) (*models.Interview, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var interview models.Interview
	if err := tx.NewSelect().
		Model(&interview).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &interview, nil
}

func (s *Storage) Interviews(ctx context.Context, limit, offset int) (int, []*models.Interview, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var interviews []*models.Interview
	total, err := tx.NewSelect().
		Model(&interviews).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, interviews, nil
}

func (s *Storage) InterviewInsert(ctx context.Context, interview *models.Interview) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(interview).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) InterviewUpdate(ctx context.Context, interview *models.Interview) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(interview).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) InterviewDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.Interview)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
