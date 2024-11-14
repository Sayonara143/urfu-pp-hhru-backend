package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) JobApplicationByID(ctx context.Context, id uuid.UUID) (*models.JobApplication, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var application models.JobApplication
	if err := tx.NewSelect().
		Model(&application).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &application, nil
}

func (s *Storage) JobApplications(ctx context.Context, limit, offset int) (int, []*models.JobApplication, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var applications []*models.JobApplication
	total, err := tx.NewSelect().
		Model(&applications).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, applications, nil
}

func (s *Storage) JobApplicationInsert(ctx context.Context, application *models.JobApplication) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(application).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) JobApplicationUpdate(ctx context.Context, application *models.JobApplication) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(application).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) JobApplicationDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.JobApplication)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
