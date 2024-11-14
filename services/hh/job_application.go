package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) JobApplicationByID(ctx context.Context, id uuid.UUID) (*models.JobApplication, error) {
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

	application, err := s.store.JobApplicationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return application, s.store.TxCommit(ctx)
}

func (s *Service) JobApplications(ctx context.Context, limit, offset int) (int, []*models.JobApplication, error) {
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

	total, applications, err := s.store.JobApplications(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, applications, s.store.TxCommit(ctx)
}

func (s *Service) JobApplicationInsert(ctx context.Context, application *models.JobApplication) error {
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

	if err = s.store.JobApplicationInsert(ctx, application); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) JobApplicationUpdate(ctx context.Context, application *models.JobApplication) error {
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

	if err = s.store.JobApplicationUpdate(ctx, application); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) JobApplicationDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.JobApplicationDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
