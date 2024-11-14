package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) EmployerProfileByID(ctx context.Context, id uuid.UUID) (*models.EmployerProfile, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var profile models.EmployerProfile
	if err := tx.NewSelect().
		Model(&profile).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (s *Storage) EmployerProfiles(ctx context.Context, limit, offset int) (int, []*models.EmployerProfile, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var profiles []*models.EmployerProfile
	total, err := tx.NewSelect().
		Model(&profiles).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, profiles, nil
}

func (s *Storage) EmployerProfileInsert(ctx context.Context, profile *models.EmployerProfile) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(profile).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) EmployerProfileUpdate(ctx context.Context, profile *models.EmployerProfile) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(profile).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) EmployerProfileDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.EmployerProfile)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
