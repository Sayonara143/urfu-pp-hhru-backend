package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) InterviewByID(ctx context.Context, id uuid.UUID) (*models.Interview, error) {
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

	interview, err := s.store.InterviewByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return interview, s.store.TxCommit(ctx)
}

func (s *Service) Interviews(ctx context.Context, limit, offset int) (int, []*models.Interview, error) {
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

	total, interviews, err := s.store.Interviews(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, interviews, s.store.TxCommit(ctx)
}

func (s *Service) InterviewInsert(ctx context.Context, interview *models.Interview) error {
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

	if err = s.store.InterviewInsert(ctx, interview); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) InterviewUpdate(ctx context.Context, interview *models.Interview) error {
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

	if err = s.store.InterviewUpdate(ctx, interview); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) InterviewDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.InterviewDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
