package postgres

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Storage) BlacklistByID(ctx context.Context, id uuid.UUID) (*models.BlacklistEntry, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return nil, ErrTxNotFound
	}

	var entry models.BlacklistEntry
	if err := tx.NewSelect().
		Model(&entry).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &entry, nil
}

func (s *Storage) Blacklist(ctx context.Context, limit, offset int) (int, []*models.BlacklistEntry, error) {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return 0, nil, ErrTxNotFound
	}

	var entries []*models.BlacklistEntry
	total, err := tx.NewSelect().
		Model(&entries).
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if err != nil {
		return 0, nil, err
	}

	return total, entries, nil
}

func (s *Storage) BlacklistInsert(ctx context.Context, entry *models.BlacklistEntry) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewInsert().Model(entry).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) BlacklistUpdate(ctx context.Context, entry *models.BlacklistEntry) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewUpdate().
		Model(entry).
		OmitZero().
		WherePK().
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Storage) BlacklistDelete(ctx context.Context, id uuid.UUID) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return ErrTxNotFound
	}

	if _, err := tx.NewDelete().
		Model((*models.BlacklistEntry)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
