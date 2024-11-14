package hh

import (
	"context"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/google/uuid"
)

// EmployerProfileByID получает профиль работодателя по ID.
func (s *Service) EmployerProfileByID(ctx context.Context, id uuid.UUID) (*models.EmployerProfile, error) {
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

	profile, err := s.store.EmployerProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return profile, s.store.TxCommit(ctx)
}

// EmployerProfiles возвращает список профилей работодателей с пагинацией.
func (s *Service) EmployerProfiles(ctx context.Context, limit, offset int) (int, []*models.EmployerProfile, error) {
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

	total, profiles, err := s.store.EmployerProfiles(ctx, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	return total, profiles, s.store.TxCommit(ctx)
}

// EmployerProfileInsert добавляет новый профиль работодателя.
func (s *Service) EmployerProfileInsert(ctx context.Context, profile *models.EmployerProfile) error {
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

	if err = s.store.EmployerProfileInsert(ctx, profile); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// EmployerProfileUpdate обновляет профиль работодателя.
func (s *Service) EmployerProfileUpdate(ctx context.Context, profile *models.EmployerProfile) error {
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

	if err = s.store.EmployerProfileUpdate(ctx, profile); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}

// EmployerProfileDelete удаляет профиль работодателя по ID.
func (s *Service) EmployerProfileDelete(ctx context.Context, id uuid.UUID) error {
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

	if err = s.store.EmployerProfileDelete(ctx, id); err != nil {
		return err
	}

	return s.store.TxCommit(ctx)
}
