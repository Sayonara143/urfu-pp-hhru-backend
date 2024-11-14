package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) BlacklistByID(ctx context.Context, id uuid.UUID) (*models.BlacklistEntry, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	entry, err := s.store.BlacklistEntryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entry, s.store.TxCommit(ctx)
}

func (s *Service) Blacklist(ctx context.Context, limit, offset int) (int, []*models.BlacklistEntry, error) {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	total, entries, err := s.store.BlacklistEntries(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, entries, s.store.TxCommit(ctx)
}

func (s *Service) BlacklistInsert(ctx context.Context, entry *models.BlacklistEntry) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err = s.store.BlacklistEntryInsert(ctx, entry); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) BlacklistUpdate(ctx context.Context, entry *models.BlacklistEntry) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err = s.store.BlacklistEntryUpdate(ctx, entry); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) BlacklistDelete(ctx context.Context, id uuid.UUID) error {
	var err error
	ctx, err = s.store.CtxWithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.store.TxRollback(ctx)
		}
	}()

	if err = s.store.BlacklistEntryDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
