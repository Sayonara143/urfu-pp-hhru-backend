package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

func (s *Service) JobVacancyByID(ctx context.Context, id uuid.UUID) (*models.JobVacancy, error) {
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

	vacancy, err := s.store.JobVacancyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return vacancy, s.store.TxCommit(ctx)
}

func (s *Service) JobVacancies(ctx context.Context, limit, offset int) (int, []*models.JobVacancy, error) {
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

	total, vacancies, err := s.store.JobVacancies(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, vacancies, s.store.TxCommit(ctx)
}

func (s *Service) JobVacancyInsert(ctx context.Context, vacancy *models.JobVacancy) error {
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

	if err = s.store.JobVacancyInsert(ctx, vacancy); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) JobVacancyUpdate(ctx context.Context, vacancy *models.JobVacancy) error {
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

	if err = s.store.JobVacancyUpdate(ctx, vacancy); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

func (s *Service) JobVacancyDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.JobVacancyDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
