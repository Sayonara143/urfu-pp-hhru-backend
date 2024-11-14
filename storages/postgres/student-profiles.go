package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) StudentProfileByID(ctx context.Context, id uuid.UUID) (*models.StudentProfile, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var profile models.StudentProfile
	if err := tx.NewSelect().
		Model(&profile).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (s *Storage) StudentProfiles(ctx context.Context, limit, offset int) (int, []*models.StudentProfile, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var profiles []*models.StudentProfile
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

func (s *Storage) StudentProfileInsert(ctx context.Context, profile *models.StudentProfile) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(profile).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) StudentProfileUpdate(ctx context.Context, profile *models.StudentProfile) error {
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

func (s *Storage) StudentProfileDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.StudentProfile)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
