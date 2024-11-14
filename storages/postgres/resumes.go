package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) ResumeByID(ctx context.Context, id uuid.UUID) (*models.Resume, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var resume models.Resume
	if err := tx.NewSelect().
		Model(&resume).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &resume, nil
}

func (s *Storage) Resumes(ctx context.Context, limit, offset int) (int, []*models.Resume, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var resumes []*models.Resume
	total, err := tx.NewSelect().
		Model(&resumes).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, resumes, nil
}

func (s *Storage) ResumeInsert(ctx context.Context, resume *models.Resume) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(resume).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ResumeUpdate(ctx context.Context, resume *models.Resume) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(resume).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ResumeDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.Resume)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
