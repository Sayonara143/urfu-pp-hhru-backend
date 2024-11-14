package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) UserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var user models.User
	if err := tx.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var user models.User
	if err := tx.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) Users(ctx context.Context, limit, offset int) (int, []*models.User, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var users []*models.User
	total, err := tx.NewSelect().
		Model(&users).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, users, nil
}

func (s *Storage) UserInsert(ctx context.Context, user *models.User) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UserUpdate(ctx context.Context, user *models.User) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(user).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UserDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.User)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
