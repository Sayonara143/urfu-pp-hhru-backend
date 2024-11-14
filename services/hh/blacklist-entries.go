package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

// BlacklistEntryByID получает запись чёрного списка по ID.
func (s *Service) BlacklistEntryByID(ctx context.Context, id uuid.UUID) (*models.BlacklistEntry, error) {
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

// BlacklistEntries возвращает список записей чёрного списка с пагинацией.
func (s *Service) BlacklistEntries(ctx context.Context, limit, offset int) (int, []*models.BlacklistEntry, error) {
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

// BlacklistEntryInsert добавляет новую запись в чёрный список.
func (s *Service) BlacklistEntryInsert(ctx context.Context, entry *models.BlacklistEntry) error {
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

// BlacklistEntryUpdate обновляет запись в чёрном списке.
func (s *Service) BlacklistEntryUpdate(ctx context.Context, entry *models.BlacklistEntry) error {
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

// BlacklistEntryDelete удаляет запись из чёрного списка по ID.
func (s *Service) BlacklistEntryDelete(ctx context.Context, id uuid.UUID) error {
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
