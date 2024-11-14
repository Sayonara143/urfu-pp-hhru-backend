package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) JobVacancyByID(ctx context.Context, id uuid.UUID) (*models.JobVacancy, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var vacancy models.JobVacancy
	if err := tx.NewSelect().
		Model(&vacancy).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &vacancy, nil
}

func (s *Storage) JobVacancies(ctx context.Context, limit, offset int) (int, []*models.JobVacancy, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var vacancies []*models.JobVacancy
	total, err := tx.NewSelect().
		Model(&vacancies).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, vacancies, nil
}

func (s *Storage) JobVacancyInsert(ctx context.Context, vacancy *models.JobVacancy) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(vacancy).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) JobVacancyUpdate(ctx context.Context, vacancy *models.JobVacancy) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(vacancy).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) JobVacancyDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.JobVacancy)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
